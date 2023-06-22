//A simple program that listens on localhost port 8080, and sends the message Hello, world!
//when the page is accessed. The number of bytes written is output in the terminal

package main

import (
	"fmt"
	"net/http"
)

func main() {
	//HandleFunc defines what is sent to the page, when the page is accessed. "/" with nothing else
	//refers to the homepage of the website in question. Below, a function is defined inside the
	//HandleFunc statement
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//A value of "Hello, world!" is written to w. The function fmt.Fprintf returns the
		//number of bytes written and an error code (if any), so they must be assigned to either
		//variables or underscores (_, aka "Blank Identifier")
		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Number of bytes written: %d \n", n)
	})
	//Establishes a server and listens at port 8080. A handler can be specified; when nil is used
	//a default handler is chosen
	_ = http.ListenAndServe(":8080", nil)
}
