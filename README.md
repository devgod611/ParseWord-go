# gobackend
This is a simple golang backend and simple web page app that parse book text content(show ten most-used words)

#Install Dependencies
- init go module
$ go mod init 

- To load .env file, install this package
$ go get github.com/joho/godotenv

- For the automatic compile & run, intall air package
$ go install github.com/cosmtrek/air@latest

- install gin & gorm package
$ go get github.com/gin-gonic/gin github.com/jinzhu/gorm

- install static
$ go get -u github.com/gin-gonic/contrib/static

# How to run
- First download go dependencies
$ go mod download

- How to run project 
$ go run main.go

Opne http://localhost:8080/ in your browser
Type the text into the input field and press "PARSE" button to show the list of most used words

![This is an image](https://github.com/ittechman101/gobackend/blob/main/screenshot.PNG)
