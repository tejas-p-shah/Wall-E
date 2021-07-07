package model

type User struct {
	UserID       int64  `json:"user_id"`
	UserFullName string `json:"user_fullname"`
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserBio      string `json:"user_bio"`
}

type Credentials struct {
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`
}
