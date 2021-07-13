package dao

import (
	"context"
	"fmt"

	"github.com/tejas-p-shah/Wall-E/config"
	"github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPostByID(postID primitive.ObjectID) ([]model.Post, error) {

	filter := bson.M{"_id": postID}
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

func GetUserPosts(wallID string) ([]model.Post, error) {
	fmt.Println(wallID)
	filter := bson.M{"wall_user_name": wallID}
	posts := []model.Post{}
	client, err := config.GetMongoClient()
	if err != nil {
		return posts, err
	}
	collection := client.Database(config.DB).Collection(config.POSTS)
	cur, findError := collection.Find(context.TODO(), filter)
	// fmt.Println("Find ERR : ", findError.Error())
	if findError != nil {
		return posts, findError
	}
	//Map result to slice
	for cur.Next(context.TODO()) {
		t := model.Post{}
		err := cur.Decode(&t)
		// fmt.Println("ERR : ", err.Error())
		if err != nil {
			// fmt.Errorf(err.Error())
			return posts, err
		}
		posts = append(posts, t)
	}
	// once exhausted, close the cursor
	cur.Close(context.TODO())
	if len(posts) == 0 {
		fmt.Println(len(posts))
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

func DeletePost(postID primitive.ObjectID) error {

	client, err := config.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(config.DB).Collection(config.POSTS)
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": postID})
	if err != nil {
		return err
	}
	return nil
}

func UpdatePostReaction(userName string, postID primitive.ObjectID, reactionValue int) error {

	posts, err := GetPostByID(postID)

	if reactionValue != 0 {
		for _, ele := range posts[0].PostLikeList {
			if ele == userName {
				break
			}
		}
		posts[0].PostLikeList = append(posts[0].PostLikeList, userName)
	} else {
		for i, v := range posts[0].PostLikeList {
			if v == userName {
				posts[0].PostLikeList = append(posts[0].PostLikeList[:i], posts[0].PostLikeList[i+1:]...)
				break
			}
		}
	}
	// posts[0].PostLikeList
	filter := bson.M{"$set": bson.M{"post_reaction": posts[0].PostLikeList}}

	client, err := config.GetMongoClient()
	if err != nil {
		return err
	}
	collection := client.Database(config.DB).Collection(config.POSTS)
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": postID}, filter)
	if err != nil {
		return err
	}
	return nil
}
