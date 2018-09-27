package database

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
	//	id       int
	Name     string
	Username string
}

type User struct {
	Username     string
	Email        string
	PasswordHash string
	//	reputation   int
	//	role         int
}

type Message struct {
	//	id            int
	Message string
	//	timestamp     string
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
		log.Fatalln("Could not connect to the database")
	}
}

//	How to use
//	AddThread(Thread{intId, "name", "existingUsername"})
func AddThread(class interface{}) {
	if c, ok := class.(Thread); ok { // type assert on it
		stmtIns, err := db.Prepare("INSERT INTO Thread (`name`, `username`) VALUES( ?, ? )") // ? = placeholder
		if err != nil {
			panic(err.Error()) // TODO: Implement proper handlig
		}
		defer stmtIns.Close() // Close the statement when we leave function() / the program terminates

		_, err = stmtIns.Exec(c.Name, c.Username) // Insert tuples (id, name, userName)
		if err != nil {
			errorHandling(err, "thread")
		}
	}
}

//	How to use
//	AddUser(User{"userName", "email", "passwordHash" reputation, role})
func AddUser(class interface{}) {
	if c, ok := class.(User); ok { // type assert on it
		//	if checkUserExists(c.username) == true {
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
		//	}
	}
}

//	How to use
//	AddMessage(Message{intId, "the message to be posted", \
//		"timestamp on mysql accepted format e.g 1971-01-01 00:00:00 as a string", "username", parentMessage})
func AddMessage(class interface{}) {
	if c, ok := class.(Message); ok { // type assert on it
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
	} else if mysqlErr.Number == 1452 && function == "message" && string(runes[135:143]) == "username" { // Non existent user
		log.Println("Hmm, that's not supposed to happen. User not found when posting message")
	} else if mysqlErr.Number == 1452 && function == "message" && string(runes[135:148]) == "parentmessage" { // Non existent parent message
		log.Println("Hmm, that's not supposed to happen. Parent message message not found") // output might need to be changed
	} else { // Unkown error
		//log.Println(string(runes[135:148]))
		panic(err.Error()) // TODO: Implement proper handling.
	}
}
