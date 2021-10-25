package controllers

import (
	"html/template"
	"net/http"

	"knocker/models"
	"knocker/services/sessions"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

const (
	COOKIE_NAME = "sessionId"
)

var InMemorySession *sessions.Session

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("").Funcs(template.FuncMap{"mod": func(i, j int) bool { return (i+1)%j == 0 }}).ParseFiles(
		// "views/index.html",
		// "views/shared/header.html",
		// "views/shared/footer.html",
		// "views/cards/user_card.html",
		// "views/cards/header_not_authorized.html",
		// "views/shared/header_authorized.html")
		"views/index.html")

	users := models.Get_users()
	session := get_session(r)

	m := map[string]interface{}{
		"Users":   users,
		"Session": session,
	}

	t.ExecuteTemplate(w, "index", m)
}

func Profile(w http.ResponseWriter, r *http.Request) {
	session := get_session(r)

	if session == "" {
		http.Redirect(w, r, "/registration", 301)
	}

	vars := mux.Vars(r)

	t, _ := template.ParseFiles("views/profile.html", "views/shared/header.html",
		"views/shared/footer.html", "views/shared/header_not_authorized.html",
		"views/shared/header_authorized.html")

	user := models.Get_user(vars["id"])

	m := map[string]interface{}{
		"Session": session,
		"User":    user,
	}

	t.ExecuteTemplate(w, "profile", m)
}

func My_page(w http.ResponseWriter, r *http.Request) {
	email := get_session(r)

	if email == "" {
		http.Redirect(w, r, "/registration", 301)
	}

	t, _ := template.ParseFiles("views/my_page.html", "views/shared/header.html",
		"views/shared/footer.html", "views/shared/header_not_authorized.html",
		"views/shared/header_authorized.html")

	user := models.Get_user_by_email(email)

	m := map[string]interface{}{
		"User":    user,
		"Session": email,
	}

	t.ExecuteTemplate(w, "my_page", m)
}

func Contacts(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/contacts.html", "views/shared/header.html",
		"views/shared/footer.html", "views/shared/header_not_authorized.html",
		"views/shared/header_authorized.html")

	t.ExecuteTemplate(w, "contacts", nil)
}
