//The below code is written following Section 3 of Building Modern Web Applications with Go (Golang)
//URL: https://www.udemy.com/course/building-modern-web-applications-with-go/
//This program will be iterated upon through the lessons in Section 3

//A simple program that listens on localhost port 8080, and sends the message Hello, world!
//when the page is accessed. The number of bytes written is output in the terminal

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// PortNumber is set as a const as we do not want the value of the variable to change at any point
const portNumber = ":8080"

// In order for a function to respond to a request from a web browser, it has to handle two parameters:
// A response writer and a request. The request is always a pointer

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

//The renderTemplate function is used to render templates on our various pages. Requires a response (w) and the
//type of template that is being rendered (html, tmpl, etc)

func renderTemplate(w http.ResponseWriter, tmpl string) {

	//ParsedTemplate is used to store the template that will be rendered. The template that is stored
	//will depend on the string provided when renderTemplate is called
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

// main is the main application function
func main() {

	//Each HandleFunc specifies an address and a function to call when that address is accessed

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	//Prints the port in use by the application to the terminal
	fmt.Printf("Starting application on port %s", portNumber)

	//Establishes a server and listens at the port defined at portNumber. A handler can be specified; when nil is used
	//a default handler is chosen
	_ = http.ListenAndServe(portNumber, nil)
}
