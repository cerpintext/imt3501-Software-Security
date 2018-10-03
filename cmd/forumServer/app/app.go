package app

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"strconv"
	"github.com/nu7hatch/gouuid"



	"github.com/krisshol/imt3501-Software-Security/SQLDatabase"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/util"
)

// DefaultHandler returns index.html.
func DefaultHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	expire := time.Now().AddDate(0, 0, 1)
	cookie, err := r.Cookie("session-id")
	if err != nil {											//TODO: Verify that it is okay to not check error
		id, _ := uuid.NewV4()
		cookie = &http.Cookie {Name: "session", Value: id.String(), Expires: expire}
		http.SetCookie(w, cookie)
	}

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to DefualtHandler\n")
	util.PrintURLAsSlice(r.URL.Path)

	if r.Method != "GET" { //Default handler is only GET. No Method switch.
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) >= 3 && parts[1] == "page" { // If there is 2 components in URL and the first one is "page". >= 3 Because there is a / at the end of the path as well.

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

	// fmt.Printf("User input accepted. Inserting message into db: \nmessage(first 20 chars): \"%s\"\nusername: %s\nparent: %d\n\n",
	// 	message.Message[0:20], message.Username, message.ParentMessage) // TODO: Remove test outprint.

	database.AddMessage(message)
	fmt.Fprint(w, "Message sent.\n")
}

// PostMessageHandler returns html page if GET, logs in user if POST.
func PostThreadHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to PostThreadHandler\n")
	util.PrintURLAsSlice(r.URL.Path) // TODO: Remove this outprint.

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

	var thread database.Thread
	thread.Name = r.FormValue("threadname")
	thread.Username = message.Username

	if !util.BasicValidate(thread.Name) {
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}

	// fmt.Printf("User input accepted. Inserting thread into db:\nthread name(first 20 chars): %s\nmessage(first 20 chars): \"%s\"\n\n",
	// 	thread.Name[0:20], message.Message[0:20]) // TODO: Remove test outprint.
	database.AddThread(thread, message)

	fmt.Fprint(w, "Message sent.\n")
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to CategoriesHandler\n")
	util.PrintURLAsSlice(r.URL.Path) // TODO: Remove this outprint.

	switch r.Method {
	case "GET":
		fmt.Fprint(w, util.FetchHTML("categories.html"))
		viewThreads := database.ShowThreads()
		w.Write([]byte(`<ul>`))
		for _, vThread := range viewThreads {
        fmt.Println("Id: ", vThread.Id)
				fmt.Println("Name: ", vThread.Name)
        fmt.Println("Username: ", vThread.Username)
        fmt.Println("")

				w.Write([]byte(`
  <li><h3>
	<input type="hidden" id="threadId" name="custId" value="`))
				w.Write([]byte(strconv.Itoa(vThread.Id)))
				w.Write([]byte(`"> `))
				w.Write([]byte(`
    <div id="textbox">
      	<p class="alignleft">`))
				w.Write([]byte(vThread.Name))
				w.Write([]byte(`</p>
	  	<p align=right><small><small>Username: `))
				w.Write([]byte(vThread.Username))
				w.Write([]byte(`</small></small></p>
    </div>
  </h3></li>
`))
		}
		w.Write([]byte(`
<ul>`))
		break

	case "POST":

		break
	default:
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}

//func Cookie (w http.ResponseWriter, r *http.Request) {


	//Alt1
	/*
	http.SetCookie(w, &http.Cookie) {
		Name: "my-cookie",
		Value: "some value",
	}*/

	//Alt2
	//cookie := http.Cookie {Name: "username", Value: "some value"}
	//http.SetCookie(w, &cookie)
//}
