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

  t.ExecuteTemplate(w, "registration", nil)
}

func authorization(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/authorization.html", "templates/header_log.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  t.ExecuteTemplate(w, "authorization", nil)
}

func log_in(w http.ResponseWriter, r *http.Request) {
  email := r.FormValue("email")
  password := r.FormValue("password")

  if user_verification(email, password) {
//сессию создаем
  } else {
    // error := "Такого пользователя не существует проверьте свой логин и пароль"
  }

}

func log_up(w http.ResponseWriter, r *http.Request) {
  email := r.FormValue("mail")
  password := r.FormValue("password")
  name := r.FormValue("name")
  surname := r.FormValue("surname")
  repeat_password := r.FormValue("repeat_password")
  sex := r.FormValue("sex")
  birthday := r.FormValue("birthday")
  key := r.FormValue("key")

  error := data_validation(email, password, repeat_password, name, surname, sex, birthday, key)

  if error == "ok" {
    // почемууууууууууу меил не добавляется я не понимаююююю
    create_user(email, password, name, surname, sex, birthday)
    //вынести вход в отдельный метод и вызывать тут после создания тоже
    http.Redirect(w, r, "/", 301)
  } else {
    t, _ := template.ParseFiles("templates/registration.html", "templates/header_log.html", "templates/form1.html", "templates/form2.html", "templates/form3.html" )
    t.ExecuteTemplate(w, "registration", error)
  }
}

func contacts(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/contacts.html", "templates/header.html", "templates/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }

  t.ExecuteTemplate(w, "profile", nil)
}
