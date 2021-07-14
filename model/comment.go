package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	CommentID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CommentParentID        primitive.ObjectID `json:"comment_parent_id,omitempty" bson:"comment_parent_id,omitempty"`
	PostID                 primitive.ObjectID `json:"post_id,omitempty" bson:"post_id,omitempty"`
	WallUserName           string             `json:"wall_user_name,omitempty" bson:"wall_user_name,omitempty"`
	UserName               string             `json:"user_name,omitempty" bson:"user_name,omitempty"`
	CommentContent         string             `json:"comment_content,omitempty" bson:"comment_content,omitempty"`
	CommentReactionList    []string           `json:"comment_reaction,omitempty" bson:"comment_reaction,omitempty"`
	CommentReactionCount   int                `json:"comment_reaction_count,omitempty" bson:"comment_reaction_count,omitempty"`
	CommentCreatedDateTime time.Time          `json:"comment_datetime,omitempty" bson:"comment_datetime,omitempty"`
	CommentEditDateTime    time.Time          `json:"comment_edit_datetime,omitempty" bson:"comment_edit_datetime,omitempty"`
}
