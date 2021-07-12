package services

import (
	"github.com/tejas-p-shah/Wall-E/dao"
	"github.com/tejas-p-shah/Wall-E/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(userEmail string) (*model.User, bool, error) {
	user, foundStatus, err := dao.GetUser(userEmail)

	if err != nil {
		return nil, foundStatus, err
	}

	return user, foundStatus, nil
}

func AddUser(user model.User) (*mongo.InsertOneResult, error) {
	result, err := dao.AddUser(user)
	return result, err
}

func UpdateUserBio(username string, bio string) {

}
