package service

import (
	"context"
	"errors"
	"math/rand"

	"strangerbot/repository"
	"strangerbot/repository/model"
)

var (
	ErrUserNotFillAllQuestion = errors.New("user not fill all question")
)

func ServiceCheckUserFillFull(ctx context.Context, chatId int64) (bool, error) {

	repo := repository.GetRepository()

	// find all question
	questions, err := repo.GetAllQuestion(ctx)
	if err != nil {
		return false, err
	}

	// find user all user question data
	userQuestionData, err := repo.GetUserQuestionDataByUser(ctx, chatId)
	if err != nil {
		return false, err
	}

	// check user fill full
	fillFull, _ := model.Questions(questions).CheckUserFillFull(userQuestionData)
	if fillFull {
		return true, nil
	}

	return false, nil
}

func ServiceMatch(ctx context.Context, chatId int64) (*model.User, []string, []string, error) {

	repo := repository.GetRepository()

	// find all question
	questions, err := repo.GetAllQuestion(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	// find all question option
	options, err := repo.GetAllOption(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	// find user all user question data
	userQuestionData, err := repo.GetUserQuestionDataByUser(ctx, chatId)
	if err != nil {
		return nil, nil, nil, err
	}

	// find matching options and user matching options value
	chatIds, userMatchingData, err := repo.GetChatByMatching(ctx, chatId, questions, options, userQuestionData)
	if err != nil {
		return nil, nil, nil, err
	}

	if len(chatIds) == 0 {
		return nil, nil, nil, nil
	}

	// shuffle chat id
	matchChatId := shuffleChatId(chatIds)

	// find user
	user, err := repo.GetUserByChatId(ctx, matchChatId)
	if err != nil {
		return nil, nil, nil, err
	}

	matchUserQuestionData, err := repo.GetUserQuestionDataByUser(ctx, matchChatId)
	if err != nil {
		return nil, nil, nil, err
	}

	// return option string
	optionIds := userMatchingData.GetOptionIds()
	matchingQuestion := model.Questions(questions).GetMatchingQuestion()
	matchingOptions := model.Options(options).GetQuestionOptions(ctx, matchingQuestion)
	userMatchingDataBe := model.UserQuestionDataList(matchUserQuestionData).GetUserQuestionDataByOptions(ctx, matchingOptions)
	matchOptionsIds := userMatchingDataBe.GetOptionIds()

	return user, model.Options(options).GetOptionsByIds(optionIds), model.Options(options).GetOptionsByIds(matchOptionsIds), nil
}

func shuffleChatId(a []int64) int64 {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a[0]
}
