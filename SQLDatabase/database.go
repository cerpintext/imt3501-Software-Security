package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB // global value for keeping database open

type Category struct {
	Name string
}

type Thread struct {
	Id       int
	Name     string
	Username string
}

type User struct {
	Username     string
	Email        string
	PasswordHash string
	Reputation   int
	Role         int
}

type Message struct {
	Id            int
	Message       string
	Timestamp     string
	Username      string
	ParentMessage int
}

func OpenDB() {
	var err error
	db, err = sql.Open("mysql", "USERNAME:PASSWORD@tcp(IP-ADDRESS:PORT)/DATABASE")
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		//log.Fatalln("Could not connect to the database")
		errorHandling(err, "ping")
	}
}

//	Only uses fields Name and Username
//	How to use
//	AddThread(Thread{intId, "name", "existingUsername"})
func AddThread(c Thread) {
	stmtIns, err := db.Prepare("INSERT INTO Thread (`name`, `username`) VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		log.Fatalln("wut?")
		panic(err.Error()) // TODO: Implement proper handlig
	}
	defer stmtIns.Close()                     // Close the statement when we leave function() / the program terminates
	_, err = stmtIns.Exec(c.Name, c.Username) // Insert tuples (name, userName)
	if err != nil {
		errorHandling(err, "thread")
	}
}

//	Only uses fields Username, Email, and passwordHash
//	How to use
//	AddUser(User{"userName", "email", "passwordHash" reputation, role})
func AddUser(c User) {
	stmtIns, err := db.Prepare("INSERT INTO User (`username`, `email`, `passwordHash`) VALUES( ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}
	defer stmtIns.Close() // Close the statement when we leave function() / the program terminates
	// Insert tuples (username, email, passwordHash, reputation, role)
	_, err = stmtIns.Exec(c.Username, c.Email, c.PasswordHash)
	if err != nil {
		errorHandling(err, "user")
	}
}

//	Only uses fields Message, Username and ParentMessage
//	How to use
//	AddMessage(Message{intId, "the message to be posted", \
//		"timestamp on mysql accepted format e.g 1971-01-01 00:00:00 as a string", "username", parentMessage})
func AddMessage(c Message) {
	stmtIns, err := db.Prepare("INSERT INTO Message (`message`, `username`, `parentmessage`) VALUES( ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}
	defer stmtIns.Close() // Close the statement when we leave function() / the program terminates
	// Insert tuples (message, username, parentMessage)
	_, err = stmtIns.Exec(c.Message, c.Username, c.ParentMessage)
	if err != nil {
		errorHandling(err, "message")
	}
}

/*************** Help functions ***************/

func errorHandling(err error, function string) {
	mysqlErr := err.(*mysql.MySQLError) // Asserting mysql error struct
	runes := []rune(mysqlErr.Message)
	if mysqlErr.Number == 1062 && function == "user" { // Duplicate username
		log.Println("Username already exists")
	} else if mysqlErr.Number == 1062 && function == "thread" { // Duplicate thread
		log.Println("Something strange happend a thread with this ID already exists")
	} else if mysqlErr.Number == 1062 && function == "message" { // Duplicate message
		log.Println("Something strange happend a message with this ID already exists")
	} else if mysqlErr.Number == 1452 && function == "message" && string(runes[134:143]) == "username" { // Non existent user
		log.Println("Hmm, that's not supposed to happen. User not found when posting message")
	} else if mysqlErr.Number == 1452 && function == "message" && string(runes[134:148]) == "parentmessage" { // Non existent parent message
		log.Println("Hmm, that's not supposed to happen. Parent message message not found") // output might need to be changed
	} else if mysqlErr.Number == 1452 && function == "thread" && string(runes[133:141]) == "username" { // Non existent parent message
		log.Println("Hmm, that's not supposed to happen. Username does not exist when posting a new thread") // output might need to be changed
	} else if function == "ping" { // Non existent parent message
		log.Println("Could not connect to the database") // output might need to be changed
	} else { // Unkown error
		//log.Println(string(runes[135:148]))
		panic(err.Error()) // TODO: Implement proper handling.
	}
}
