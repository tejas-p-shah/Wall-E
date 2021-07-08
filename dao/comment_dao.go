package dao

import (
	"fmt"

	"github.com/tejas-p-shah/Wall-E/model"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/bson"
)

func GetPostComments(PostID int64) (*model.Comment, error) {

	db := GetDb().Collection("comment")
	fmt.Println(db)

	return nil, fmt.Errorf("user %v was not found", PostID)
}
