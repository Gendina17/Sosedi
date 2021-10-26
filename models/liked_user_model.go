package models

import "fmt"

//	добавить в таблицу лайкнутых одну строку
func Add_to_liked(user_id uint16, liked_user_id uint16) {
	db := connect_db()
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf("INSERT INTO liked_users (user_id, liked_user_id) VALUES ('%d','%d')", user_id, liked_user_id))
	defer res.Close()
}

//	взять из таблицы все айди лайкнутых одним пользователем
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

//	взять из таблицы все данные лайкнутых пользователей
func Get_liked_users(user_id uint16) []UserWithComment {
	liked_users := []UserWithComment{}

	db := connect_db()
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf("SELECT users.id, users.name, users.price_min, users.price_max, users.photo_key, liked_users.comment FROM users, liked_users WHERE liked_users.user_id = %d AND users.id = liked_users.liked_user_id", user_id))

	for res.Next() {
		var liked_user UserWithComment
		res.Scan(&liked_user.Id, &liked_user.Name, &liked_user.PriceMin, &liked_user.PriceMax, &liked_user.Photo, &liked_user.Comment)
		if liked_user.Photo == "" {
			liked_user.Photo = "614fa07d2e9e9f844c05567e.jpeg"
		}
		liked_users = append(liked_users, liked_user)
	}

	return liked_users
}

//	добавить (или обновить) значение в поле коммент
func Add_comment(comment string, liked_user_id uint16, user_id uint16) {
	db := connect_db()
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf("UPDATE liked_users SET comment = '%s' WHERE liked_user_id = %d AND user_id = %d", comment, liked_user_id, user_id))
	defer res.Close()
}

//	удалить запись из таблицы лайкнутых по данным айди
func Remove_from_liked(liked_user_id uint16, user_id uint16) {
	db := connect_db()
	defer db.Close()

	res, _ := db.Query(fmt.Sprintf("DELETE FROM liked_users WHERE liked_user_id = %d AND user_id = %d", liked_user_id, user_id))
	defer res.Close()
}
