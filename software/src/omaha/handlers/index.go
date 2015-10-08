package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"omaha/util"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	omahaDir := util.GetOmahaPath()
	templatePath := fmt.Sprintf("%s/templates/index.html", omahaDir)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Printf("Error parsing template at %s\n", templatePath)
		fmt.Println(err.Error())
		fmt.Fprint(w, "Something bad happened!")
		return
	}
	t.Execute(w, "/css/index.css")
}
