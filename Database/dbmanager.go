package dbmanagaer

import (
	"database/sql"
	"log"

	"github.com/gomodule/redigo/redis"
)

var Db *sql.DB
var Cache redis.Conn

func InitDB() {
	var err error
	log.Println("Initializing Database .....")
	//sql.open() function only gives the handler to talk to db, it does not create any connection with Db.
	Db, err = sql.Open("mysql", "root:Rajendra95@@tcp(127.0.0.1:3306)/userlogindata")
	if err != nil {
		panic(err)
	}
	// ping function will actually verify the connection with database.
	err = Db.Ping()
	if err != nil {
		panic(err.Error())
	}
	log.Println("Database Connection Establised")
}

func InitCache() {
	// Initialize the redis connection to a redis instance running on your local machine
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	// Assign the connection to the package level `cache` variable
	Cache = conn
	log.Println("Cache connction", Cache)
}
