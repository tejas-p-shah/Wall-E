package dao

import (
	"context"
	"fmt"

	"github.com/tejas-p-shah/Wall-E/config"
	"github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPostComments(PostID int64) (*model.Comment, error) {

	return nil, fmt.Errorf("user %v was not found", PostID)
}

func AddNewComment(comment *model.Comment) (*mongo.InsertOneResult, error) {

	client, err := config.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(config.DB).Collection(config.POSTS)
	result, err := collection.InsertOne(context.TODO(), comment)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateComment(comment *model.Comment) error {

	return fmt.Errorf("comment was not found")
}

func DeleteComment(comment *model.Comment) error {

	return fmt.Errorf("comment was not found")
}
