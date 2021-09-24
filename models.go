package main

import (
  "fmt"
  "regexp"
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
  Sex string
}

func get_users() []User {
  users := []User{}

  db := connect_db()
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

func user_verification(email string, password string) bool {
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

  if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password + user.Sault)) != nil {
    return false
  }

  return true
}

func create_user(email string, password string, name string, surname string, sex string, birthday string)  {
  sault := uniuri.NewLen(10)
  encrypted_password, _ := bcrypt.GenerateFromPassword([]byte(password + sault), 5)

  db := connect_db()
  defer db.Close()

  res, _ := db.Query(fmt.Sprintf("INSERT INTO users (name, surname, birthday, price_min, price_max, password, sault, email, sex) VALUES ('%s','%s','%s','34000','23000','%s','%s','%s', '%s')", name, surname, birthday, encrypted_password, sault, email, sex))
  defer res.Close()
}

func data_validation(email string, password string, repeat_password string, name string, surname string, sex string, birthday string, key string) string {
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

func connect_db() *sql.DB {
  db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/sosedi")
  return db
}

func get_user_by_email(email string) User {
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
