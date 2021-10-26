package main

import (
	"net/http"

	"knocker/controllers"
	"knocker/services/sessions"

	"github.com/gorilla/mux"
)

var inMemorySession *sessions.Session

func handleFunc() {
	router := mux.NewRouter()

	controllers.InMemorySession = sessions.NewSession()
	inMemorySession = controllers.InMemorySession

	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/profile/{id:[0-9]+}", controllers.Profile).Methods("GET")
	router.HandleFunc("/my_page", controllers.My_page).Methods("GET")
	router.HandleFunc("/favorite", controllers.Favorite).Methods("GET")
	router.HandleFunc("/registration", controllers.Registration).Methods("GET")
	router.HandleFunc("/authorization", controllers.Authorization).Methods("GET")
	router.HandleFunc("/log_up", controllers.Log_up).Methods("POST")
	router.HandleFunc("/log_in", controllers.Log_in).Methods("POST")
	router.HandleFunc("/like", controllers.Like).Methods("POST")
	router.HandleFunc("/comment", controllers.Comment).Methods("POST")
	router.HandleFunc("/dislike", controllers.Dislike).Methods("POST")
	router.HandleFunc("/check_session", controllers.Check_session).Methods("POST")
	router.HandleFunc("/log_out", controllers.Log_out).Methods("GET")
	router.HandleFunc("/contacts", controllers.Contacts).Methods("GET")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}

func main() {
	handleFunc()
}
