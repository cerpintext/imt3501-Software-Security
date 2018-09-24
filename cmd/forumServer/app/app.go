package app

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
// }

// DefaultHandler returns index.html.
func DefaultHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to DefualtHandler\n")

	// Default handler is only GET.
	fmt.Fprint(w, FetchHTML("index.html"))

}

//SignInHandler returns html page if GET, registers new user if POST.
func SignUpHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to SignUpHandler\n")

	switch r.Method {
	case "GET":
		fmt.Fprint(w, FetchHTML("signup.html"))
		break

	case "POST":

		break
	}

}

// LoginHandler returns html page if GET, logs in user if POST.
func LoginHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to SignUpHandler\n")

	switch r.Method {
	case "GET":
		fmt.Fprint(w, FetchHTML("login.html"))
		break

	case "POST":

		break
	}

}
