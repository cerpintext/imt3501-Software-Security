# IMT3501 Software Security Project - Forum application

## Members
Jone Skaara (Student no: 473181)  
Kristian S. Holm (Student no: 473114)  
Olav H. Hoggen (Student no: 473138)
Martin Kvalvåg (Student no: 473144)

## Description
This is the assignment in IMT3501 Software Security at NTNU in Gjøvik Norway. The task is to create a web forum with a concept of users, threads, categories, messages, and message replies.
Users can sign up then login and keep the session for their time browsing the forum, post messages and have other users up- and downvote your messages earning you reputation points. 
There is a string focus on security in the application.

## Components 

## Installation
```bash
	git clone git@github.com:krisshol/imt3501-Software-Security.git

	.
	.

```

local db:
```
    #basic setup first eg. add root user.
    mysql -u root -p
    create database forum	
    mysql -u root -p forum < create-db.sql 
```
