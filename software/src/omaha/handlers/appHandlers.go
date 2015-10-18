package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"omaha/util"
)

func AppHandler(w http.ResponseWriter, r *http.Request) {
	omahaDir := util.GetOmahaPath()
	templatePath := fmt.Sprintf("%s/templates/app.html", omahaDir)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Printf("Error parsing template at %s\n", templatePath)
		fmt.Println(err.Error())
		fmt.Fprint(w, "Something bad happened!")
		return
	}
	t.Execute(w, nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	omahaDir := util.GetOmahaPath()
	templatePath := fmt.Sprintf("%s/templates/login.html", omahaDir)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Printf("Error parsing template at %s\n", templatePath)
		fmt.Println(err.Error())
		fmt.Fprint(w, "Something bad happened!")
		return
	}
	t.Execute(w, nil)
}
