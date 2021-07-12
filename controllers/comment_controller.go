package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/tejas-p-shah/Wall-E/model"
	"github.com/tejas-p-shah/Wall-E/services"
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

	var comment model.Comment

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.AddNewComment(claims.UserName, comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
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

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.UpdateComment(claims.UserName, comment)
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

	var comment model.Comment

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.DeleteComment(claims.UserName, comment)
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

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	var comment model.Comment

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = services.UpdateCommentReaction(claims.UserName, comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
