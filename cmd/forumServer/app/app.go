package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/krisshol/imt3501-Software-Security/SQLDatabase"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/util"
)

// DefaultHandler returns index.html.
func DefaultHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to DefualtHandler\n")
	util.PrintURLAsSlice(r.URL.Path)

	if r.Method != "GET" { //Default handler is only GET. No Method switch.
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) >= 3 && parts[1] == "page" { // If there is 2 komponents in URL and the first one is "page". >= 3 Because there is a / at the end of the path as well.

		fmt.Fprint(w, util.FetchHTML(parts[2]+".html"))

	} else {

		fmt.Fprint(w, util.FetchHTML("index.html"))
	}

}

//SignInHandler returns html page if GET, registers new user if POST.
func SignUpHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to SignUpHandler\n")
	util.PrintURLAsSlice(r.URL.Path)

	switch r.Method {
	case "GET":
		fmt.Fprint(w, util.FetchHTML("signup.html"))
		break

	case "POST":
		r.ParseForm()
		// Get fields from form:
		userName := r.FormValue("username")
		userEmail := r.FormValue("email")
		password := r.FormValue("password")
		passwordConfirm := r.FormValue("passwordconfirm")

		// Validate input:
		if !util.BasicValidate(userName) ||
			!util.BasicValidate(userEmail, -1, config.MAX_EMAIL_LENGTH) ||
			!util.BasicValidate(password, config.MIN_PASSWORD_LENGTH) ||
			!util.BasicValidate(passwordConfirm, config.MIN_PASSWORD_LENGTH) {

			w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
			fmt.Fprint(w, "Sorry, some of that was wrong.")
			return
		}
		if password != passwordConfirm {

			w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
			fmt.Fprint(w, "Passowords did not match.")
			return
		}

		// TODO: Hash password.

		var user database.User // Create struct to store user input to be sent to db:
		user.Username = userName
		user.Email = userEmail
		user.PasswordHash = password
		fmt.Printf("User input accepted. Inserting user into db: \nusername: %s\n\n", userName) // TODO: Remove test outprint.
		database.OpenDB()
		database.AddUser(user) // Send struct to db.

		fmt.Fprint(w, "All good, welcome to the team "+userName+"! :D")
		break
	default:
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}
}

// LoginHandler returns html page if GET, logs in user if POST.
func LoginHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to LoginHandler\n")
	util.PrintURLAsSlice(r.URL.Path)

	switch r.Method {
	case "GET":
		fmt.Fprint(w, util.FetchHTML("login.html"))
		break

	case "POST":

		break
	default:
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}

}

// PostMessageHandler returns html page if GET, logs in user if POST.
func PostMessageHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to PostMessageHandler\n")
	util.PrintURLAsSlice(r.URL.Path)

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}

	r.ParseForm()

	message, err := util.ValidateMessage(r.FormValue("message"), r.FormValue("username"), r.FormValue("parentmessage"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		fmt.Fprint(w, err)
		return
	}

	fmt.Printf("User input accepted. Inserting message into db: \nmessage(first 20 chars): \"%s\"\nusername: %s\nparent: %d\n\n",
		message.Message[0:20], message.Username, message.ParentMessage) // TODO: Remove test outprint.

	database.OpenDB()
	database.AddMessage(message)
	fmt.Fprint(w, "Message sent.\n")
}
