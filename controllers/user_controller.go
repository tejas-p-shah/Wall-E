package controllers

import (
	"html/template"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
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

	params := mux.Vars(r)

	user, _, err := services.GetUser(params["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	posts, _ := services.GetUserPosts(user[len(user)-1].UserName)

	var postComments []model.Comment

	for _, v := range posts {
		comments, _ := services.GetPostComment(v.PostID)
		postComments = append(postComments, comments...)
	}

	sort.Slice(postComments, func(i, j int) bool {
		return postComments[i].CommentParentID.String() < postComments[j].CommentParentID.String()
	})

	type Data struct {
		User     *model.User
		Posts    []model.Post
		Comments []model.Comment
	}

	// fmt.Println("new : ", postComments)
	data := &Data{User: &user[len(user)-1], Posts: posts, Comments: postComments}

	t := template.Must(template.ParseFiles("views/templates/wall.gohtml"))
	t.Execute(w, data)
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(data)
}
