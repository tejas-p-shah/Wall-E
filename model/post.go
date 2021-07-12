package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	PostID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	WallUserName        string             `json:"wall_user_name,omitempty" bson:"wall_user_name,omitempty"`
	UserName            string             `json:"user_name,omitempty" bson:"user_name,omitempty"`
	PostTitle           string             `json:"post_title,omitempty" bson:"post_title,omitempty"`
	PostContent         string             `json:"post_content,omitempty" bson:"post_content,omitempty"`
	PostLikeList        []string           `json:"post_reaction,omitempty" bson:"post_reaction,omitempty"` // slice of usernames
	PostCreatedDateTime time.Time          `json:"post_created_datetime,omitempty" bson:"post_created_datetime,omitempty"`
	PostEditedDateTime  time.Time          `json:"post_edited_datetime,omitempty" bson:"post_edited_datetime,omitempty"`
}
