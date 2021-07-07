package dao

import (
	"fmt"

	"github.com/tejas-p-shah/Wall-E/model"
)

func GetPostComments(PostID int64) (*model.Comment, error) {

	return nil, fmt.Errorf("user %v was not found", PostID)
}
