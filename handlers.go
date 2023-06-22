package main

import (
	"fmt"
	"net/http"
	"text/template"
)

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
