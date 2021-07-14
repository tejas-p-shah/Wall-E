package services

import (
	"time"

	"github.com/tejas-p-shah/Wall-E/dao"
	"github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserPosts(userName string) ([]model.Post, error) {
	posts, err := dao.GetUserPosts(userName)
	// for _, v := range posts {
	// 	v.PostLikeCount = len(v.PostLikeList)
	// 	fmt.Println(v.PostLikeCount)
	// }
	return posts, err
}

func UpdatePost(post model.Post) error {
	post.PostEditedDateTime = time.Now()
	err := dao.UpdatePost(&post)
	return err
}

func DeletePost(postID primitive.ObjectID) error {
	err := dao.DeletePost(postID)
	comments, _ := dao.GetPostComments(postID)
	for _, v := range comments {
		dao.DeleteComment(v.CommentID)
	}
	return err
}

func UpdatePostReaction(userName string, postID primitive.ObjectID, reactionValue int) error {
	err := dao.UpdatePostReaction(userName, postID, reactionValue)
	return err
}

func AddNewPost(userName string, post model.Post) (*mongo.InsertOneResult, error) {
	post.UserName = userName
	post.PostCreatedDateTime = time.Now()
	post.PostEditedDateTime = time.Now()

	result, err := dao.AddNewPost(&post)
	return result, err
}
