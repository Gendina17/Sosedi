package main

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
  "github.com/dchest/uniuri"
  "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
  Id uint16
  Name, Surname string
  Age uint16
  PriceMin, PriceMax uint32
  Password, Sault, Email string
}

func get_users() []User {
  users := []User{}

  db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/sosedi")
  defer db.Close()

  res, _ := db.Query("SELECT id, name, price_min, price_max FROM users ")

   for res.Next() {
     var user User
     res.Scan(&user.Id, &user.Name, &user.PriceMin, &user.PriceMax)
     users = append(users, user)
   }

   return users
}

func get_user(id string) User {
  db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/sosedi")
  defer db.Close()

  res, _ := db.Query(fmt.Sprintf(
    "SELECT id, name, price_min, price_max FROM users WHERE id = %s ", id))

    var user User
    for res.Next() {
      res.Scan(&user.Id, &user.Name, &user.PriceMin, &user.PriceMax)
    }

  return user
}

func user_verification(email string, password string) bool  {
  db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/sosedi")
  defer db.Close()

  res, _ := db.Query(fmt.Sprintf(
    "SELECT id, email, password, sault FROM users WHERE email = %s ", email))

    var user User
    for res.Next() {
      res.Scan(&user.Id, &user.Name, &user.Password, &user.Sault)
    }

    if user != nil && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password + user.Sault)) == 0 {
        return true
    }
    return false
}
//подумать как передавать кучу параметров мб хэш
func create_user(name string, urname string, birthday string, price_min string, price_max string, password string, sault string, email string)  {
  sault = uniuri.NewLen(10)
  encrypted_password, _ := bcrypt.GenerateFromPassword([]byte(password + sault), 5)

  db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/sosedi")
  defer db.Close()

  res, _ := db.Query(fmt.Sprintf(
    `INSERT INTO users (name, surname, birthday, price_min, price_max, password, sault, email)
     VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s')`,
     name, urname, birthday, price_min, price_max, encrypted_password, sault, email))
}
//мб вынесни в отдельный метод обращение к базе ток вот какого типа возвращаемое значение узнать
