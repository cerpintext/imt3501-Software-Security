package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const address = "127.0.0.1" // Localhost.
const port = "5000"         // Must be an open port. On linux open with $source PORT=5000
const htmlPath = "../html/"

func defaultHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	fmt.Print("Received a request to /.")
	switch r.Method {
	case "GET":
		data, err := ioutil.ReadFile(htmlPath + "index.html")
		if err != nil {
			fmt.Print("Something went wrong: " + string(data))
			panic(err)
		}
		fmt.Fprint(w, string(data))
		break
	case "POST":
		fmt.Fprintf(w, "Hello world! You did a post request.") // Write hello world to respone writer.
		break
	}
}

func main() {

	fmt.Print("Starting server listening on " + address + " with port " + port)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(address+":"+port, nil) // Start serving incomming requests. Will continue to serve forever.
}
