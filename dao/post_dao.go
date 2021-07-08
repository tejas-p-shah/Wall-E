package dao

import (
	"fmt"

	"github.com/tejas-p-shah/Wall-E/model"
)

func GetUserPosts(userID int64) (*model.Post, error) {

	db := GetDb().Collection("comment")
	fmt.Println(db)
	return nil, fmt.Errorf("user %v was not found", userID)
}
