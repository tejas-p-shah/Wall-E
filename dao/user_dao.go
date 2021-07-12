package dao

import (
	"context"
	"fmt"

	"github.com/tejas-p-shah/Wall-E/config"
	"github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Users = map[string]*model.User{
		"a@gmail.com": {UserFullName: "ABC", UserName: "A", UserEmail: "a@gmail.com", UserBio: "Bio of A"},
		"L@gmail.com": {UserFullName: "LMN", UserName: "L", UserEmail: "L@gmail.com", UserBio: "Bio of L"},
		"X@gmail.com": {UserFullName: "XYZ", UserName: "X", UserEmail: "X@gmail.com", UserBio: "Bio of X"},
	}
)

func GetUser(userEmail string) (*model.User, bool, error) {
	if user := Users[userEmail]; user != nil {
		return user, true, nil
	}

	return nil, false, fmt.Errorf("user %v was not found", userEmail)
}

func AddUser(user model.User) (*mongo.InsertOneResult, error) {
	if user.UserEmail == "" {
		return nil, fmt.Errorf("user email address cannot be empty")
	}

	client, err := config.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(config.DB).Collection(config.POSTS)
	result, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return result, nil
}
