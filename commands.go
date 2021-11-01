package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/Machiel/telegrambot"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strangerbot/repository"
	"strangerbot/service"
	"strangerbot/vars"
)

// CommandHandler supplies an interface for handling messages
type commandHandler func(u User, m *tgbotapi.Message) bool

func RetrySendMessage(id int64, text string, options telegrambot.SendMessageOptions) (string, error) {

	var rsp string
	var err error

	for i := 0; i < 20; i++ {

		rsp, err = telegram.SendMessage(id, text, options)
		if err == nil {
			break
		}

		time.Sleep(10 * time.Duration(i) * time.Millisecond)

	}

	return rsp, err
}

func RetrySend(c tgbotapi.Chattable) (tgbotapi.Message, error) {

	var rsp tgbotapi.Message
	var err error

	for i := 0; i < 20; i++ {

		rsp, err = telegramBot.Send(c)
		if err == nil {
			break
		}

		time.Sleep(10 * time.Duration(i) * time.Millisecond)

	}

	return rsp, err
}

func commandDisablePictures(u User, m *tgbotapi.Message) bool {
	if len(m.Text) < 7 || strings.ToLower(m.Text[0:7]) != "/nopics" {
		return false
	}

	if u.AllowPictures {
		db.Exec("UPDATE users SET allow_pictures = 0 WHERE id = ?", u.ID)
		_, _ = RetrySendMessage(u.ChatID, "Users won't be able to send you photos anymore!", emptyOpts)
		return true
	}

	db.Exec("UPDATE users SET allow_pictures = 1 WHERE id = ?", u.ID)
	_, _ = RetrySendMessage(u.ChatID, "Users can now send you photos!", emptyOpts)
	return true
}

func commandStart(u User, m *tgbotapi.Message) bool {

	if len(m.Text) < 6 {
		return false
	}

	if strings.ToLower(m.Text[0:6]) != "/start" {
		return false
	}

	if u.Available {
		return false
	}

	if u.MatchChatID.Valid {
		return false
	}

	isFull, err := service.ServiceCheckUserFillFull(nil, u.ChatID)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	if !isFull {
		_, _ = RetrySendMessage(u.ChatID, vars.NotProfileFinishMessage, emptyOpts)
		return false
	}

	db.Exec("UPDATE users SET available = 1 WHERE chat_id = ?", u.ChatID)

	_, _ = RetrySendMessage(u.ChatID, vars.StartMatchingMessage, emptyOpts)
	startJobs <- u.ChatID

	return true
}

func commandStop(u User, m *tgbotapi.Message) bool {

	if len(m.Text) < 4 {
		return false
	}

	rightCommand := strings.ToLower(m.Text[0:4]) == "/bye" || strings.ToLower(m.Text[0:4]) == "/end"

	if !rightCommand {
		return false
	}

	if !u.Available {
		return false
	}

	_, _ = RetrySendMessage(u.ChatID, "We're ending the conversation...", emptyOpts)

	endConversationQueue <- EndConversationEvent{ChatID: u.ChatID}

	return true
}

func commandReport(u User, m *tgbotapi.Message) bool {

	if len(m.Text) < 7 || strings.ToLower(m.Text[0:7]) != "/report" {
		return false
	}

	if !u.Available || !u.MatchChatID.Valid {
		return false
	}

	report := m.Text[7:]
	report = strings.TrimSpace(report)

	if len(report) == 0 {
		_, _ = RetrySendMessage(u.ChatID, "Usage /report: /report <reason>", emptyOpts)
		return true
	}

	partner, err := retrieveUser(u.MatchChatID.Int64)

	if err != nil {
		log.Println("Error retrieving partner")
		return true
	}

	db.Exec("INSERT INTO reports (user_id, report, reporter_id, created_at) VALUES (?, ?, ?, ?)", partner.ID, report, u.ID, time.Now())

	_, _ = RetrySendMessage(u.ChatID, "User has been reported!", emptyOpts)

	return true
}

func commandMessage(u User, m *tgbotapi.Message) bool {

	if !u.Available {
		return false
	}

	if !u.MatchChatID.Valid {
		return false
	}

	chatID := u.MatchChatID.Int64
	partner, err := retrieveUser(chatID)

	if err != nil {
		log.Println("[ERROR] Could not retrieve partner %d", chatID)
		return false
	}

	if m.Photo != nil && len(*m.Photo) > 0 {

		if !partner.AllowPictures {
			_, _ = RetrySendMessage(chatID, "User tried to send you a photo, but you disabled this,  you can enable photos by using the /nopics command", emptyOpts)
			_, _ = RetrySendMessage(u.ChatID, "User disabled photos, and will not receive your photos", emptyOpts)
			return true
		}

		var toSend tgbotapi.PhotoSize

		for _, t := range *m.Photo {
			if t.FileSize > toSend.FileSize {
				toSend = t
			}
		}

		_, _ = RetrySendMessage(chatID, "User sends you a photo!", emptyOpts)
		_, err = telegram.SendPhoto(chatID, toSend.FileID, emptyOpts)

	} else if m.Sticker != nil {
		_, _ = RetrySendMessage(chatID, "User sends you a sticker!", emptyOpts)
		_, err = telegram.SendSticker(chatID, m.Sticker.FileID, emptyOpts)
	} else if m.Location != nil {
		_, _ = RetrySendMessage(chatID, "User sends you a location!", emptyOpts)
		_, err = telegram.SendLocation(chatID,
			m.Location.Latitude,
			m.Location.Longitude,
			emptyOpts,
		)
	} else if m.Document != nil {
		_, _ = RetrySendMessage(chatID, "User sends you a document!", emptyOpts)
		_, err = telegram.SendDocument(chatID, m.Document.FileID, emptyOpts)
	} else if m.Audio != nil {
		_, _ = RetrySendMessage(chatID, "User sends you an audio file!", emptyOpts)
		_, err = telegram.SendAudio(chatID, m.Audio.FileID, emptyOpts)
	} else if m.Video != nil {
		_, _ = RetrySendMessage(chatID, "User sends you a video file!", emptyOpts)
		_, err = telegram.SendVideo(chatID, m.Video.FileID, emptyOpts)
	} else {
		_, err = RetrySendMessage(chatID, "User: "+m.Text, emptyOpts)
	}

	if err != nil {
		log.Printf("Forward error: %s", err)
	}

	return true

}

func commandHelp(u User, m *tgbotapi.Message) bool {

	if len(m.Text) < 5 {
		return false
	}

	if strings.ToLower(m.Text[0:5]) != "/help" {
		return false
	}

	_, _ = RetrySendMessage(m.Chat.ID, vars.HelpMessage, emptyOpts)

	return true
}

func commandSetup(u User, m *tgbotapi.Message) bool {

	ctx := context.TODO()

	if len(m.Text) < 6 {
		return false
	}

	if strings.ToLower(m.Text[0:6]) != "/setup" {
		return false
	}

	repo := repository.GetRepository()
	menus, err := repo.GetMenuList(ctx, 0)
	if err != nil {
		return true
	}

	msg := tgbotapi.NewMessage(m.Chat.ID, vars.TopMenuMessage)
	msg.ReplyMarkup = menus.GetKeyboardMarkup()

	if _, err := RetrySend(msg); err != nil {
		log.Printf("Send err: %s", err.Error())
	}

	return true
}
