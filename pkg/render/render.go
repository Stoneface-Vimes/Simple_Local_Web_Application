package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// The renderTemplate function is used to render templates on our various pages. Requires a response (w) and the
// type of template that is being rendered (html, tmpl, etc)
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {

	//ParsedTemplate is used to store the template that will be rendered. The template that is stored
	//will depend on the string provided when renderTemplate is called
	//Added the base layout tmpl file to the files being parsed
	//Currently, every time you visit a page the template must be read from disk. In the future,
	//the parsed templates will be saved to a local variable
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}

// Package level Template Cache, stores parsed templates to avoid redundant disc reads
var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//We want to check to see if we already have the template in our cache
	//If the variable t (defined in RenderTemplateTest) is in the map, inMap is True
	_, inMap := tc[t]

	if !inMap {
		//If inMap is false, we need to create the template and add it to the Template Cache
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		//IF inMap is true, we have the template in the Template Cache
		log.Println("using cached template")
	}

	//Assigns tmpl the value we recieved from the createTemplateCache function (t, string)
	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

// Creates and returns a templace cache assigned to a string "t"
// Assings all templates we will be using to a slice of stringsnamed "templates".
// The stored templates are then parsed and assigned to tmpl. If no errors
// are produced during parsing, a string is returned
func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	//Parse the template and return an error if an error occurs
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	//IF no errors were produced when the template was parsed, add parsed template to Template Cache
	tc[t] = tmpl

	return nil

}
