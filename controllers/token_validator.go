package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/tejas-p-shah/Wall-E/model"
)

var jwtKey = []byte("secret_key")

func setToken(w http.ResponseWriter, r *http.Request, email string, username string) {
	expirationTime := time.Now().Add(time.Minute * 30) // Set expiration time to 5 mins
	claims := &model.Claims{
		UserEmail: email,
		UserName:  username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(
		w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
			Path:    "/",
		})
}

func validate_token(w http.ResponseWriter, r *http.Request) (bool, model.Claims) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return false, model.Claims{}
		}
		w.WriteHeader(http.StatusBadRequest)
		return false, model.Claims{}
	}

	tokenString := cookie.Value
	claims := &model.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return false, model.Claims{}
		}
		w.WriteHeader(http.StatusBadRequest)
		return false, model.Claims{}
	}

	if !token.Valid {
		w.WriteHeader(http.StatusBadRequest)
		return false, model.Claims{}
	}

	refreshToken(w, claims)
	return true, *claims
}

func refreshToken(w http.ResponseWriter, claims *model.Claims) {
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		return
	}

	expirationTime := time.Now().Add(time.Minute * 30)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
			Path:    "/",
		})
}
