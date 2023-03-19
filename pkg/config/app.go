package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() *sql.DB{

	// cfg := mysql.Config{
	// 	User:   os.Getenv("root@localhost"),
	// 	Passwd: os.Getenv(""),
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "movie",
	// }

	// db, err = sql.Open("mysql", cfg.FormatDSN())
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/movie?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }
	//defer DB.Close()
	return DB
}

// func GetDB() *sql.DB{
// 	//fmt.Println(DB, "db <-------->")
// 	return DB
// }
