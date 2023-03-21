package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() *sql.DB{
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/movie?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	return DB
}

