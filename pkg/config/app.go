package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Connect() {

	// cfg := mysql.Config{
	// 	User:   os.Getenv("root@localhost"),
	// 	Passwd: os.Getenv(""),
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "movie",
	// }

	// db, err = sql.Open("mysql", cfg.FormatDSN())
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/movie?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
}
func GetDB() *sql.DB{
	return db
}
