package main

import (
	"context"
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strangerbot/keyboard"
	"strangerbot/service"
)

func handleCallbackQuery(callbackQuery *tgbotapi.CallbackQuery) {

	ctx := context.TODO()
	_ = ctx

	if len(callbackQuery.Data) == 0 {
		return
	}

	u, err := retrieveOrCreateUser(callbackQuery.Message.Chat.ID)
	_ = u
	if err != nil {
		log.Println(err)
		return
	}

	data := new(keyboard.KeyboardCallbackDataPlus)
	if err := json.Unmarshal([]byte(callbackQuery.Data), &data); err != nil {
		log.Println("json unamrshal error", err.Error())
		return
	}

	var (
		msgs []*tgbotapi.MessageConfig
		cbs  []tgbotapi.CallbackConfig
	)
	switch data.ButtonType {
	case keyboard.BUTTON_TYPE_MENU:
		msgs, err = service.ServiceMenu(ctx, callbackQuery.Message.Chat.ID, data)
		if err != nil {
			return
		}
	case keyboard.BUTTON_TYPE_QUESTION:

	case keyboard.BUTTON_TYPE_OPTION:
		msgs, cbs, err = service.ServiceQuestionOption(ctx, callbackQuery.Message.Chat.ID, data)
		if err != nil {
			return
		}
	}

	// send callback
	for _, cb := range cbs {
		cb.CallbackQueryID = callbackQuery.ID
		_, err = telegramBot.AnswerCallbackQuery(cb)
		if err != nil {
			log.Println(err.Error())
		}
	}

	if len(msgs) == 0 {
		return
	}

	// first delete pre msg
	{
		msg := tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID)
		_, err = telegramBot.Send(msg)
		if err != nil {
			return
		}
	}

	// send new message
	for _, msg := range msgs {
		_, err = telegramBot.Send(msg)
		if err != nil {
			return
		}
	}

	//switch obj.OptionType {
	//case GenderOptionType:
	//
	//	// update gender
	//	updateGender(u.ID, obj.OptionValue)
	//

	//
	//	{
	//		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, fmt.Sprintf("You selected %s as your Gender. %s", obj.GetOptionText(), obj.GetOptionNoteText()))
	//		telegramBot.Send(msg)
	//	}
	//
	//	{
	//		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, `What gender do you want to match with?`)
	//
	//		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{
	//			{
	//				Text:         MatchModeOptionMaleText + MatchModeOptionMaleNoteText,
	//				CallbackData: &MatchModeMale,
	//			},
	//			{
	//				Text:         MatchModeOptionFemaleText + MatchModeOptionFemaleNoteText,
	//				CallbackData: &MatchModeFemale,
	//			},
	//			{
	//				Text:         MatchModeOptionAnythingText + MatchModeOptionAnythingNoteText,
	//				CallbackData: &MatchModeAnything,
	//			},
	//		})
	//
	//		_, err := telegramBot.Send(msg)
	//		if err != nil {
	//			log.Println(err.Error())
	//		}
	//	}
	//
	//case MatchModeOptionType:
	//
	//	// update gender
	//	updateMathMode(u.ID, obj.OptionValue)
	//
	//	// handle message
	//	{
	//		msg := tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID)
	//		telegramBot.Send(msg)
	//	}
	//
	//	{
	//		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, fmt.Sprintf("You selected %s as your Match. %s", obj.GetOptionText(), obj.GetOptionNoteText()))
	//		telegramBot.Send(msg)
	//	}
	//
	//	{
	//		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, `What are you here for?`)
	//
	//		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{
	//			{
	//				Text:         GoalOptionDatingText + GoalOptionDatingNoteText,
	//				CallbackData: &GoalDating,
	//			},
	//			{
	//				Text:         GoalOptionFriendsText + GoalOptionFriendsNoteText,
	//				CallbackData: &GoalFriends,
	//			},
	//		})
	//
	//		_, err := telegramBot.Send(msg)
	//		if err != nil {
	//			log.Println(err.Error())
	//		}
	//	}
	//
	//case GoalOptionType:
	//
	//	// update tags
	//	updateTags(u.ID, obj.GetOptionText())
	//
	//	// handle message
	//	{
	//		msg := tgbotapi.NewDeleteMessage(callbackQuery.Message.Chat.ID, callbackQuery.Message.MessageID)
	//		telegramBot.Send(msg)
	//	}
	//
	//	{
	//		msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, fmt.Sprintf("You selected %s as your Goal. %s", obj.GetOptionText(), obj.GetOptionNoteText()))
	//		telegramBot.Send(msg)
	//	}
	//
	//}

}
