package main

import (
  "fmt"
  "net/http"
  "html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html", "templates/user_card.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  t.ExecuteTemplate(w, "index", nil)
}

func profile(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/profile.html", "templates/header.html", "templates/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  t.ExecuteTemplate(w, "profile", nil)
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
