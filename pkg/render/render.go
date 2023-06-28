package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// The renderTemplate function is used to render templates on our various pages. Requires a response (w) and the
// type of template that is being rendered (html, tmpl, etc)
func RenderTemplate(w http.ResponseWriter, tmpl string) {

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
