package database

import (
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
}

type Thread struct {
}

type User struct {
}

type Message struct {
}

func openDB() {

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
