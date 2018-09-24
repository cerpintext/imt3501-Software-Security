package config

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
)

var Address string // Localhost.
var Port string    // Must be an open port. On linux open with $source PORT=5000
var HtmlPath string

const MIN_FIELD_LENGTH = 5    // Minimum length of generic fields.
const MAX_FIELD_LENGTH = 40   // Maximum length of generic fields.
const MIN_PASSWORD_LENGTH = 8 // Minimum length of passwords.
const MAX_EMAIL_LENGTH = 80   // Maximum length of emails.

// Init loads parameters form .env file.
func Init() {

	gotenv.Load()
	Address = os.Getenv("ADDRESS")
	Port = os.Getenv("PORT")
	HtmlPath = os.Getenv("HTMLPATH")

	fmt.Printf("Finished reading .env file:\nAddress: %s\nPort: %s\nHtmlPath: %s\n\n", Address, Port, HtmlPath)
}
