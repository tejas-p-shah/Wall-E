package dao

import (
	"context"

	"github.com/tejas-p-shah/Wall-E/config"
	"github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCommentByKey(key string, value primitive.ObjectID) ([]model.Comment, error) {

	filter := bson.M{key: value}
	comment := []model.Comment{}
	client, err := config.GetMongoClient()
	if err != nil {
		return comment, err
	}
	collection := client.Database(config.DB).Collection(config.COMMENTS)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return comment, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := model.Comment{}
		err := cur.Decode(&t)
		if err != nil {
			return comment, err
		}
		comment = append(comment, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(comment) == 0 {
		return comment, mongo.ErrNoDocuments
	}
	return comment, nil
}

func GetPostComments(PostID primitive.ObjectID) ([]model.Comment, error) {
	filter := bson.M{"post_id": PostID}
	comments := []model.Comment{}
	client, err := config.GetMongoClient()
	if err != nil {
		return comments, err
	}
	collection := client.Database(config.DB).Collection(config.COMMENTS)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return comments, findError
	}
	for cur.Next(context.TODO()) {
		t := model.Comment{}
		err := cur.Decode(&t)
		if err != nil {
			return comments, err
		}
		t.CommentReactionCount = len(t.CommentReactionList)
		comments = append(comments, t)
	}
	cur.Close(context.TODO())
	if len(comments) == 0 {
		return comments, mongo.ErrNoDocuments
	}
	return comments, nil
}

func AddNewComment(comment *model.Comment) (*mongo.InsertOneResult, error) {

	client, err := config.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(config.DB).Collection(config.COMMENTS)
	result, err := collection.InsertOne(context.TODO(), comment)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateComment(comment *model.Comment) error {

	filter := bson.M{"$set": bson.M{"comment_content": comment.CommentContent}}

	client, err := config.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(config.DB).Collection(config.COMMENTS)
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": comment.CommentID}, filter)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(commentID primitive.ObjectID) error {

	client, err := config.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(config.DB).Collection(config.COMMENTS)
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": commentID})
	if err != nil {
		return err
	}
	return nil
}

func UpdateCommentReaction(userName string, commentID primitive.ObjectID, reactionValue int) error {

	comments, _ := GetCommentByKey("_id", commentID)

	if reactionValue != 0 {
		for _, ele := range comments[0].CommentReactionList {
			if ele == userName {
				return nil
			}
		}
		comments[0].CommentReactionList = append(comments[0].CommentReactionList, userName)
	} else {
		for i, v := range comments[0].CommentReactionList {
			if v == userName {
				comments[0].CommentReactionList = append(comments[0].CommentReactionList[:i], comments[0].CommentReactionList[i+1:]...)
				break
			}
		}
	}
	filter := bson.M{"$set": bson.M{"comment_reaction": comments[0].CommentReactionList}}

	client, err := config.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(config.DB).Collection(config.COMMENTS)
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": commentID}, filter)
	if err != nil {
		return err
	}
	return nil
}
