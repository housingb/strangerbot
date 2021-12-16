package repository

import (
	"context"

	"github.com/jinzhu/gorm"
	"strangerbot/repository/model"
)

func (p *Repository) GetChatCnt(ctx context.Context) (int64, error) {

	var cnt struct{ Cnt int64 }

	err := p.db.Select("COUNT(*) AS cnt").Model(model.User{}).Scan(&cnt).Error

	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return cnt.Cnt, nil
}

func (p *Repository) GetChatList(ctx context.Context, offset, limit int64) ([]*model.User, error) {

	var list []*model.User
	err := p.db.Offset(offset).Limit(limit).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (p *Repository) GetUserByChatId(ctx context.Context, chatId int64) (*model.User, error) {
	po := &model.User{}

	if err := p.db.Where("chat_id = ?", chatId).First(&po).Error; err != nil {
		return nil, err
	}

	return po, nil
}

func (p *Repository) GetEmailCnt(ctx context.Context, email string) (int64, error) {

	var cnt struct{ Cnt int64 }

	err := p.db.Select("COUNT(*) AS cnt").Where("email = ? AND is_verify = 1", email).Model(model.User{}).Scan(&cnt).Error

	if err == gorm.ErrRecordNotFound {
		return 0, nil
	}

	if err != nil {
		return 0, err
	}

	return cnt.Cnt, nil
}

func (p *Repository) GetVerifyUser(ctx context.Context, chatIds []int64, isVerify bool) ([]int64, error) {

	var list []*model.User
	err := p.db.Where("chat_id IN(?) AND is_verify = ?", chatIds, isVerify).Find(&list).Error
	if err != nil {
		return nil, err
	}

	rs := make([]int64, 0, len(list))
	for _, item := range list {
		rs = append(rs, item.ChatID)
	}

	return rs, nil
}
