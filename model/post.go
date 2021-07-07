package model

import "time"

type Post struct {
	PostID        int64     `json:"post_id"`
	UserID        int64     `json:"user_id"`
	UserName      string    `json:"user_name"`
	PostContent   string    `json:"post_content"`
	PostReactions string    `json:"post_reaction"`
	PostDateTime  time.Time `json:"post_datetime"`
}
