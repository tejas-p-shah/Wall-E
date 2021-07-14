package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tejas-p-shah/Wall-E/model"
	"github.com/tejas-p-shah/Wall-E/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddNewComment(w http.ResponseWriter, r *http.Request) {

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
	var comment model.Comment
	_ = json.NewDecoder(r.Body).Decode(&comment)
	comment.WallUserName = params["wall_id"]
	var err error
	comment.PostID, err = primitive.ObjectIDFromHex(params["post_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := services.AddNewComment(claims.UserName, comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func UpdateComment(w http.ResponseWriter, r *http.Request) {

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

	var comment model.Comment
	_ = json.NewDecoder(r.Body).Decode(&comment)

	if comment.UserName != claims.UserName {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	params := mux.Vars(r)
	objectID, err := primitive.ObjectIDFromHex(params["comment_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment.CommentID = objectID
	err = services.UpdateComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
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

	var comment model.Comment

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if comment.UserName != claims.UserName {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	objectID, err := primitive.ObjectIDFromHex(params["comment_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = services.DeleteComment(objectID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func UpdateCommentReaction(w http.ResponseWriter, r *http.Request) {
	tokenStatus, claims := validate_token(w, r)

	if !tokenStatus {
		redirectURL := "/"
		http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
		return
	}

	// headerContentTtype := r.Header.Get("Content-Type")
	// if headerContentTtype != "application/json" {
	// 	w.WriteHeader(http.StatusUnsupportedMediaType)
	// 	return
	// }

	params := mux.Vars(r)
	objectID, err := primitive.ObjectIDFromHex(params["comment_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	value, ok := params["reaction"]
	if ok {
		reactionValue, err := strconv.Atoi(value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = services.UpdateCommentReaction(claims.UserName, objectID, reactionValue)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
