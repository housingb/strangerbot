package model

import (
	"context"

	"strangerbot/repository/gorm_global"
)

type UserQuestionData struct {
	gorm_global.ColumnCreateModifyDeleteTime
	ChatId     int64
	QuestionId int64
	OptionId   int64
	Value      string
}

func (u *UserQuestionData) TableName() string {
	return "bot_user_question_data"
}

type UserQuestionDataList []*UserQuestionData

func (u UserQuestionDataList) GetUserQuestionDataByOptions(ctx context.Context, options Options) UserQuestionDataList {

	rs := make([]*UserQuestionData, 0, len(u))
	for i, item := range u {
		for _, option := range options {
			if item.OptionId == option.ID {
				rs = append(rs, u[i])
			}
		}
	}

	return rs
}
