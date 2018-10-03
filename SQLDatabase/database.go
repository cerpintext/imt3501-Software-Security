package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
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
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		config.Config.DatabaseUser,
		config.Config.DatabasePassword,
		config.Config.DatabaseAddress,
		config.Config.DatabasePort,
		config.Config.DatabaseDatabase))
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		log.Fatalln("Could not connect to the database")
	}
}

/*************** Add functions ***************/

//	Only uses fields Name and Username
//	How to use
//	AddThread(Thread{anInt, "name", "existingUsername"})
func AddThread(c Thread, m Message) {
	stmtIns, err := db.Prepare("INSERT INTO Thread (`name`, `username`) VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}
	defer stmtIns.Close()                        // Close the statement when we leave function() / the program terminates
	res, err := stmtIns.Exec(c.Name, c.Username) // Insert tuples (name, userName)
	if err != nil {
		errorHandling(err, "addThread")
	}

	threadID, err := res.LastInsertId() // Get IDs to create realtion between the Thread and its root Message.
	if err != nil {
		errorHandling(err, "addThread")
	}
	messageID := AddMessage(m)

	stmtIns, err = db.Prepare("INSERT INTO ThreadMessages (`threadId`, `messageId`) VALUES( ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}
	res, err = stmtIns.Exec(threadID, messageID) // Insert tuples (name, userName)
	if err != nil {
		errorHandling(err, "addThread")
	}
}

//	Only uses fields Username, Email, and passwordHash
//	How to use
//	AddUser(User{"userName", "email", "passwordHash" anInt, anInt})
func AddUser(c User) {
	stmtIns, err := db.Prepare("INSERT INTO User (`username`, `email`, `passwordHash`) VALUES( ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}
	defer stmtIns.Close() // Close the statement when we leave function() / the program terminates
	// Insert tuples (username, email, passwordHash, reputation, role)
	_, err = stmtIns.Exec(c.Username, c.Email, c.PasswordHash)
	if err != nil {
		errorHandling(err, "addUser")
	}
}

//	Only uses fields Message, Username and ParentMessage
//	How to use
//	AddMessage(Message{anInt, "the message to be posted", \
//		"", "username", parentMessage})
func AddMessage(c Message) int {
	stmtIns, err := db.Prepare("INSERT INTO Message (`message`, `username`, `parentmessage`) VALUES( ?, ?, ? )") // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}
	defer stmtIns.Close() // Close the statement when we leave function() / the program terminates
	// Insert tuples (message, username, parentMessage)

	var res sql.Result
	if c.ParentMessage == -1 { // No parent. ints can't be nil in golang.

		res, err = stmtIns.Exec(c.Message, c.Username, nil)
	} else {

		res, err = stmtIns.Exec(c.Message, c.Username, c.ParentMessage)
	}

	if err != nil {
		errorHandling(err, "addMessage")
	}
	messageID, err := res.LastInsertId()
	if err != nil {
		errorHandling(err, "addThread")
	}
	return int(messageID)
}

/*************** Delete functions ***************/

//	Only uses field Username
//	How to use
//	DeleteUser(User{"userName", "", "" anInt, anInt})
func DeleteUser(c User) {
	stmtIns, err := db.Prepare("DELETE FROM User WHERE username = ?") // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}
	defer stmtIns.Close() // Close the statement when we leave function() / the program terminates
	// Insert tuples (username, email, passwordHash, reputation, role)
	Result, err := stmtIns.Exec(c.Username)
	if err != nil {
		//panic(err.Error())
		errorHandling(err, "delUser")
	}
	rows, err := Result.RowsAffected()
	if rows == 0 {
		fmt.Println("Hmm, something strange happend; \n\tUser not found -> user not deleted")
	}
}

//	Only uses fields Name and Username
//	How to use
//	DeleteThread(Thread{threadId, "", ""})
func DeleteThread(c Thread) {
	stmtIns, err := db.Prepare("DELETE FROM Thread WHERE id = ?") // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}
	defer stmtIns.Close() // Close the statement when we leave function() / the program terminates
	// Insert tuples (username, email, passwordHash, reputation, role)
	Result, err := stmtIns.Exec(c.Id)
	if err != nil {
		errorHandling(err, "delThread")
	}
	rows, err := Result.RowsAffected()
	if rows == 0 {
		fmt.Println("Hmm, something strange happend; \n\tThread not found -> thread not deleted")
	}
}

//	Only uses field Username
//	How to use
//	DeleteMessage(Message{anInt, "", "", "", anInt})
func DeleteMessage(c Message) {
	stmtIns, err := db.Prepare("DELETE FROM Message WHERE id = ?") // ? = placeholder
	if err != nil {
		panic(err.Error()) // TODO: Implement proper handlig
	}
	defer stmtIns.Close() // Close the statement when we leave function() / the program terminates
	// Insert tuples (username, email, passwordHash, reputation, role)
	Result, err := stmtIns.Exec(c.Id)
	if err != nil {
		errorHandling(err, "delMessage")
	}
	rows, err := Result.RowsAffected()
	if rows == 0 {
		fmt.Println("Hmm, something strange happend; \n\tMessage not found -> message not deleted")
	}
}

/*************** Help functions ***************/
// so far only for mysql statements
func errorHandling(err error, function string) {
	mysqlErr := err.(*mysql.MySQLError) // Asserting mysql error struct
	runes := []rune(mysqlErr.Message)
	if mysqlErr.Number == 1062 && function == "addUser" { // Duplicate username
		log.Println("Username already exists")
	} else if mysqlErr.Number == 1062 && function == "addThread" { // Duplicate thread
		log.Println("Something strange happend a thread with this ID already exists")
	} else if mysqlErr.Number == 1062 && function == "addMessage" { // Duplicate message
		log.Println("Something strange happend a message with this ID already exists")
	} else if mysqlErr.Number == 1452 && function == "addMessage" && string(runes[135:143]) == "username" { // Non existent user
		log.Println("Hmm, that's not supposed to happen. User not found when posting message")
	} else if mysqlErr.Number == 1452 && function == "addMessage" && string(runes[135:148]) == "parentmessage" { // Non existent parent message
		log.Println("Hmm, that's not supposed to happen. Parent message message not found") // output might need to be changed
	} else if mysqlErr.Number == 1452 && function == "addThread" && string(runes[133:141]) == "username" { // Non existent parent message
		log.Println("Hmm, that's not supposed to happen. Username does not exist when posting a new thread") // output might need to be changed
	} else if mysqlErr.Number == 1054 { // Non existent parent message
		log.Println(mysqlErr.Message) // output might need to be changed
	} else { // Unkown error
		//log.Println(string(runes[135:143])) // Causes out of bounds exceptions if error not long enough. Careful.
		panic(err.Error())
	}
}
