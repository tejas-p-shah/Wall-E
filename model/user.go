package model

import (
	"github.com/golang-jwt/jwt"
)

type User struct {
	UserFullName string `json:"user_fullname,omitempty" bson:"user_fullname,omitempty"`
	UserName     string `json:"user_name,omitempty" bson:"user_name,omitempty"`
	UserEmail    string `json:"user_email,omitempty" bson:"user_email,omitempty"`
	UserBio      string `json:"user_bio,omitempty" bson:"user_bio,omitempty"`
}

type Claims struct {
	UserEmail string `json:"user_email,omitempty" bson:"user_email,omitempty"`
	UserName  string `json:"user_name,omitempty" bson:"user_name,omitempty"`
	jwt.StandardClaims
}
