package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/websocket"
)

var inMemorySession *Session
var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
var upgrader = websocket.Upgrader{}          // Configure the upgrader

func handleFunc() {
  router := mux.NewRouter()

  inMemorySession = NewSession()

  router.HandleFunc("/", index).Methods("GET")
  router.HandleFunc("/profile/{id:[0-9]+}", profile).Methods("GET")
  router.HandleFunc("/registration", registration).Methods("GET")
  router.HandleFunc("/authorization/", authorization).Methods("GET")
  router.HandleFunc("/log_up", log_up).Methods("POST")
  router.HandleFunc("/log_in", log_in).Methods("POST")
  router.HandleFunc("/contacts/", contacts).Methods("GET")
  router.HandleFunc("/chat/", chat).Methods("GET")

  router.HandleFunc("/ws", handleConnections)  // websocket
  go handleMessages()                          // запускаем горутину

  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
  http.Handle("/", router)


  http.ListenAndServe(":8080", nil)
}

func main()  {
  handleFunc()
}
