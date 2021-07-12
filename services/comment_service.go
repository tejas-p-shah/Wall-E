package services

import (
	"fmt"

	"github.com/tejas-p-shah/Wall-E/dao"
	"github.com/tejas-p-shah/Wall-E/model"
)

func GetPostComment(PostID int64) (*model.Comment, error) {
	comments, err := dao.GetPostComments(PostID)

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func UpdateComment(userName string, comment model.Comment) error {

	if comment.UserName != userName {
		return fmt.Errorf("user %s Does not have permissions", userName)
	}

	// Call DAO

	return nil
}

func DeleteComment(userName string, comment model.Comment) error {

	if comment.UserName != userName {
		return fmt.Errorf("user %s Does not have permissions", userName)
	}

	// Call DAO

	return nil
}

func UpdateCommentReaction(userName string, comment model.Comment) error {
	comment.UserName = userName

	// Call DAO

	return nil
}

func AddNewComment(userName string, comment model.Comment) error {
	comment.UserName = userName

	// Call DAO

	return nil
}
