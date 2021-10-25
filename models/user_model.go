package models

import (
	"fmt"

	"github.com/dchest/uniuri"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

func Get_users() []User {
	users := []User{}

	db := connect_db()
	defer db.Close()

	res, _ := db.Query("SELECT id, name, price_min, price_max, photo_key FROM users ")

	for res.Next() {
		var user User
		res.Scan(&user.Id, &user.Name, &user.PriceMin, &user.PriceMax, &user.Photo)
		if user.Photo == "" {
			user.Photo = "614fa07d2e9e9f844c05567e.jpeg"
		}
		users = append(users, user)
	}

	return users
}

func Get_user(id string) User {
	db := connect_db()
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf(
		"SELECT id, name, price_min, price_max FROM users WHERE id = %s ", id))

	var user User
	for res.Next() {
		res.Scan(&user.Id, &user.Name, &user.PriceMin, &user.PriceMax)
	}

	return user
}

func Create_user(email string, password string, name string, surname string, sex string, birthday string, photo_key string) {
	sault := uniuri.NewLen(10)
	encrypted_password, _ := bcrypt.GenerateFromPassword([]byte(password+sault), 5)

	db := connect_db()
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf("INSERT INTO users (name, surname, birthday, price_min, price_max, password, sault, email, sex, photo_key) VALUES ('%s','%s','%s','34000','23000','%s','%s','%s','%s','%s')", name, surname, birthday, encrypted_password, sault, email, sex, photo_key))
	defer res.Close()
}

func Get_user_by_email(email string) User {
	db := connect_db()
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf(
		"SELECT id, name, price_min, price_max FROM users WHERE email = \"%s\" ", email))

	var user User
	for res.Next() {
		res.Scan(&user.Id, &user.Name, &user.PriceMin, &user.PriceMax)
	}

	return user
}

// TODO: передавать невидимым инпутом токен посмотреть как его можн шифровать
// TODO: подумать мб как т лучше передавать эту кучу параметров
// TODO: мб др писать
