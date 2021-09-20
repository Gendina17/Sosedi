package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func handleFunc() {
  router := mux.NewRouter()

  router.HandleFunc("/", index).Methods("GET")
  router.HandleFunc("/profile/{id:[0-9]+}", profile).Methods("GET")
  router.HandleFunc("/registration/", registration).Methods("GET")
  router.HandleFunc("/authorization/", authorization).Methods("GET")
  router.HandleFunc("/log_up/", log_up).Methods("POST")
  router.HandleFunc("/log_in/", log_in).Methods("POST")
  router.HandleFunc("/contacts/", contacts).Methods("GET")

  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
  http.Handle("/", router)

  http.ListenAndServe(":8080", nil)
}

func main()  {
  handleFunc()
}
