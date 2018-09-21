package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/subosito/gotenv"
)

var address string // Localhost.
var port string    // Must be an open port. On linux open with $source PORT=5000
var htmlPath string

// Init loads parameters form .env file.
func Init() {

	gotenv.Load()
	address = os.Getenv("ADDRESS")
	port = os.Getenv("PORT")
	htmlPath = os.Getenv("HTMLPATH")
}

func defaultHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.
	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to /.\n")
	switch r.Method { // Do different things depending on type of request.

	case "GET":

		fileNames := strings.Split(r.URL.Path, "/")
		fileName := strings.ToLower(fileNames[len(fileNames)-2]) + ".html" // Create desired filename path.
		if fileName == ".html" {
			fileName = "index.html"
		}
		fmt.Printf("Http request for html file: %s\n", fileName)
		data, err := ioutil.ReadFile(htmlPath + fileName) // Attempt to read desired file.
		if err != nil {
			fmt.Printf("Something went wrong fetching file: %s:\n %s\n\n", fileName, string(data))
		} else {

			fmt.Printf("Serving\n\n")
			fmt.Fprint(w, string(data)) // If read went ok, send file.
		}
		break

	case "POST":

		fmt.Fprint(w, "Hello world! You did a post request.\n") // Write hello world to response writer.
		break
	}
}

func main() {

	Init()
	fmt.Print("Starting server listening on " + address + " with port " + os.Getenv("PORT") + "\n")
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(address+":"+os.Getenv("PORT"), nil) // Start serving incomming requests. Will continue to serve forever.
}
