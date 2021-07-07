package services

import (
	"github.com/tejas-p-shah/Wall-E/dao"
	"github.com/tejas-p-shah/Wall-E/model"
)

func GetUserPost(userID int64) (*model.Post, error) {
	post, err := dao.GetUserPosts(userID)

	if err != nil {
		return nil, err
	}

	return post, nil
}
