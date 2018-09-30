package config

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"

	"github.com/tkanos/gonfig"
)

const MIN_FIELD_LENGTH = 1      // Minimum length of generic fields.
const MAX_FIELD_LENGTH = 40     // Maximum length of generic fields.
const MIN_PASSWORD_LENGTH = 8   // Minimum length of passwords.
const MAX_EMAIL_LENGTH = 80     // Maximum length of emails.
const MAX_MESSAGE_LENGTH = 2000 // Maximum length of messages.

type Configuration struct {
	Port             int    // The port the server will listen on.
	Address          string // The address the server will serve requests from.
	DatabasePort     int    // The port the remote database uses.
	DatabaseAddress  string // The address to the remote database.
	DatabaseDatabase string // The database to use within the running remote sql server.
	DatabaseUser     string // The username for access to the remote database.
	DatabasePassword string // The password for that user.
	HtmlPath         string // The absolute or relative path to the directory where the html docs are stored.
}

var Config Configuration

// Init loads parameters json file.
func Init() {

	ReadConfigFromFile("config.json")
}

// ReadConfigFromFile reads a json file with configuration values using external liberary: github.com/tkanos/gonfig.
func ReadConfigFromFile(fileName string) {

	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), fileName)

	err := gonfig.GetConf(filePath, &Config)
	if err != nil {

		fmt.Print("Could not read " + filePath + ", or its malformed.\n\n")
		panic(err)
	}

}
