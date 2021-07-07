package dao

import (
	"fmt"

	"github.com/tejas-p-shah/Wall-E/model"
)

func GetUserPosts(userID int64) (*model.Post, error) {

	return nil, fmt.Errorf("user %v was not found", userID)
}
