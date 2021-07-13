package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tejas-p-shah/Wall-E/model"
	"github.com/tejas-p-shah/Wall-E/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddNewPost(w http.ResponseWriter, r *http.Request) {
	tokenStatus, claims := validate_token(w, r)
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

	params := mux.Vars(r)
	var post model.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.WallUserName = params["wall_id"]

	result, err := services.AddNewPost(claims.UserName, post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	tokenStatus, claims := validate_token(w, r)

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

	// params := mux.Vars(r)
	// fmt.Println("Wall : ", params["wall_id"])
	// fmt.Println("Post : ", params["post_id"])
	var post model.Post
	_ = json.NewDecoder(r.Body).Decode(&post)

	if post.UserName != claims.UserName {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err := services.UpdatePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	tokenStatus, claims := validate_token(w, r)

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

	params := mux.Vars(r)

	var post model.Post

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if post.UserName != claims.UserName {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	objectID, err := primitive.ObjectIDFromHex(params["post_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = services.DeletePost(objectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UpdatePostReaction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("1")
	tokenStatus, claims := validate_token(w, r)
	fmt.Println("2")
	if !tokenStatus {
		redirectURL := "/"
		http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
		return
	}
	fmt.Println("3")
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}
	fmt.Println("4")
	// var post model.Post

	// decoder := json.NewDecoder(r.Body)
	// err := decoder.Decode(&post)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	params := mux.Vars(r)
	objectID, err := primitive.ObjectIDFromHex(params["post_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("5")
	value, ok := params["reaction"]
	if ok {
		reactionValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println("6")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = services.UpdatePostReaction(claims.UserName, objectID, reactionValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Println("bad")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}
