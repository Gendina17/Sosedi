package main

import (
  "net/http"
)

func handleFunc() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
  http.HandleFunc("/", index)
  http.HandleFunc("/profile/", profile)
  http.HandleFunc("/registration/", registration)
  http.HandleFunc("/authorization/", authorization)
  http.HandleFunc("/contacts/", contacts)
  http.ListenAndServe(":8080", nil)
}

func main()  {
  handleFunc()
}
