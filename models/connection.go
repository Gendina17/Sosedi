package models

import "database/sql"

type User struct {
	Id                     uint16
	Name, Surname          string
	Age                    uint16
	PriceMin, PriceMax     uint32
	Password, Sault, Email string
	Sex                    string
	Photo                  string
}

func connect_db() *sql.DB {
	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/sosedi")
	return db
}
