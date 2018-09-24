package config

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
)

var Address string // Localhost.
var Port string    // Must be an open port. On linux open with $source PORT=5000
var HtmlPath string

// Init loads parameters form .env file.
func Init() {

	gotenv.Load()
	Address = os.Getenv("ADDRESS")
	Port = os.Getenv("PORT")
	HtmlPath = os.Getenv("HTMLPATH")

	fmt.Printf("Finished reading .env file:\nAddress: %s\nPort: %s\nHtmlPath: %s\n\n", Address, Port, HtmlPath)
}
