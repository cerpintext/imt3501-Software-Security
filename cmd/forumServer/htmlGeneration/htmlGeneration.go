package htmlGeneration

import (
	"fmt"
	"strconv"

	database "github.com/krisshol/imt3501-Software-Security/SQLDatabase"
)

func GenerateCategoryList() string {
	var htmlDoc string
	htmlDoc += "<ul>"

	database.OpenDB()
	viewCategories := database.ShowCategories()

	for _, vCategory := range viewCategories {
		htmlDoc += "<a href=\"/category/" + vCategory.Name + "\">" + vCategory.Name + "</a>"
		htmlDoc += "<br>"
	}

	htmlDoc += "/<ul>"

	return htmlDoc
}

// GenerateTreadList gets all threads in the db for a certain category, and formats them to html text.
func GenerateTreadList(category string) string {

	var htmlDoc string

	database.OpenDB()
	viewThreads := database.ShowThreads(category)
	htmlDoc += "<ul>"
	fmt.Println("Displaying all threads in category: " + category)

	for _, vThread := range viewThreads {
		fmt.Printf("ThreadId: %d \tName:%s \n", vThread.Id, vThread.Name)

		htmlDoc += "<li><h3>"
		htmlDoc += "<input type=\"hidden\" id=\"threadId\" name=\"threadId\" value=\""
		htmlDoc += strconv.Itoa(vThread.Id) + "\">\n"
		htmlDoc += "<div id=\"textbox\">\n<a href=\"/thread/" + fmt.Sprint(vThread.Id) + "\" class=\"alignleft\">" + vThread.Name + "</a>"
		htmlDoc += "<p align=right><small><small>Username: " + vThread.Username
		htmlDoc += "</small></small></p>\n</div>\n</h3></li>\n"
	}

	htmlDoc += "/<ul>"
	return htmlDoc
}

func GenerateMessagesList(threadId int, username string, moderator bool) string {
	var htmlDoc string

	database.OpenDB()
	viewMessages := database.GetThread(database.Thread{60, "", ""})
	// viewMessages := database.ShowMessages(threadId)

	htmlDoc += "<ul>"
	fmt.Printf("GenerateMessagesList(): Messeage count in thread(%d): %d\n", threadId, len(viewMessages))

	for _, vMessage := range viewMessages {
		fmt.Printf("MessageID: %d \tMessageText:%s \n", vMessage.Id, vMessage.Message) // TODO: Remove bad use of printf.

		htmlDoc += "<b>" + vMessage.Username + "</b>"
		htmlDoc += "<input type=\"hidden\" id=\"messageId\" name=\"messageId\" value=\""
		htmlDoc += strconv.Itoa(vMessage.Id) + "\">\n"
		htmlDoc += "<p>" + vMessage.Message + "</p>"

		if moderator || vMessage.Username == username {
			htmlDoc += "<form action=\"/message/\" method=\"delete\"><input type=\"submit\" value =\"Remove\"></form>"
		}
		htmlDoc += "<br>"
	}

	htmlDoc += "</ul>"

	return htmlDoc
}
