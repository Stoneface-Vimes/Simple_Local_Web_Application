//The below code is written following Section 3 of Building Modern Web Applications with Go (Golang)
//URL: https://www.udemy.com/course/building-modern-web-applications-with-go/
//This program will be iterated upon through the lessons in Section 3

//A simple program that listens on localhost port 8080, and sends the message Hello, world!
//when the page is accessed. The number of bytes written is output in the terminal

package main

import (
	"fmt"
	"net/http"
)

// PortNumber is set as a const as we do not want the value of the variable to change at any point
const portNumber = ":8080"

// In order for a function to respond to a request from a web browser, it has to handle two parameters:
// A response writer and a request. The request is always a pointer

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)

	//Fprintf will format a string and write it to w. The string is defined after w.
	//Fprintf also returns the amount of bytes used as an interger and an error message, if any
	//Since we're not using those returned values, they've been set to Blank Identifiers

	_, _ = fmt.Fprintf(w, "This is the about page and 2 + 2 is %d", sum)
}

// Functions that start with a capital letter are accesbile outside their package. For best practices,
// if a function is only called from inside it's package, it should start with a lower case letter

// addValues adds and returns the passed values of x and y
func addValues(x, y int) int {
	return x + y
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
