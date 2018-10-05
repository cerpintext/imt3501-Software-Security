package app

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/nu7hatch/gouuid"

	"github.com/krisshol/imt3501-Software-Security/SQLDatabase"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/htmlGeneration"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/util"
)

// DefaultHandler returns index.html.
func DefaultHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to DefualtHandler\n")

	if r.Method != "GET" { //Default handler is only GET. No Method switch.
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) >= 3 && parts[1] == "page" { // If there is 2 components in URL and the first one is "page". >= 3 Because there is a / at the end of the path as well.

		fmt.Fprint(w, util.FetchHTML(parts[2]+".html"))

	} else {

		fmt.Fprint(w, util.FetchHTML("index.html"))
		fmt.Fprint(w, htmlGeneration.GenerateCategoryList())
	}

}

//SignInHandler returns html page if GET, registers new user if POST.
func SignUpHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to SignUpHandler\n")

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
		database.AddUser(user)                                                                  // Send struct to db.

		fmt.Fprint(w, "All good, welcome to the team "+userName+"! :D<br> <a href=\"/login/\">Login with your new user</a>")
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

	switch r.Method {
	case "GET":
		fmt.Fprint(w, util.FetchHTML("login.html"))
		break

	case "POST":
		r.ParseForm()
		// Get fields from form:
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Validate input:
		if !util.BasicValidate(username) ||
			!util.BasicValidate(password, config.MIN_PASSWORD_LENGTH) {

			w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
			fmt.Fprint(w, util.FetchHTML("login.html"))
			fmt.Fprint(w, "Sorry, some of that was wrong.")
			return
		}
		fmt.Print("User input accepted.\n")

		if util.IsLoggedIn(r) { // The user has registered session and  still has their cookie(not expired).

			w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
			fmt.Fprint(w, "You are already logged in.")
			fmt.Printf("User: %s Is already logged in. Aboring login.\n\n", username)
			return

		}

		database.OpenDB()
		user, err := database.GetUser(username)

		if err != nil { // Check if user in db.
			w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
			fmt.Fprint(w, util.FetchHTML("login.html"))
			fmt.Fprint(w, "Sorry, some of that was wrong.")
			fmt.Printf("User: %s doesn't exist. Aboring login.\n\n", username)
			return
		}
		if user.PasswordHash != password { // Check if password matches.
			w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
			fmt.Fprint(w, util.FetchHTML("login.html"))
			fmt.Fprint(w, "Sorry, some of that was wrong.")
			fmt.Print("Wrong password. Aboring login.\n\n")
			return
		}

		fmt.Printf("User exists. Password match. Attempting to store session id for user.")

		expire := time.Now().AddDate(0, 0, config.SESSION_EXPIRE_DAYS)
		cookie, err := r.Cookie("session")
		if err != nil { //TODO: Verify that it is okay to not check error
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{Name: "username", Value: username, Path: "/", Expires: expire}
			http.SetCookie(w, cookie)
			cookie = &http.Cookie{Name: "session", Value: id.String(), Path: "/", Expires: expire}
			http.SetCookie(w, cookie)
		}
		config.SessionMap[username] = cookie.Value
		fmt.Printf("Session generated:\n%s.\n \n", cookie.Value)

		fmt.Fprint(w, "Logged in successfully<br> <a href=\"/\">Back to front page</a>")
		break
	default:
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}

}

// MessageHandler returns html page if GET, logs in user if POST.
func MessageHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to PostMessageHandler\n")

	switch r.Method {
	case "GET":
		util.FetchHTML("message.html")
		break
	case "POST":
		message, err := util.ValidateMessage(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
			fmt.Fprint(w, err)
			return
		}

		fmt.Print("User input accepted. Inserting message into db\n")
		database.AddMessage(message)
		fmt.Fprint(w, "Message sent.<br> <a href=\"/\">Back to front page</a>")
		break

	case "DELETE":
		break

	default:
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}
}

// MessageHandler returns html page if GET, logs in user if POST.
func ThreadHandler(w http.ResponseWriter, r *http.Request) { // Default request handler handles domain/ requests.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to ThreadHandler\n")

	switch r.Method {
	case "GET":
		// Parse input from URL.
		parts := strings.Split(r.URL.Path, "/")
		threadId, err := strconv.Atoi(parts[2])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
			fmt.Fprint(w, "No thread with that ID.")
			return
		}

		moderator := false
		username := ""

		if util.IsLoggedIn(r) {
			cookieUsername, _ := r.Cookie("username")
			username = cookieUsername.Value

			database.OpenDB()
			user, _ := database.GetUser(username)
			if user.Role > 0 {
				moderator = true
			}
		}

		fmt.Println("Displaying thread: " + fmt.Sprint(threadId))
		fmt.Fprint(w, util.FetchHTML("thread.html"))
		fmt.Fprint(w, htmlGeneration.GenerateMessagesList(threadId, username, moderator))
	case "POST":

		message, err := util.ValidateMessage(r)
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
		database.AddThread(thread, message)

		fmt.Fprint(w, "Message sent.<br>	<a href=\"/\">Back to front page</a>")
		break

	default:
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}

}

func CategoryHandler(w http.ResponseWriter, r *http.Request) { // Generates a list of categories and sends it the user.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to CategoriesHandler\n")
	util.PrintURLAsSlice(r.URL.Path)
	parts := strings.Split(r.URL.Path, "/")
	category := parts[2]
	fmt.Println("Displaying category: " + category)

	switch r.Method {
	case "GET":
		fmt.Fprint(w, util.FetchHTML("category.html"))
		fmt.Fprint(w, htmlGeneration.GenerateTreadList(category))
		break

	case "POST":

		break
	default:
		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.
		return
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) { // Logs user out by setting their cookie to be expired so they get deleted.

	w.Header().Set("Content-Type", "text/html") // The response will be an html document.
	fmt.Print("Received a request to Logout handler\n")

	if r.Method == "POST" {

		cookieUsername, err := r.Cookie("username")
		if err != nil { //The user has registered session but their cookie is expired.
			return
		} else {
			delete(config.SessionMap, cookieUsername.Value) // Delete stored session id.
		}

		expire := time.Now().AddDate(0, 0, -1)
		cookie := &http.Cookie{Name: "username", Value: "LOGGEDOUT", Path: "/", Expires: expire}
		http.SetCookie(w, cookie)
		cookie = &http.Cookie{Name: "session", Value: "NOSESSION", Path: "/", Expires: expire}
		http.SetCookie(w, cookie)

		fmt.Fprint(w, "Logged out.<br>	<a href=\"/\">Back to front page</a>")

	} else {

		w.WriteHeader(http.StatusBadRequest) // Bad input give errorcode 400 bad request.

	}

}
