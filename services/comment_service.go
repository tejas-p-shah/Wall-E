package services

import (
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
