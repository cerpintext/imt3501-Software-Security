package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // global value for keeping database open

type Category struct {
}

type Thread struct {
}

type User struct {
}

type Message struct {
}

func openDB() {
	var err error
	db, err = sql.Open("mysql", "USERNAME:PASSWORD@tcp(IP-ADDRESS:PORT)/DATABASE")
	if err != nil {
		panic(err.Error()) // Implement proper handlig
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // Implement proper handlig
	}
}

func AddThread() {
	openDB()
}

func AddUser(username string, passwordHash string) {
	openDB()
}

func AddMessage(message string, username string) {
	openDB()
}
