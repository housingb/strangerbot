package main

import (
	"context"
	"log"

	"strangerbot/service"
	"strangerbot/vars"
)

func matchUsers(chatIDs <-chan int64) {

	ctx := context.TODO()

	for c := range chatIDs {

		user, err := retrieveUser(c)

		if err != nil {
			log.Printf("Error in matcher: %s", err)
			continue
		}

		if !user.Available || user.MatchChatID.Valid {
			log.Println("User already assigned")
			continue
		}

		matchUser, err := service.ServiceMatch(ctx, user.ChatID)
		if err != nil {
			log.Printf("Error retrieving available users: %s", err)
			continue
		}

		if matchUser == nil {
			continue
		}

		createMatch(user.ChatID, user.ID, matchUser.ChatID, matchUser.ID)

	}

}

func createMatch(userChatId, userId, matchUserChatId, matchUserId int64) {
	query := "UPDATE users SET match_chat_id = ? WHERE id = ?"

	db.Exec(query, userChatId, matchUserId)
	db.Exec(query, matchUserChatId, userId)

	telegram.SendMessage(matchUserChatId, vars.MatchedMessage, emptyOpts)
	telegram.SendMessage(userChatId, vars.MatchedMessage, emptyOpts)
}
