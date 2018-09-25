package util

import (
	"fmt"
	"io/ioutil"

	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
)

// FetchHTML takes a filename of an html doc in the htmldirectory configured, reads and returns it.
func FetchHTML(fileName string) string {

	fmt.Printf("Http request for html file: %s\n", config.HtmlPath+fileName)
	data, err := ioutil.ReadFile(config.HtmlPath + fileName) // Attempt to read desired file.
	if err != nil {
		fmt.Printf("Something went wrong fetching file: %s:\n %s\n\n", config.HtmlPath+fileName, string(data))
	} else {

		fmt.Printf("Serving\n\n")
		return string(data)
	}
	return ""
}

// func parseURL(URL string) string {

// 	fileNames := strings.Split(URL, "/")
// 	fileName := strings.ToLower(fileNames[len(fileNames)-2]) + ".html" // Create desired filename path.

// 	if fileName == ".html" {
// 		fileName = "index.html"
// 	}
// 	return fileName
//

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

		fmt.Print("BaiscValidate. Input not valid: " + field)
		return false
	}

	// TODO: More validation.

	return true
}