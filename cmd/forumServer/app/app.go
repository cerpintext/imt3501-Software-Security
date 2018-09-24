package app

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
)

func RegisterUser(r *http.Request) {

}

// DefaultHandler returns html page for requested file if GET, and handles input to create new message or user, or upvote if POST.
func DefaultHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to /.\n")

	fileNames := strings.Split(r.URL.Path, "/")
	fileName := strings.ToLower(fileNames[len(fileNames)-2]) + ".html" // Create desired filename path.
	if fileName == ".html" {
		fileName = "index.html"
	}

	switch r.Method { // Do different things depending on type of request.
	case "GET":

		fmt.Printf("Http request for html file: %s\n", config.HtmlPath+fileName)
		data, err := ioutil.ReadFile(config.HtmlPath + fileName) // Attempt to read desired file.
		if err != nil {
			fmt.Printf("Something went wrong fetching file: %s:\n %s\n\n", config.HtmlPath+fileName, string(data))
		} else {

			fmt.Printf("Serving\n\n")
			fmt.Fprint(w, string(data)) // If read went ok, send file.
		}
		break

	case "POST":

		switch fileName {
		case "signup":
			RegisterUser(r)
			break
		}

		break
	}
}
