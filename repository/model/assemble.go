package model

type MatchUserData struct {
	ChatId                 int64
	User                   *User
	PersonalInfoQuestions  map[int64]*ProfileQuestion
	MatchCriteriaQuestions map[int64]*MatchingQuestion
	MatchChatId            int64
	MatchMatchUserData     *MatchUserData
}

type ProfileQuestion struct {
	ProfileQuestion *Question
	ProfileOptions  []*QuestionOption
}

type MatchingQuestion struct {
	MatchingQuestion *Question
	MatchingOptions  []*QuestionOption
}
