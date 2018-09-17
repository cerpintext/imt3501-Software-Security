package main

import (
	"fmt"
	"net/http"
)

const address = "127.0.0.1" // Localhost.
const port = "5000"         // Must be an open port. On linux open with $source PORT=5000

func defaultHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.
	fmt.Fprintf(w, "Hello world!") // Write hello world to respone writer.
}

func main() {
	fmt.Print("Starting server listening on " + address + " with port " + port)
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(address+":"+port, nil) // Start serving incomming requests. Will continue to serve forever.
}
