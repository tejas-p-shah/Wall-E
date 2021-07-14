package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tejas-p-shah/Wall-E/authenticators"
	"github.com/tejas-p-shah/Wall-E/controllers"
	"github.com/tejas-p-shah/Wall-E/services"

	"github.com/gorilla/mux"
)

func setRouters() {
	router := mux.NewRouter()
	// Index Route
	router.HandleFunc("/", controllers.IndexHandler).Methods("GET")
	router.HandleFunc("/wall/{wall_id}", controllers.HomePageHandler).Methods("GET")
	router.HandleFunc("/logout", controllers.LogOutHandler).Methods("GET")
	router.HandleFunc("/test/wall/{wall_id}/post/{post_id}", controllers.TestHandler).Methods("GET")
	router.HandleFunc("/test", controllers.TestHandler).Methods("POST")
	router.HandleFunc("/user/email", controllers.TestHandler).Methods("POST")

	// User Routers
	router.HandleFunc("/user/{user_id}", controllers.GetUserProfile).Methods("GET")
	router.HandleFunc("/search/{user_id}", controllers.SearchUser).Methods("GET")

	//Post Routers
	router.HandleFunc("/wall/{wall_id}/post", controllers.AddNewPost).Methods("POST")
	router.HandleFunc("/wall/{wall_id}/post/{post_id}", controllers.UpdatePost).Methods("PUT")
	router.HandleFunc("/wall/{wall_id}/post/{post_id}", controllers.DeletePost).Methods("DELETE")
	router.HandleFunc("/wall/{wall_id}/post/{post_id}/{reaction}", controllers.UpdatePostReaction).Methods("PUT")

	//Comment Routers
	router.HandleFunc("/wall/{wall_id}/post/{post_id}/comment", controllers.AddNewComment).Methods("POST")
	router.HandleFunc("/wall/{wall_id}/post/{post_id}/comment/{comment_id}", controllers.UpdateComment).Methods("PUT")
	router.HandleFunc("/wall/{wall_id}/post/{post_id}/comment/{comment_id}", controllers.DeleteComment).Methods("DELETE")
	router.HandleFunc("/wall/{wall_id}/post/{post_id}/comment/{comment_id}/{reaction}", controllers.UpdateCommentReaction).Methods("PUT")

	http.Handle("/", router)

}

func BootApplication() {

	port := services.GetPort()

	setRouters()
	authenticators.Init_Github_Authenticator()

	fmt.Println("[ Server Running on http://localhost" + port + "]")
	log.Panic(
		http.ListenAndServe(port, nil),
	)
}
