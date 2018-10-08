# IMT3501 Software Security Project - Forum application

## Members
| Name             | Studentno |
| ---------------- | --------- |
| Jone Skaara      | 473181    |
| Kristian S. Holm | 473114    |
| Olav H. Hoggen   | 473138    |
| Martin Kvalvåg   | 473144    |
| Magnus Bringe    | 473155    |

## Description
This is the assignment in IMT3501 Software Security at NTNU in Gjøvik Norway. The task is to create a web forum with a concept of users, threads, categories, messages, and message replies.
Users can sign up then login and keep the session for their time browsing the forum, post messages and have other users up- and downvote your messages earning you reputation points (that is not completed).
There is a strong focus on security in the application.

## Components

### Golang Server
- The Golang server takes requests from the different HTML documents. We have a
- DefaultHandler that returns the index document with a list of categories, it can also return any generic html document using the /page prefix. 
- SignupHandler that returns the signup document or register new user.
- LoginHandler that returns the login html document. It logs user in, creating session cookies.
- LogoutHandler that logs a user out by deleting their sessions. 
- MessageHandler displays the content of a thread. Create new message if post request.
- ThreadHandler displays content of a category. Create new thread if post request.
- CategoryHandler generates a list of threads and sends it to the user.

### HTML Websites
In the html documents we use basic html with form actions which contacts the go handlers in the golang server. We have a htmlGeneration package that generates threads, categories and messages from the database.

### Database 
We have made a struct for each of our database types, add functions, get functions, show functions and delete functions. The add functions use prepare statements to insert into the database. The get functions is used to retrieve the user, threads, messages or replies. Show is used to generate a slice of category and treads.

## Dependencies
- github.com/go-sql-driver/mysql
    My SQL driver for golang.
- github.com/tkanos/gonfig
    Reading for json files for system configuration.
- github.com/nu7hatch/gouuid
    Generating session ids.
- golang.org/x/crypto/scrypt
    Generating salt and hashing of passwords



## Installation
Prerequisite for installation on any OS is that golang 1.11 or later is installed and set up on the intended server computer. 
Link to Golang download: 
https://golang.org/dl/


#### Config.json setup (All operating systems)
Time to fill in all parameters the server uses to communicate with the database and requesting clients.  
##### Syntax
``` json
	"Port" :        	int, 	# The port the golang server will be using.
	"Address" :     	string, # The address of the host computer on the local network.

	"DatabasePort" :	int, 	# The port the SQL database will be using. 
	"DatabaseAddress" :  	string, # The address to the database that the golang server will be using.
	"DatabaseDatabase" : 	string,	# The database within the SQL databse server to use.
	"DatabaseUser" :	string, # The user the golang server will be logging into the database with.
	"DatabasePassword" :	string, # The password for that databse user.

	"HtmlPath" :   		string  # The relative or absolute path to the html folder in the repository containing all the html docs.
```
##### Example
``` json
{
   "Port" :            	5000,
   "Address" :         	"127.0.0.1",

   "DatabasePort" :    	3306,
   "DatabaseAddress":  	"21.65.27.111",
   "DatabaseDatabase" :	"forumdatabase",
   "DatabaseUser" :    	"forumuser",
   "DatabasePassword" :	"password",

   "HtmlPath" :   	"/home/name/golang/src/github.com/krisshol/imt3501-Software-Security/html/"
}
```



### Linux
```bash
    # cd $GOPATH
    mkdir -p src/github.com/krisshol && cd src/github.com/krisshol
    git clone https://github.com/krisshol/imt3501-Software-Security.git
    cd imt3501-Software-Security
    go get github.com/tkanos/gonfig
    go get github.com/go-sql-driver/mysql
    go get github.com/nu7hatch/gouuid
    go get -u golang.org/x/crypto/scrypt
    cp docs/config.json.example cmd/forumServer/config/config.json
    # See  Config.json setup above.
    go build ./cmd/forumServer
     
    ./forumServer
```
### Windows
```powershell
  # In powershell or cmd
  # Navigate to $GOPATH
  mkdir src
  mkdir github.com
  mkdir krisshol
  git clone https://github.com/krisshol/imt3501-Software-Security.git
  go get github.com/tkanos/gonfig
  go get github.com/go-sql-driver/mysql
  go get github.com/nu7hatch/gouuid
  go get -u golang.org/x/crypto/scrypt
  cd src\github.com\krisshol\imt3501-Software-Security\
  copy docs\config.json.example cmd\forumserver\config\config.json
  # Setup config file, see start of installation guide
  mkdir bin
  cd bin
  go build ..\cmd\forumServer
  forumserver.exe
```
  
  
  


local db:
``` bash
    # basic setup login in with root or other authorised user
    mysql -u root -p
    create database forum
    # go to $GOPATH/src/github.com/krisshol/imt3501-Software-Security/
    mysql -u root -p forum < create-db.sql
```
