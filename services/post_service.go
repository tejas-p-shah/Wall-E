package services

import (
	"time"

	"github.com/tejas-p-shah/Wall-E/dao"
	"github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserPosts(userName string) ([]model.Post, error) {
	posts, err := dao.GetUserPosts(userName)
	return posts, err
}

func UpdatePost(post model.Post) error {
	post.PostEditedDateTime = time.Now()
	err := dao.UpdatePost(&post)
	return err
}

func DeletePost(post model.Post) error {
	err := dao.DeletePost(&post)
	return err
}

func UpdatePostReaction(userName string, post model.Post) error {
	post.UserName = userName

	// Call DAO

	return nil
}

func AddNewPost(userName string, post model.Post) (*mongo.InsertOneResult, error) {
	post.UserName = userName
	post.PostCreatedDateTime = time.Now()
	post.PostEditedDateTime = time.Now()

	result, err := dao.AddNewPost(&post)
	return result, err
}
