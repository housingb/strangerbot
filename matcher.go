package main

import (
	"context"
	"fmt"
	"log"
	"strings"

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

		matchUser, matchingOptions, matchUsersOptions, err := service.ServiceMatch(ctx, user.ChatID)
		if err != nil {
			log.Printf("Error retrieving available users: %s", err)
			continue
		}

		if matchUser == nil {
			continue
		}

		createMatch(user.ChatID, user.ID, matchUser.ChatID, matchUser.ID, matchingOptions, matchUsersOptions)

	}

}

func createMatch(userChatId, userId, matchUserChatId, matchUserId int64, matchingOptions []string, matchingUserOptions []string) {
	query := "UPDATE users SET match_chat_id = ? WHERE id = ?"

	db.Exec(query, userChatId, matchUserId)
	db.Exec(query, matchUserChatId, userId)

	telegram.SendMessage(matchUserChatId, fmt.Sprintf(vars.MatchedMessage, strings.Join(matchingUserOptions, ",")), emptyOpts)
	telegram.SendMessage(userChatId, fmt.Sprintf(vars.MatchedMessage, strings.Join(matchingOptions, ",")), emptyOpts)
}
