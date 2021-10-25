package controllers

import (
	"html/template"
	"knocker/models"
	"knocker/services/images_uploading"
	"net/http"
)

func Registration(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/registration.html", "views/header_log.html",
		"views/form1.html", "views/form2.html", "views/form3.html")

	t.ExecuteTemplate(w, "registration", nil)
}

func Authorization(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/authorization.html", "views/header_log.html")

	t.ExecuteTemplate(w, "authorization", nil)
}

func Log_in(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if models.User_verification(email, password) {
		create_session(email, &w)
		http.Redirect(w, r, "/my_page", 301)
	} else {
		error := "Такого пользователя не существует проверьте свой логин и пароль"
		t, _ := template.ParseFiles("views/authorization.html", "views/header_log.html")
		t.ExecuteTemplate(w, "authorization", error)
	}
}

func Log_up(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/registration.html", "views/header_log.html",
		"views/form1.html", "views/form2.html", "views/form3.html")

	email := r.FormValue("mail")
	password := r.FormValue("password")
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	repeat_password := r.FormValue("repeat_password")
	sex := r.FormValue("sex")
	birthday := r.FormValue("birthday")

	photo_key, error := images_uploading.Getting_image_from_request(r)

	if error != "ok" {
		t.ExecuteTemplate(w, "registration", error)
	}

	error = models.Data_validation(email, password, repeat_password, name, surname, sex, birthday)

	if error == "ok" {
		models.Create_user(email, password, name, surname, sex, birthday, photo_key)
		create_session(email, &w)
		http.Redirect(w, r, "/", 301)
	} else {
		t.ExecuteTemplate(w, "registration", error)
	}
}

func Log_out(w http.ResponseWriter, r *http.Request) {
	//сделать выход - очистку куки

	http.Redirect(w, r, "/", 301)
}
