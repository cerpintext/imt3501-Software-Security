package htmlGeneration

import (
	"fmt"
	"strconv"

	database "github.com/krisshol/imt3501-Software-Security/SQLDatabase"
	"github.com/krisshol/imt3501-Software-Security/cmd/forumServer/config"
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

	htmlDoc += "<form action=\"/thread/\" method=\"post\">Thread Name <input type=\"text\" name=\"threadname\"> Thread Description <input type=\"text\" name=\"message\"><input type=\"hidden\" name=\"parentmessage\" value =\"-1\"><input type=\"hidden\" name=\"threadid\" value =\"-1\"><input type=\"hidden\" name=\"categoryname\" value =\"" + category + "\"><input type=\"submit\" value =\"Create thread\"></input></form>\n"

	database.OpenDB()
	viewThreads := database.ShowThreads(category)
	htmlDoc += "<ul>"
	fmt.Println("Displaying all threads in category: " + category)

	for _, vThread := range viewThreads {
		fmt.Printf("ThreadId: %d \tName:%s \n", vThread.Id, vThread.Name)

		htmlDoc += "<li><h3>"
		htmlDoc += "<input type=\"hidden\" id=\"threadId\" name=\"custId\" value=\""
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
	viewMessages := database.GetThread(database.Thread{threadId, "", ""})

	htmlDoc += "<ul>\n"
	fmt.Printf("GenerateMessagesList(): Messeage count in thread(%d): %d\n\tMod:%t\n", threadId, len(viewMessages), moderator)

	for _, vMessage := range viewMessages {
		fmt.Printf("MessageID: %d \tMessageText:%s \n", vMessage.Id, vMessage.Message) // TODO: Remove bad use of printf.

		htmlDoc += "<b>" + vMessage.Username + "</b>\n"
		htmlDoc += "<p>" + vMessage.Message + "</p>\n"

		viewComments := database.GetReplies(vMessage.Id)
		for _, vComment := range viewComments {
			htmlDoc += "<hr>\n"

			htmlDoc += "<b style=\"margin-left:" + fmt.Sprint(config.COMMENT_INTENT) + "px\">" + vComment.Username + "</b>\n"
			htmlDoc += "<p style=\"margin-left:" + fmt.Sprint(config.COMMENT_INTENT) + "px\">" + vComment.Message + "</p>\n"
			if moderator || vMessage.Username == username {
				htmlDoc += "<form style=\"margin-left:" + fmt.Sprint(config.COMMENT_INTENT) + "px\" action=\"/message/\" method=\"delete\"><input type=\"hidden\" name=\"messageid\" value =\"" + fmt.Sprint(vComment.Id) + "\"><input type=\"submit\" value =\"Remove\"></form>\n"
			}
		}

		if len(username) > 0 {
			htmlDoc += "<form action=\"/message/\" method=\"post\"><input type=\"text\" name=\"message\"><input type=\"hidden\" name=\"parentmessage\" value =\"" + fmt.Sprint(vMessage.Id) + "\"><input type=\"hidden\" name=\"threadid\" value =\"-1\"><input type=\"submit\" value =\"Comment\"></input></form>\n"
		}
		if moderator || vMessage.Username == username {
			htmlDoc += "<form action=\"/message/\" method=\"delete\"><input type=\"hidden\" name=\"messageid\" value =\"" + fmt.Sprint(vMessage.Id) + "\"><input type=\"submit\" value =\"Remove\"></form>\n"
		}
		htmlDoc += "<hr>\n"
		htmlDoc += "<br>\n"
	}
	htmlDoc += "<p>Post a message to this thread:</p>"
	htmlDoc += "<form action=\"/message/\" method=\"post\"><input type=\"text\" name=\"message\"><input type=\"hidden\" name=\"parentmessage\" value =\"-1\"><input type=\"hidden\" name=\"threadid\" value =\"" + fmt.Sprint(threadId) + "\"><input type=\"submit\" value =\"Post message\"></input></form>\n"

	htmlDoc += "</ul>\n"

	return htmlDoc
}
