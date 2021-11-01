package repository

import (
	"context"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"strangerbot/repository/model"
)

func (r *Repository) UserQuestionDataAdd(ctx context.Context, po *model.UserQuestionData) error {

	if err := r.db.Create(&po).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return nil
		}
		return err
	}

	return nil
}

func (p *Repository) GetUserQuestionDataByOptionAndChat(ctx context.Context, optionId int64, chatId int64) (*model.UserQuestionData, error) {

	po := &model.UserQuestionData{}

	if err := p.db.Where("option_id = ? AND chat_id = ? AND is_del = 0", optionId, chatId).First(&po).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return po, nil
}

func (p *Repository) GetUserQuestionDataByQuestion(ctx context.Context, questionId int64, chatId int64) ([]*model.UserQuestionData, error) {

	q := p.db.Where("question_id = ? and chat_id = ? AND is_del = 0", questionId, chatId)

	var list []*model.UserQuestionData

	if err := q.Model(&model.UserQuestionData{}).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, err
	}

	return list, nil
}

func (p *Repository) DeleteUserQuestionData(ctx context.Context, po *model.UserQuestionData) error {

	if err := p.db.Delete(&po).Error; err != nil {
		return err
	}

	return nil
}

func (p *Repository) DeleteUserQuestionDataByQuestion(ctx context.Context, chatId int64, questionId int64) error {

	if err := p.db.Where("chat_id = ? AND question_id = ?", chatId, questionId).Delete(&model.UserQuestionData{}).Error; err != nil {
		return err
	}

	return nil
}

func (p *Repository) GetUserQuestionDataByUser(ctx context.Context, chatId int64) ([]*model.UserQuestionData, error) {
	q := p.db.Where("chat_id = ? AND is_del = 0", chatId)

	var list []*model.UserQuestionData

	if err := q.Model(&model.UserQuestionData{}).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		log.Println(err.Error())
		return nil, err
	}

	return list, nil
}

func (p *Repository) GetChatByMatching(ctx context.Context, chatId int64, questions model.Questions, options model.Options, userQuestionData model.UserQuestionDataList) ([]int64, error) {

	// matching options data
	matchingQuestion := questions.GetMatchingQuestion()
	profileQuestion := questions.GetMappingQuestion(matchingQuestion)
	matchingOptions := options.GetQuestionOptions(ctx, matchingQuestion)
	profileOptions := options.GetQuestionOptions(ctx, profileQuestion)
	userMatchingData := userQuestionData.GetUserQuestionDataByOptions(ctx, matchingOptions)
	userProfileData := userQuestionData.GetUserQuestionDataByOptions(ctx, profileOptions)

	//selectOptions := make([]int64, 0, len(userMatchingData))

	questionMatchingMap := make(map[int64]bool)
	questionProfileMatchingMap := make(map[int64]bool)
	allMatchingOptions := make([]int64, 0, len(userMatchingData))
	allProfileMatchingOptions := make([]int64, 0, len(userProfileData))
	matchingQuestionNum := 0
	profileMatchingQuestionNum := 0

	// group by question
	for _, item := range userMatchingData {

		option := matchingOptions.GetOption(item.OptionId)
		if option == nil {
			continue
		}

		if option.MatchingOptionId == 0 {
			continue
		}

		allMatchingOptions = append(allMatchingOptions, option.MatchingOptionId)
		if _, ok := questionMatchingMap[item.QuestionId]; !ok {
			questionMatchingMap[item.QuestionId] = true
			matchingQuestionNum++
		}

	}

	for _, item := range userProfileData {

		// get profile option mapping matching option
		option := matchingOptions.GetOptionByMapping(item.OptionId)
		if option == nil {
			continue
		}

		if option.MatchingOptionId == 0 {
			continue
		}

		allProfileMatchingOptions = append(allProfileMatchingOptions, option.ID)
		if _, ok := questionProfileMatchingMap[option.QuestionId]; !ok {
			questionProfileMatchingMap[option.QuestionId] = true
			profileMatchingQuestionNum++
		}

	}

	for questionId, _ := range questionProfileMatchingMap {
		anyOption := options.IsHasAnythingOption(questionId)
		if anyOption != nil {
			allProfileMatchingOptions = append(allProfileMatchingOptions, anyOption.ID)
		}
	}

	// build sql
	var sub *gorm.DB

	//sub = p.db.Raw("SELECT chat_id FROM (SELECT chat_id,COUNT(*) AS cnt FROM (SELECT chat_id,question_id FROM bot_user_question_data WHERE option_id IN(?) AND chat_id IN((SELECT chat_id FROM users WHERE chat_id != ? AND available = 1 AND match_chat_id IS NULL)) GROUP BY chat_id,question_id) AS bot_user_question_data GROUP BY chat_id) AS bot_user_question_data WHERE cnt = ?", allMatchingOptions, chatId, matchingQuestionNum)

	sub = p.db.Raw("SELECT chat_id FROM (SELECT chat_id,COUNT(*) AS cnt FROM (SELECT chat_id,question_id FROM bot_user_question_data WHERE option_id IN(?) AND chat_id IN((SELECT chat_id FROM users WHERE available = 1 AND match_chat_id IS NULL)) GROUP BY chat_id,question_id) AS bot_user_question_data GROUP BY chat_id) AS bot_user_question_data WHERE cnt = ?", allMatchingOptions, matchingQuestionNum)

	sub = p.db.Raw("SELECT chat_id FROM (SELECT chat_id,COUNT(*) AS cnt FROM (SELECT chat_id,question_id FROM bot_user_question_data WHERE chat_id IN(?) AND option_id IN(?)) AS bot_user_question_data GROUP BY chat_id) AS bot_user_question_data WHERE cnt = ?", sub.QueryExpr(), allProfileMatchingOptions, profileMatchingQuestionNum)

	var data []struct {
		ChatId int64
	}

	if err := sub.Scan(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	rs := make([]int64, 0, len(data))
	for _, item := range data {
		rs = append(rs, item.ChatId)
	}

	return rs, nil
}
