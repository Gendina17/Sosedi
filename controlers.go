package main

import (
  "fmt"
  "net/http"
  "html/template"
  "github.com/gorilla/mux"
  "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html", "templates/user_card.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  var users = []User{}

  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/sosedi")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  defer db.Close()

  res, err := db.Query("SELECT id, name, price_min, price_max FROM users ")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

   for res.Next() {
     var user User
     err = res.Scan(&user.Id, &user.Name, &user.PriceMin, &user.PriceMax)
     if err != nil {
       fmt.Fprintf(w, err.Error())
     }

      users = append(users, user)
   }

  t.ExecuteTemplate(w, "index", users)
}

func profile(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  t, err := template.ParseFiles("templates/profile.html", "templates/header.html", "templates/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/sosedi")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  defer db.Close()

  res, err := db.Query(fmt.Sprintf("SELECT id, name, price_min, price_max FROM users WHERE id = %s ", vars["id"]))

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  var user = User{}
  for res.Next() {
     var new_user User
     err = res.Scan(&new_user.Id, &new_user.Name, &new_user.PriceMin, &new_user.PriceMax)
     if err != nil {
       fmt.Fprintf(w, err.Error())
     }
     user = new_user
  }

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
