package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const address = "127.0.0.1" // Localhost.
const port = "5000"         // Must be an open port. On linux open with $source PORT=5000
const htmlPath = "../html/"

func defaultHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.
	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to /.\n")
	switch r.Method { // Do different things depending on type of request.

	case "GET":

		fileNames := strings.Split(r.URL.Path, "/")
		fileName := fileNames[len(fileNames)-2] + ".html" // Create desired filename path.
		if fileName == ".html" {
			fileName = "index.html"
		}
		fmt.Printf("Http request for html file: %s\n", fileName)
		data, err := ioutil.ReadFile(htmlPath + fileName) // Attempt to read desired file.
		if err != nil {
			fmt.Printf("Something went wrong fetching file: %s:\n %s\n\n", fileName, string(data))
		} else {

			fmt.Printf("Serving\n\n")
			fmt.Fprint(w, string(data)) // If read went ok, wend file.
		}
		break

	case "POST":

		fmt.Fprint(w, "Hello world! You did a post request.\n") // Write hello world to response writer.
		break
	}
}

func main() {

	fmt.Print("Starting server listening on " + address + " with port " + port + "\n")
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(address+":"+port, nil) // Start serving incomming requests. Will continue to serve forever.
}
