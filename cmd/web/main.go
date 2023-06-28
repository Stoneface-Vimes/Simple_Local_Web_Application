//The below code is written following Section 3 of Building Modern Web Applications with Go (Golang)
//URL: https://www.udemy.com/course/building-modern-web-applications-with-go/
//This program will be iterated upon through the lessons in Section 3

package main

import (
	"Simple_Local_Web_Application/pkg/handlers"
	"fmt"
	"net/http"
)

// PortNumber is set as a const as we do not want the value of the variable to change at any point
const portNumber = ":8080"

// main is the main application function
func main() {

	//Each HandleFunc specifies an address and a function to call when that address is accessed

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	//Prints the port in use by the application to the terminal
	fmt.Printf("Starting application on port %s\n", portNumber)

	//Establishes a server and listens at the port defined at portNumber. A handler can be specified; when nil is used
	//a default handler is chosen
	_ = http.ListenAndServe(portNumber, nil)
}
