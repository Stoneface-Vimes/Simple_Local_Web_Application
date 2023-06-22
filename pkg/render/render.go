package render

import (
	"fmt"
	"net/http"
	"text/template"
)

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
