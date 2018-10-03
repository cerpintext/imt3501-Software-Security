package app

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	

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
	// setup cookie for deployment
         // see http://golang.org/pkg/net/http/#Request.Cookie

         // we will try to drop the cookie, if there's error
         // this means that the same cookie has been dropped
         // previously and display different message
         c, err := r.Cookie("timevisited") //

         expire := time.Now().AddDate(0, 0, 1)

         cookieMonster := &http.Cookie{
                 Name:  "timevisited",
                 Expires: expire,
                 Value: strconv.FormatInt(time.Now().Unix(), 10),
         }

         // http://golang.org/pkg/net/http/#SetCookie
         // add Set-Cookie header
         http.SetCookie(w, cookieMonster)

         if err != nil {
                 w.Write([]byte("Welcome! first time visitor!"))
         } else {
                 lasttime, _ := strconv.ParseInt(c.Value, 10, 0)
                 html := "Hey! Hello again!, your last visit was at "
                 html = html + time.Unix(lasttime, 0).Format("15:04:05")
                 w.Write([]byte(html))
         }
	/*
	expire := time.Now().AddDate(0, 0, 1)
	cookie := &http.Cookie {Name: "username", Value: "some value", Expires: expire}
	http.SetCookie(w, cookie)*/
	/*	
	http.SetCookie(w, &http.Cookie {
		Name: "my-cookie",
		Value: "some value",
	})*/

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
	fmt.Print("Received a request to SignUpHandler\n")
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
	fmt.Print("Received a request to SignUpHandler\n")
	util.PrintURLAsSlice(r.URL.Path)

	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}

	r.ParseForm()
	var message database.Message
	message.Message = r.FormValue("message")
	message.Username = r.FormValue("username")

	parent, err := strconv.Atoi(r.FormValue("parentmessage"))
	if err != nil {
		fmt.Printf("Failed to parse message.parentmessage, got: %s\n\n\n", r.FormValue("parentmessage"))
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}
	message.ParentMessage = parent

	if !util.BasicValidate(message.Message, -1, config.MAX_MESSAGE_LENGTH) ||
		!util.BasicValidate(message.Username) ||
		message.ParentMessage < 0 {

		fmt.Fprint(w, "Message was invalid")
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		fmt.Print("Message rejected.\n\n")
		return
	}

	fmt.Printf("User input accepted. Inserting message into db: \nmessage(first 20 chars): \"%s\"\nusername: %s\nparent: %d\n\n",
		message.Message[0:20], message.Username, parent) // TODO: Remove test outprint.

	database.OpenDB()
	database.AddMessage(message)
	fmt.Fprint(w, "Message sent.\n")
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
