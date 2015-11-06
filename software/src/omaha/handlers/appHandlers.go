package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"omaha/database"
	"omaha/util"
)

func AppHandler(w http.ResponseWriter, r *http.Request) {
	omahaDir := util.GetOmahaPath()
	templatePath := fmt.Sprintf("%s/templates/app.html", omahaDir)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Error parsing template at %s\n", templatePath)
		log.Println(err)
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	t.Execute(w, nil)
}

func redirectToLoginHandler(w http.ResponseWriter, r *http.Request) {
	omahaDir := util.GetOmahaPath()
	templatePath := fmt.Sprintf("%s/templates/loginFromRedirect.html", omahaDir)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Error parsing template at %s\n", templatePath)
		log.Println(err)
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	t.Execute(w, nil)
}

type loginPostRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginPostResponse struct {
	Hash string `json:"hash"`
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LOGIN")
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	if host == "::1" {
		w.Write(getGenericSuccessResponse())
		return
	}
	loginRequest := &loginPostRequest{}
	err := json.NewDecoder(r.Body).Decode(loginRequest)
	if err != nil {
		w.Write(getGenericErrorResponse(err.Error()))
		return
	} else {
		hash, err := database.LoginAccount(loginRequest.Username, loginRequest.Password)
		if err != nil {
			w.Write(getGenericErrorResponse(err.Error()))
			return
		}
		response := &loginPostResponse{hash}
		var responseObj []byte
		responseObj, err = json.Marshal(response)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(responseObj)
	}
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	omahaDir := util.GetOmahaPath()
	templatePath := fmt.Sprintf("%s/templates/login.html", omahaDir)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Printf("Error parsing template at %s\n", templatePath)
		log.Println(err)
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	t.Execute(w, nil)
}
