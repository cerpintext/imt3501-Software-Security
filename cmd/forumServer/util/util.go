package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	database "github.com/krisshol/imt3501-Software-Security/SQLDatabase"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
)

// FetchHTML takes a filename of an html doc in the htmldirectory configured, reads and returns it.
func FetchHTML(fileName string) string {

	fmt.Printf("Http request for html file: %s\n", config.Config.HtmlPath+fileName)
	data, err := ioutil.ReadFile(config.Config.HtmlPath + fileName) // Attempt to read desired file.
	if err != nil {
		fmt.Printf("Something went wrong fetching file: %s:\n %s\n\n", config.Config.HtmlPath+fileName, string(data))
	} else {

		fmt.Printf("Serving\n\n")
		return string(data)
	}
	return ""
}

// PrintURLAsSlice prints the individual URL indecies and values string split in "/".
func PrintURLAsSlice(URL string) {

	fmt.Printf("Printing URL as slice: %s\n", URL)
	parts := strings.Split(URL, "/")
	for i, part := range parts {
		fmt.Printf("%d: %s\n", i, part)
	}
}

// BasicValidate returns false if any anomalies are detected, like empty string. Optional parameters are an int for custom min length, an int for custom max length.
func BasicValidate(field string, param ...int) bool {

	minLength := config.MIN_FIELD_LENGTH
	maxLength := config.MAX_FIELD_LENGTH

	if len(param) >= 1 && param[0] >= 0 {

		minLength = param[0]
	}

	if len(param) >= 2 {

		maxLength = param[1]
	}

	if len(field) < minLength || len(field) >= maxLength {

		fmt.Printf("BaiscValidate() Input not valid: %s\n\n", field)
		return false
	}

	// TODO: More validation.

	fmt.Printf("\nBasicValidate() Input valid: (len: %d string: %s)\n\tString is less than max %d. String is more than min %d\n", len(field), field, maxLength, minLength)
	return true
}

func ValidateMessage(messageText string, username string, parentMessage string) (database.Message, error) {

	var message database.Message

	message.Message = messageText
	message.Username = username

	parent, err := strconv.Atoi(parentMessage)
	if err != nil {
		fmt.Printf("Validate Message: Failed to parse message.parentmessage, got: %s\n\n\n", parentMessage)
		return message, errors.New("Message was invalid")
	}
	message.ParentMessage = parent

	fmt.Printf("Validate Message(): parentMessage: %d", parent)

	// TODO: check if user exist.

	if !BasicValidate(message.Message, -1, config.MAX_MESSAGE_LENGTH) ||
		!BasicValidate(message.Username) ||
		message.ParentMessage < -1 {

		fmt.Print("Validate Message(): Message rejected.\n\n")
		return message, errors.New("Message was invalid")
	}

	return message, nil

}

func IsLoggedIn(r *http.Request, username string) bool {

	fmt.Printf("IsLoggedIn(): Checking if user %s is logged in.\n", username)
	storedSession, exist := config.SessionMap[username]

	if exist {

		cookieSession, err := r.Cookie("session")
		if err != nil { //The user has registered session but their cookie is expired.

			fmt.Printf("IsLoggedIn(): User %s session was stored but session was expired. Deleting old session. User is NOT logged in.\n", username)
			delete(config.SessionMap, username) // Delete stored session id.
			return false
		}

		if cookieSession.Value == storedSession { // The user has registered session and  still has their cookie(not expired).

			fmt.Printf("IsLoggedIn(): User %s is logged in.\n", username)
			return true

		}
	}
	fmt.Printf("IsLoggedIn(): User %s is NOT logged in.\n", username)
	return false
}
