package controllers

import (
	"encoding/json"
	"html/template"
	"knocker/models"
	"net/http"
)

func Favorite(w http.ResponseWriter, r *http.Request) {
	session := get_session(r)
	liked_users := models.Get_liked_users(models.Get_user_by_email(session).Id)

	if session == "" {
		http.Redirect(w, r, "/registration", 301)
	}

	t, _ := template.ParseFiles("views/favorite.html", "views/shared/header.html",
		"views/shared/footer.html", "views/shared/header_not_authorized.html",
		"views/shared/header_authorized.html", "views/cards/long_card.html")

	m := map[string]interface{}{
		"Session": session,
		"Users":   liked_users,
	}

	t.ExecuteTemplate(w, "favorite", m)
}

//	здесь просто вызывается метод ад_ту_лайкд, строк много потому что создается видимость обработки ошибок
func Like(w http.ResponseWriter, r *http.Request) {
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
		cucurrent_user_id := models.Get_user_by_email(get_session(r)).Id
		models.Add_to_liked(cucurrent_user_id, liked_user_id)
		js, err := json.Marshal("OK")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	}
}
