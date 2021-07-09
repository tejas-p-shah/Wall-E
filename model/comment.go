package model

import "time"

type Comment struct {
	CommentID       int64     `json:"comment_id"`
	CommentParentID int64     `json:"comment_parent_id"`
	PostID          int64     `json:"post_id"`
	UserID          int64     `json:"user_id"`
	UserName        string    `json:"user_name"`
	CommentContent  string    `json:"comment_content"`
	CommentReaction string    `json:"comment_reaction"`
	CommentDateTime time.Time `json:"comment_datetime"`
	Subcomment      []string
}
