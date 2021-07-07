package dao

import (
	"fmt"

	"github.com/tejas-p-shah/Wall-E/model"
)

var (
	Users = map[int64]*model.User{
		101: {UserID: 101, UserFullName: "ABC", UserName: "A", UserEmail: "a@gmail.com", UserBio: "Bio of A"},
	}

	Credentials = map[int64]*model.Credentials{
		101: {UserName: "ABC", UserPassword: "123"},
	}
)

func GetUser(userID int64) (*model.User, error) {
	if user := Users[userID]; user != nil {
		return user, nil
	}

	return nil, fmt.Errorf("user %v was not found", userID)
}
