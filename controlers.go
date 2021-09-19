package main

import (
  "fmt"
  "net/http"
  "html/template"
  "github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request) {
  t, _ := template.ParseFiles("templates/index.html", "templates/header.html",
    "templates/footer.html", "templates/user_card.html")

  users := get_users()

  t.ExecuteTemplate(w, "index", users)
}

func profile(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  t, _ := template.ParseFiles("templates/profile.html", "templates/header.html",
     "templates/footer.html")

  user := get_user(vars["id"])

  t.ExecuteTemplate(w, "profile", user)
}

func registration(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/registration.html", "templates/header_log.html", "templates/form1.html", "templates/form2.html", "templates/form3.html" )

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  t.ExecuteTemplate(w, "profile", nil)
}

func authorization(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/authorization.html", "templates/header_log.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  t.ExecuteTemplate(w, "profile", nil)
}

func contacts(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/contacts.html", "templates/header.html", "templates/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  t.ExecuteTemplate(w, "profile", nil)
}
