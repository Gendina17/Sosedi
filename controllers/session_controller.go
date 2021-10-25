package controllers

import (
	"encoding/json"
	"knocker/models"
	"net/http"
	"time"
)

func create_session(email string, w *http.ResponseWriter) {
	sessionId := InMemorySession.Init(email)

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
		return InMemorySession.Get(cookie.Value)
	}
	return ""
}

//	просто метод который использует твой гет_сессион и возвращает либо инэктив либо список айди тех кого лайкнул авторизованный
func Check_session(w http.ResponseWriter, r *http.Request) {
	current_session := get_session(r)
	var js []byte
	var err error
	if current_session != "" {
		js, err = json.Marshal(models.Get_liked_users_id(models.Get_user_by_email(current_session).Id))
	} else {
		js, err = json.Marshal("Inactive")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
