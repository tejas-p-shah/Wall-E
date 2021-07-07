package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tejas-p-shah/Wall-E/controllers"
	"github.com/tejas-p-shah/Wall-E/services"

	"github.com/gorilla/mux"
)

func setRouters() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.IndexHandler).Methods("GET")
	router.HandleFunc("/login/default/", controllers.IndexHandler).Methods("GET")
	router.HandleFunc("/users", nil).Methods("GET")
	router.HandleFunc("/user", controllers.GetUserProfile).Methods("GET")

	http.Handle("/", router)

}

func BootApplication() {

	port := services.GetPort()

	setRouters()
	services.Init_Github_Authenticator()

	fmt.Println("[ Server Running on http://localhost" + port + "]")
	log.Panic(
		http.ListenAndServe(port, nil),
	)
}
