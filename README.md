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
Users can sign up then login and keep the session for their time browsing the forum, post messages and have other users up- and downvote your messages earning you reputation points. 
There is a string focus on security in the application.

## Components 

### Golang Server

### HTML Websites

### Database Handler

### MYSQL Database

## Dependencies
    - github.com/subosito/gotenv
        Gotenv is a very small library for reading .env files for golang programs.
    - github.com/go-sql-driver/mysql
        My SQL driver for golang.


## Installation
### Linux
```bash
	git clone git@github.com:krisshol/imt3501-Software-Security.git
    cd imt3501-Software-Security/bin
    go get github.com/subosito/gotenv
    go get github.com/go-sql-driver/mysql
    go build ../cmd/forumServer
    cp ../docs/envExample .env
    # Fill inn missing fields in .env file. Like PORT and IP.
    .
    .        
    .
    ./forumServer   # Must be ran from bin since the path is relative to terminals working dir, not the executable's location.

```

local db:
```
    #basic setup first eg. add root user.
    mysql -u root -p
    create database forum	
    mysql -u root -p forum < create-db.sql 
```
