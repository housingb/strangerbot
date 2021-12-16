package service

import (
	"context"
	"errors"
	"math/rand"

	"strangerbot/repository"
	"strangerbot/repository/model"
	"strangerbot/vars"
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

func ServiceMatch(ctx context.Context, chatId int64) (*model.User, error) {

	repo := repository.GetRepository()

	// find all question
	questions, err := repo.GetAllQuestion(ctx)
	if err != nil {
		return nil, err
	}

	// find all question option
	options, err := repo.GetAllOption(ctx)
	if err != nil {
		return nil, err
	}

	// find user all user question data
	userQuestionData, err := repo.GetUserQuestionDataByUser(ctx, chatId)
	if err != nil {
		return nil, err
	}

	// find matching options and user matching options value
	chatIds, _, err := repo.GetChatByMatching(ctx, chatId, questions, options, userQuestionData)
	if err != nil {
		return nil, err
	}

	if len(chatIds) == 0 {
		return nil, nil
	}

	// verified user matching
	userDataList := model.UserQuestionDataList(userQuestionData)
	vMOI := userDataList.GetFirstOptionIdByQuestionId(vars.MatchingQuestionId)
	switch vMOI {
	case vars.MatchingVerifiedOptionId:

		chatIds, err = repo.GetVerifyUser(ctx, chatIds, true)
		if err != nil {
			return nil, err
		}

	case vars.MatchingUnverifiedOptionId:

		chatIds, err = repo.GetVerifyUser(ctx, chatIds, false)
		if err != nil {
			return nil, err
		}

	}

	// shuffle chat id
	matchChatId := shuffleChatId(chatIds)

	// find user
	user, err := repo.GetUserByChatId(ctx, matchChatId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func shuffleChatId(a []int64) int64 {
	for i := range a {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	return a[0]
}

func ServiceCheckEmailUnique(ctx context.Context, email string) error {

	repo := repository.GetRepository()

	cnt, err := repo.GetEmailCnt(ctx, email)
	if err != nil {
		return err
	}

	if cnt > 0 {
		return errors.New("Sorry, This email has already been registered.")
	}

	return nil
}
