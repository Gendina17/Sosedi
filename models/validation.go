package models

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func User_verification(email string, password string) bool {
	db := connect_db()
	defer db.Close()

	if !email_exist(email) {
		return false
	}

	res, _ := db.Query(fmt.Sprintf("SELECT password, sault FROM users WHERE email = \"%s\" ", email))

	var user User
	for res.Next() {
		res.Scan(&user.Password, &user.Sault)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+user.Sault)) != nil {
		return false
	}

	return true
}

func Data_validation(email string, password string, repeat_password string, name string, surname string, sex string, birthday string) string {
	matched, _ := regexp.MatchString(`^[a-z0-9][a-z0-9\._-]*[a-z0-9]*@([a-z0-9]+([a-z0-9-]*[a-z0-9]+)*\.)+[a-z]+`, email)

	if len(email) < 5 || len(email) > 40 || !matched {
		return "Емаил введен некорректно"
	}

	if email_exist(email) {
		return "Аккаунт с введенным email уже существует"
	}

	if password != repeat_password {
		return "Пароли не совпадают"
	}

	if len(password) < 6 || len(password) > 40 {
		return "Пароль введен неккоректно"
	}

	return "ok"
}

func email_exist(email string) bool {
	db := connect_db()
	defer db.Close()

	var key int

	db.QueryRow(fmt.Sprintf("SELECT EXISTS(SELECT id FROM users WHERE email = \"%s\" )", email)).Scan(&key)

	if key == 1 {
		return true
	}
	return false
}
