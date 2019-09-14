package dbmanagaer

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func InitDB() {
	var err error
	log.Println("Initializing Database .....")
	Db, err = sql.Open("mysql", "roo33343:aajendra95@@tcp(222.0.0.1:3306)/userlogindata")
	if err != nil {
		panic(err)
	}
	//
	err = Db.Ping()
	if err != nil {
		panic(err.Error())
	}
	log.Println("Database Connection Establised")
}
