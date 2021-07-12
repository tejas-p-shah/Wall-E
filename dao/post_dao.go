package dao

import (
	"context"

	"github.com/tejas-p-shah/Wall-E/config"
	"github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserPosts(wallID string) ([]model.Post, error) {

	filter := bson.M{"wall_user_name": wallID}
	posts := []model.Post{}
	client, err := config.GetMongoClient()
	if err != nil {
		return posts, err
	}
	collection := client.Database(config.DB).Collection(config.POSTS)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return posts, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := model.Post{}
		err := cur.Decode(&t)
		if err != nil {
			return posts, err
		}
		posts = append(posts, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(posts) == 0 {
		return posts, mongo.ErrNoDocuments
	}
	return posts, nil
}

func AddNewPost(post *model.Post) (*mongo.InsertOneResult, error) {

	client, err := config.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(config.DB).Collection(config.POSTS)
	result, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdatePost(post *model.Post) error {

	filter := bson.M{"$set": bson.M{"post_title": post.PostTitle, "post_content": post.PostContent}}

	client, err := config.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(config.DB).Collection(config.POSTS)
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": post.PostID}, filter)
	if err != nil {
		return err
	}
	return nil
}

func DeletePost(post *model.Post) error {

	client, err := config.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(config.DB).Collection(config.POSTS)
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": post.PostID})
	if err != nil {
		return err
	}
	return nil
}
