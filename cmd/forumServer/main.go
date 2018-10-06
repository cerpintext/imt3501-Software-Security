package main

import (
	"fmt"
	"net/http"

	"github.com/krisshol/imt3501-Software-Security/SQLDatabase"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/app"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
)

func main() {

	config.Init()
	database.OpenDB()
	fmt.Printf("Starting server listening on %s with port %d\n", config.Config.Address, config.Config.Port)

	// Handlers:
	http.HandleFunc("/", app.DefaultHandler)
	http.HandleFunc("/signup/", app.SignUpHandler)
	http.HandleFunc("/login/", app.LoginHandler)
	http.HandleFunc("/logout/", app.LogoutHandler)
	http.HandleFunc("/message/", app.MessageHandler)
	http.HandleFunc("/delete/", app.DeleteMessageHandler)
	http.HandleFunc("/thread/", app.ThreadHandler)
	http.HandleFunc("/category/", app.CategoryHandler)

	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Config.Address, config.Config.Port), nil) // Start serving incomming requests. Will continue to serve forever.
}
