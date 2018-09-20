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
	w.Header().Set("Content-Type", "text/html")
	fmt.Print("Received a request to /.")
	switch r.Method {

	case "GET":

		fileNames := strings.Split(r.URL.Path, "/")
		fileName := fileNames[len(fileNames)-2] + ".html"
		fmt.Printf("Http request for html file: %s", fileName)
		data, err := ioutil.ReadFile(htmlPath + fileName)
		if err != nil {
			fmt.Print("Something went wrong: " + string(data))
			panic(err)
		}
		fmt.Fprint(w, string(data))
		break

	case "POST":

		fmt.Fprint(w, "Hello world! You did a post request.") // Write hello world to respone writer.
		break
	}
}

func main() {

	fmt.Print("Starting server listening on " + address + " with port " + port)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(address+":"+port, nil) // Start serving incomming requests. Will continue to serve forever.
}
