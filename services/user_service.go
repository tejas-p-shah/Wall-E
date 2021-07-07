package services

import (
	"github.com/tejas-p-shah/Wall-E/dao"
	"github.com/tejas-p-shah/Wall-E/model"
)

func GetUser(userID int64) (*model.User, error) {
	user, err := dao.GetUser(userID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
