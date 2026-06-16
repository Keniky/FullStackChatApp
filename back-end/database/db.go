package database

import (
	"chatApp/config"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB = nil
var mutex sync.Mutex

func createDBConnection() {
	var err error
	db, err = sql.Open("postgres", config.Dsn)

	if err != nil {

		fmt.Println("failed to open connection with the databse ")
		return
	}

	if err = db.Ping(); err != nil {

		fmt.Println("failed to connect to the database")
		return
	}

}

func GetDBConnection() *sql.DB {
	mutex.Lock()
	if db == nil {
		createDBConnection()
	}
	mutex.Unlock()

	return db
}
