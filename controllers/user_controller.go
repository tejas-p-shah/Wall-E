package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/tejas-p-shah/Wall-E/model"
	"github.com/tejas-p-shah/Wall-E/services"
)

func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	tokenStatus, claims := validate_token(w, r)

	if !tokenStatus {
		redirectURL := "/"
		http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
	}

	if claims.UserEmail == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Unable to fetch user email"))
		return
	}

	user, _, err := services.GetUser(claims.UserEmail)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		log.Fatal((err.Error()))
		return
	}

	t := template.Must(template.ParseFiles("views/templates/home.gohtml"))
	t.Execute(w, user)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	tokenStatus, _ := validate_token(w, r)

	if !tokenStatus {
		redirectURL := "/"
		http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
		return
	}

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	type searchStruct struct {
		UserEmail string `json:"user_name"`
	}
	var userStruct searchStruct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&userStruct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, _, err := services.GetUser(userStruct.UserEmail)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User Not Fond"))
		return
	}

	type Data struct {
		User     *model.User
		Posts    *model.Post
		Comments *model.Comment
	}

	data := &Data{User: user}

	t := template.Must(template.ParseFiles("views/templates/home.gohtml"))
	t.Execute(w, data)
}
