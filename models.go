package main

import (
  "fmt"
  "database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
  Id uint16
  Name, Surname string
  Age uint16
  PriceMin, PriceMax uint32
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
