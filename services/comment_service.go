package services

import (
	"time"

	"github.com/tejas-p-shah/Wall-E/dao"
	"github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPostComment(PostID primitive.ObjectID) ([]model.Comment, error) {
	comments, err := dao.GetPostComments(PostID)
	// for _, v := range comments {
	// 	v.CommentReactionCount = len(v.CommentReactionList)
	// 	fmt.Println(v.CommentReactionCount)
	// }
	return comments, err
}

func UpdateComment(comment model.Comment) error {

	comment.CommentEditDateTime = time.Now()
	err := dao.UpdateComment(&comment)
	return err
}

func DeleteComment(commentID primitive.ObjectID) error {

	err := dao.DeleteComment(commentID)
	comments, _ := dao.GetCommentByKey("comment_parent_id", commentID)
	for _, v := range comments {
		dao.DeleteComment(v.CommentID)
	}
	return err
}

func UpdateCommentReaction(userName string, commentID primitive.ObjectID, reactionValue int) error {
	err := dao.UpdateCommentReaction(userName, commentID, reactionValue)
	return err
}

func AddNewComment(userName string, comment model.Comment) (*mongo.InsertOneResult, error) {
	comment.UserName = userName
	comment.CommentCreatedDateTime = time.Now()
	comment.CommentEditDateTime = time.Now()

	result, err := dao.AddNewComment(&comment)
	return result, err
}
