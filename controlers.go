package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

const (
	COOKIE_NAME = "sessionId"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("").Funcs(template.FuncMap{"mod": func(i, j int) bool { return (i+1)%j == 0 }}).ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html",
		"templates/user_card.html",
		"templates/header_not_authorized.html",
		"templates/header_authorized.html")

	users := get_users()
	session := get_session(r)

	m := map[string]interface{}{
		"Users":   users,
		"Session": session,
	}

	t.ExecuteTemplate(w, "index", m)
}

func profile(w http.ResponseWriter, r *http.Request) {
	session := get_session(r)

	if session == "" {
		http.Redirect(w, r, "/registration", 301)
	}

	vars := mux.Vars(r)

	t, _ := template.ParseFiles("templates/profile.html", "templates/header.html",
		"templates/footer.html", "templates/header_not_authorized.html",
		"templates/header_authorized.html")

	user := get_user(vars["id"])

	m := map[string]interface{}{
		"Session": session,
		"User":    user,
	}

	t.ExecuteTemplate(w, "profile", m)
}

func my_page(w http.ResponseWriter, r *http.Request) {
	email := get_session(r)

	if email == "" {
		http.Redirect(w, r, "/registration", 301)
	}

	t, _ := template.ParseFiles("templates/my_page.html", "templates/header.html",
		"templates/footer.html", "templates/header_not_authorized.html",
		"templates/header_authorized.html")

	user := get_user_by_email(email)

	m := map[string]interface{}{
		"User":    user,
		"Session": email,
	}

	t.ExecuteTemplate(w, "my_page", m)
}

func favorite(w http.ResponseWriter, r *http.Request) {
	session := get_session(r)
	liked_users := get_liked_users(get_user_by_email(session).Id)

	if session == "" {
		http.Redirect(w, r, "/registration", 301)
	}

	t, _ := template.ParseFiles("templates/favorite.html", "templates/header.html",
		"templates/footer.html", "templates/header_not_authorized.html",
		"templates/header_authorized.html", "templates/long_card.html")

	m := map[string]interface{}{
		"Session": session,
		"Users":   liked_users,
	}

	t.ExecuteTemplate(w, "favorite", m)
}

func registration(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/registration.html", "templates/header_log.html",
		"templates/form1.html", "templates/form2.html", "templates/form3.html")

	t.ExecuteTemplate(w, "registration", nil)
}

func authorization(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/authorization.html", "templates/header_log.html")

	t.ExecuteTemplate(w, "authorization", nil)
}

func log_in(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if user_verification(email, password) {
		create_session(email, &w)
		http.Redirect(w, r, "/my_page", 301)
	} else {
		error := "Такого пользователя не существует проверьте свой логин и пароль"
		t, _ := template.ParseFiles("templates/authorization.html", "templates/header_log.html")
		t.ExecuteTemplate(w, "authorization", error)
	}
}

func log_up(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/registration.html", "templates/header_log.html",
		"templates/form1.html", "templates/form2.html", "templates/form3.html")

	email := r.FormValue("mail")
	password := r.FormValue("password")
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	repeat_password := r.FormValue("repeat_password")
	sex := r.FormValue("sex")
	birthday := r.FormValue("birthday")

	photo_key, error := getting_image_from_request(r)

	if error != "ok" {
		t.ExecuteTemplate(w, "registration", error)
	}

	error = data_validation(email, password, repeat_password, name, surname, sex, birthday)

	if error == "ok" {
		create_user(email, password, name, surname, sex, birthday, photo_key)
		create_session(email, &w)
		http.Redirect(w, r, "/", 301)
	} else {
		t.ExecuteTemplate(w, "registration", error)
	}
}

func log_out(w http.ResponseWriter, r *http.Request) {
	//сделать выход - очистку куки

	http.Redirect(w, r, "/", 301)
}

//	здесь просто вызывается метод ад_ту_лайкд, строк много потому что создается видимость обработки ошибок
func like(w http.ResponseWriter, r *http.Request) {
	var liked_user_id uint16
	err := json.NewDecoder(r.Body).Decode(&liked_user_id)
	if err != nil {
		println("An error occured")
		js, err := json.Marshal("Error")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	} else {
		cucurrent_user_id := get_user_by_email(get_session(r)).Id
		add_to_liked(cucurrent_user_id, liked_user_id)
		js, err := json.Marshal("OK")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	}
}

//	просто метод который использует твой гет_сессион и возвращает либо инэктив либо список айди тех кого лайкнул авторизованный
func check_session(w http.ResponseWriter, r *http.Request) {
	current_session := get_session(r)
	var js []byte
	var err error
	if current_session != "" {
		js, err = json.Marshal(get_liked_users_id(get_user_by_email(current_session).Id))
	} else {
		js, err = json.Marshal("Inactive")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func contacts(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/contacts.html", "templates/header.html",
		"templates/footer.html", "templates/header_not_authorized.html",
		"templates/header_authorized.html")

	t.ExecuteTemplate(w, "contacts", nil)
}

func create_session(email string, w *http.ResponseWriter) {
	sessionId := inMemorySession.Init(email)

	cookie := &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(44640 * time.Minute),
	}

	http.SetCookie(*w, cookie)
}

func get_session(r *http.Request) string {
	cookie, _ := r.Cookie(COOKIE_NAME)

	if cookie != nil {
		return inMemorySession.Get(cookie.Value)
	}
	return ""
}
