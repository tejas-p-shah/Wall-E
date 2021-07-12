package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tejas-p-shah/Wall-E/dao"
	"github.com/tejas-p-shah/Wall-E/model"
	"github.com/tejas-p-shah/Wall-E/services"
)

type IndexPage struct {
	Title string
}

type GithubDataTokenResp struct {
	UserFullName string `json:"name"`
	UserEmail    string `json:"email"`
	UserName     string `json:"login"`
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tokenStatus, _ := validate_token(w, r)

	if !tokenStatus {
		data := IndexPage{Title: "Wall-E Home"}
		t := template.Must(template.ParseFiles("views/templates/index.gohtml"))
		t.Execute(w, data)
		return
	}

	HomePageHandler(w, r)
}

func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(
		w,
		&http.Cookie{
			Name:  "token",
			Value: "",
			Path:  "/",
		})
}

func LoggedinHandler(w http.ResponseWriter, r *http.Request, githubData string) {
	if githubData == "" {
		fmt.Fprintf(w, "UNAUTHORIZED!!")
		return
	}

	var githubDataTokenResp GithubDataTokenResp
	err := json.Unmarshal([]byte(githubData), &githubDataTokenResp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, foundStatus, _ := services.GetUser(githubDataTokenResp.UserEmail)

	if !foundStatus {

		services.AddUser(model.User{UserFullName: githubDataTokenResp.UserFullName, UserEmail: githubDataTokenResp.UserEmail, UserName: githubDataTokenResp.UserName})
	}

	setToken(w, r, githubDataTokenResp.UserEmail, githubDataTokenResp.UserName)

	redirectURL := "/home"
	http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {

	tokenStatus, claims := validate_token(w, r)

	if !tokenStatus {
		redirectURL := "/"
		http.Redirect(w, r, redirectURL, http.StatusMovedPermanently)
		return
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

	// posts, _ := services.GetUserPosts(user.UserName)

	type Data struct {
		User     *model.User
		Posts    *model.Post
		Comments *model.Comment
	}

	data := &Data{User: user}

	t := template.Must(template.ParseFiles("views/templates/home.gohtml"))
	t.Execute(w, data)

}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("TEST HANDLER : ")
	// w.Header().Set("Content-Type", "application/json")

	// params := mux.Vars(r)
	// fmt.Println("Wall : ", params["wall_id"])
	// fmt.Println("Post : ", params["post_id"])
	// var post model.Post
	// _ = json.NewDecoder(r.Body).Decode(&post)
	// json.NewEncoder(w).Encode(post)

	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["wall_id"]

	// posts, err := services.GetUserPosts(id)
	result, err := dao.GetUserPosts(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(result)

	// w.Header().Set("content-type", "application/json")
	// var post model.Post
	// _ = json.NewDecoder(r.Body).Decode(&post)
	// post.PostDateTime = time.Now()

	// client, err := config.GetMongoClient()
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }

	// collection := client.Database(config.DB).Collection(config.POSTS)
	// result, err := collection.InsertOne(context.TODO(), post)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
	// 	return
	// }
	// json.NewEncoder(w).Encode(result)
}
