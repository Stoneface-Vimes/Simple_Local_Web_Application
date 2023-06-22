package handlers

import (
	"Simple_Local_Web_Application/pkg/render"
	"net/http"
)

// In order for a function to respond to a request from a web browser, it has to handle two parameters:
// A response writer and a request. The request is always a pointer

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
