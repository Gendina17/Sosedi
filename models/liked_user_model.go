package models

import "fmt"

//	добавить в таблицу лайкнутых одну строку
func Add_to_liked(user_id uint16, liked_user_id uint16) {
	db := connect_db()
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf("INSERT INTO liked_users (user_id, liked_user_id) VALUES ('%d','%d')", user_id, liked_user_id))
	defer res.Close()
}

//	взять из таблицы всех лайкнутых одним пользователем
func Get_liked_users_id(user_id uint16) []uint16 {
	db := connect_db()
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf("SELECT liked_user_id FROM liked_users WHERE user_id = %d", user_id))

	liked_users := []uint16{}

	for res.Next() {
		var liked_user uint16
		res.Scan(&liked_user)
		liked_users = append(liked_users, liked_user)
	}

	return liked_users
}

func Get_liked_users(user_id uint16) []User {
	liked_users := []User{}

	db := connect_db()
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf("SELECT users.id, users.name, users.price_min, users.price_max, users.photo_key FROM users, liked_users WHERE liked_users.user_id = %d AND users.id = liked_users.liked_user_id", user_id))

	for res.Next() {
		var liked_user User
		res.Scan(&liked_user.Id, &liked_user.Name, &liked_user.PriceMin, &liked_user.PriceMax, &liked_user.Photo)
		if liked_user.Photo == "" {
			liked_user.Photo = "614fa07d2e9e9f844c05567e.jpeg"
		}
		liked_users = append(liked_users, liked_user)
	}

	return liked_users
}
