package service

import (
	"context"

	"strangerbot/repository"
	"strangerbot/repository/model"
)

func ServiceMatchedDetailRecord(ctx context.Context, chatId int64, matchedChatId int64) error {

	repo := repository.GetRepository()

	// add
	po := &model.MatchedDetail{
		ChatId:      chatId,
		MatchChatId: matchedChatId,
	}

	if err := repo.MatchedDetailAdd(ctx, po); err != nil {
		return err
	}

	return nil
}
