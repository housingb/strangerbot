package service

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strangerbot/keyboard"
	"strangerbot/repository"
	"strangerbot/repository/model"
)

func ServiceQuestionOption(ctx context.Context, chatId int64, data *keyboard.KeyboardCallbackDataPlus) ([]*tgbotapi.MessageConfig, error) {

	repo := repository.GetRepository()

	// get option
	option, err := repo.GetOptionById(ctx, data.ButtonRelId)
	if err != nil {
		return nil, nil
	}

	// get question
	question, err := repo.GetQuestionById(ctx, option.QuestionId)
	if err != nil {
		return nil, nil
	}

	switch question.FrontendType {
	case model.FRONTEND_TYPE_SELECT:

		// delete question old option
		err := repo.DeleteUserQuestionDataByQuestion(ctx, chatId, question.ID)
		if err != nil {
			return nil, nil
		}

	case model.FRONTEND_TYPE_MULTI_SELECT:

		userQuestionData, err := repo.GetUserQuestionDataByOptionAndChat(ctx, option.ID, chatId)
		if err != nil {
			return nil, err
		}

		// delete
		if userQuestionData != nil {
			if err := repo.DeleteUserQuestionData(ctx, userQuestionData); err != nil {
				return nil, nil
			}

			// back to menu
			menu, err := repo.GetMenuByQuestionId(ctx, option.QuestionId)
			if err != nil {
				return nil, nil
			}

			return serviceBackMenu(ctx, chatId, repo, menu)
		}

	}

	// add
	po := &model.UserQuestionData{
		ChatId:     chatId,
		QuestionId: option.QuestionId,
		OptionId:   option.ID,
		Value:      option.Value,
	}

	if err := repo.UserQuestionDataAdd(ctx, po); err != nil {
		return nil, err
	}

	// back to menu
	menu, err := repo.GetMenuByQuestionId(ctx, option.QuestionId)
	if err != nil {
		return nil, nil
	}

	return serviceBackMenu(ctx, chatId, repo, menu)

}
