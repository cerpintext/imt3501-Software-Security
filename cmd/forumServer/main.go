package main

import (
	"fmt"
	"net/http"

	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/app"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
	"github.com/krisshol/imt3501-Software-Security/SQLDatabase"
)

func main() {

	config.Init()
	database.OpenDB()
	fmt.Printf("Starting server listening on %s with port %d\n", config.Config.Address, config.Config.Port)

	// Handlers:
	http.HandleFunc("/", app.DefaultHandler)
	http.HandleFunc("/signup/", app.SignUpHandler)
	http.HandleFunc("/login/", app.LoginHandler)
	http.HandleFunc("/postmessage/", app.PostMessageHandler)
	http.HandleFunc("/postthread/", app.PostThreadHandler)

	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Config.Address, config.Config.Port), nil) // Start serving incomming requests. Will continue to serve forever.
}
