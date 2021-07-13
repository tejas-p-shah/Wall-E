package dao

import (
	"context"
	"fmt"

	"github.com/tejas-p-shah/Wall-E/config"
	"github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Users = map[string]*model.User{
		"a@gmail.com": {UserFullName: "ABC", UserName: "A", UserEmail: "a@gmail.com", UserBio: "Bio of A"},
		"L@gmail.com": {UserFullName: "LMN", UserName: "L", UserEmail: "L@gmail.com", UserBio: "Bio of L"},
		"X@gmail.com": {UserFullName: "XYZ", UserName: "X", UserEmail: "X@gmail.com", UserBio: "Bio of X"},
	}
)

func GetUser(userEmail string) ([]model.User, bool, error) {
	filter := bson.M{"user_email": userEmail}
	user := []model.User{}
	client, err := config.GetMongoClient()
	if err != nil {
		return nil, false, err
	}
	collection := client.Database(config.DB).Collection(config.USERS)
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return nil, false, findError
	}
	for cur.Next(context.TODO()) {
		t := model.User{}
		err := cur.Decode(&t)
		if err != nil {
			return nil, false, err
		}
		user = append(user, t)
	}
	cur.Close(context.TODO())
	if len(user) == 0 {
		return user, false, fmt.Errorf(mongo.ErrNoDocuments.Error())
	}
	return user, true, nil
}

func AddUser(user model.User) (*mongo.InsertOneResult, error) {
	if user.UserEmail == "" {
		return nil, fmt.Errorf("user email address cannot be empty")
	}

	client, err := config.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(config.DB).Collection(config.USERS)
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
