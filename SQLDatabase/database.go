package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB // global value for keeping database open

type Category struct {
	name string
}

type Thread struct {
	id       int
	name     string
	username string
}

type User struct {
	username     string
	email        string
	passwordHash string
	reputation   int
	role         int
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

//	How to use
//	AddMessage(Message{intId, "the message to be posten", \
//		"timestamp on mysql accepted format e.g 1971-01-01 00:00:00 as a string", "parentMessage"})
func AddThread() {
	openDB()                          // should find better way to handle db connection globally
	if c, ok := class.(Message); ok { // type assert on it
		stmtIns, err := db.Prepare("INSERT INTO Thread VALUES( ?, ?, ? )") // ? = placeholder
		if err != nil {
			panic(err.Error()) // Implement proper handlig
		}
		defer stmtIns.Close() // Close the statement when we leave function() / the program terminates

		_, err = stmtIns.Exec(c.id, c.name, c.username) // Insert tuples (id, name, userName)
		if err != nil {
			panic(err.Error()) // Implement proper handlig
		}
	}
	db.close() // should find better way to handle db connection globally
}

//	How to use
//	AddUser(Message{"userName", "email", "passwordHash" "reputation", "role"})
func AddUser(username string, passwordHash string) {
	openDB()                          // should find better way to handle db connection globally
	if c, ok := class.(Message); ok { // type assert on it
		stmtIns, err := db.Prepare("INSERT INTO User VALUES( ?, ?, ?, ?, ? )") // ? = placeholder
		if err != nil {
			panic(err.Error()) // Implement proper handlig
		}
		defer stmtIns.Close() // Close the statement when we leave function() / the program terminates
		// Insert tuples (username, email, passwordHash, reputation, role)
		_, err = stmtIns.Exec(c.username, c.email, c.passwordHash, c.reputation, c.role)
		if err != nil {
			panic(err.Error()) // Implement proper handlig
		}
	}
	db.close() // should find better way to handle db connection globally
}

func AddMessage(message string, username string) {
	openDB()
}
