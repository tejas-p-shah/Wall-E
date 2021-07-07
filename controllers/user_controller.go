package controllers

import (
	"github.com/tejas-p-shah/Wall-E/services"

	"encoding/json"
	"net/http"
	"strconv"
)

func GetUserProfile(responseWriter http.ResponseWriter, request *http.Request) {

	userID, err := strconv.ParseInt(request.URL.Query().Get("user_id"), 10, 64)

	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte("Unable to fetch user_id"))
		return
	}

	user, err := services.GetUser(userID)

	if err != nil {
		responseWriter.WriteHeader(http.StatusNotFound)
		responseWriter.Write([]byte(err.Error()))
	}

	jsonValue, _ := json.Marshal(user)
	responseWriter.Write(jsonValue)
}
