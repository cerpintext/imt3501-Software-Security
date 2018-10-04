package htmlGeneration

import (
	"fmt"
	"strconv"

	database "github.com/krisshol/imt3501-Software-Security/SQLDatabase"
)

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
		htmlDoc += "<input type=\"hidden\" id=\"threadId\" name=\"custId\" value=\""
		htmlDoc += strconv.Itoa(vThread.Id) + "\">\n"
		htmlDoc += "<div id=\"textbox\">\n<p class=\"alignleft\">" + vThread.Name + "</p>"
		htmlDoc += "<p align=right><small><small>Username: " + vThread.Username
		htmlDoc += "</small></small></p>\n</div>\n</h3></li>\n"
	}

	htmlDoc += "<ul>"
	return htmlDoc
}

func GenerateCategoryList() string {
	var htmlDoc string
	htmlDoc += "<ul>"

	database.OpenDB()
	viewCategories := database.ShowCategories()

	for _, vCategory := range viewCategories {
		htmlDoc += "<a href=\"/categories/" + vCategory.Name + "\">" + vCategory.Name + "</a>"
		htmlDoc += "<br>"
	}

	htmlDoc += "<ul>"

	return htmlDoc
}
