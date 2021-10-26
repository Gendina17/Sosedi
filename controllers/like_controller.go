package controllers

import (
	"encoding/json"
	"html/template"
	"knocker/models"
	"net/http"
)

type NewComment struct {
	Comment       string
	Liked_user_id uint16
}

func Favorite(w http.ResponseWriter, r *http.Request) {
	session := get_session(r)
	liked_users := models.Get_liked_users(models.Get_user_by_email(session).Id)

	if session == "" {
		http.Redirect(w, r, "/registration", 301)
	}

	t, _ := template.ParseFiles("views/favorite.html", "views/header.html",
		"views/footer.html", "views/header_not_authorized.html",
		"views/header_authorized.html", "views/long_card.html")

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
		current_user_id := models.Get_user_by_email(get_session(r)).Id
		models.Add_to_liked(current_user_id, liked_user_id)
		js, err := json.Marshal("OK")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(js)
	}
}

//	вызывается метод ад_коммент из модели с параметрами, переданными в пост запросе
func Comment(w http.ResponseWriter, r *http.Request) {
	var new_comment NewComment
	var js []byte
	err := json.NewDecoder(r.Body).Decode(&new_comment)
	if err != nil {
		println("An error occured")
		js, _ = json.Marshal("Error")
	} else {
		current_user_id := models.Get_user_by_email(get_session(r)).Id
		models.Add_comment(new_comment.Comment, new_comment.Liked_user_id, current_user_id)
		js, _ = json.Marshal("OK")
	}
	w.Write(js)
}

//	то же самое что и выше только вызывает метод римув_фром_лайкд
func Dislike(w http.ResponseWriter, r *http.Request) {
	var liked_user_id uint16
	var js []byte
	err := json.NewDecoder(r.Body).Decode(&liked_user_id)
	if err != nil {
		println("An error occured")
		js, _ = json.Marshal("Error")
	} else {
		current_user_id := models.Get_user_by_email(get_session(r)).Id
		models.Remove_from_liked(liked_user_id, current_user_id)
		js, _ = json.Marshal("OK")
	}
	w.Write(js)
}
