package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	//"net"
	"net/http"
	"omaha/database"
	"omaha/util"
)

type loginPostRequest struct {
	Username  string  `json:"username"`
	Password  string  `json:"password"`
}

type loginPostResponse struct {
	Hash  		string  `json:"hash"`
	Level 		int			`json:"level"`
	SpeakerID int			`json:"speakerid"`
	ZoneID		int			`json:"zoneid"`
}

type accountCreationRequest struct {
	Level			int		  `json:"level"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Name			string  `json:"name"`
	Email			string	`json:"email"`
	Phone 		string	`json:"phone"`
	SpeakerID int			`json:"speakerid"`
	ZoneID		int			`json:"zoneid"`
}

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

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LOGIN")

	// uncomment this stuff later 
	//host, _, _ := net.SplitHostPort(r.RemoteAddr)
	/*if host == "::1" {
		w.Write(getGenericSuccessResponse())
		return
	}*/
	loginRequest := &loginPostRequest{}
	err := json.NewDecoder(r.Body).Decode(loginRequest)
	log.Println(*loginRequest)
	if err != nil {
		w.Write(getGenericErrorResponse(err.Error()))
		return
	} else {
		hash, err := database.LoginAccount(loginRequest.Username, loginRequest.Password)
		var level = database.GetLevelOfAccount(loginRequest.Username)
		var speakerID = database.GetSpeakerForAccount(loginRequest.Username)
		var zoneID = database.GetZoneForAccount(loginRequest.Username)

		log.Println("This is ths level that you are looking for ", level)
		if err != nil {
			w.Write(getGenericErrorResponse(err.Error()))
			return
		}

		response := &loginPostResponse{hash, level, speakerID, zoneID}
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

func AccountCreationHandler(w http.ResponseWriter, r *http.Request) {
	accountRequest := &accountCreationRequest{}
	err := json.NewDecoder(r.Body).Decode(accountRequest)
	if err != nil {
		log.Printf("AccountCreationHandler json decoding error: %s\n", err)
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	// Change this data based on something later
	err = database.CreateAccount(accountRequest.Level, accountRequest.Username, accountRequest.Password, accountRequest.Name, accountRequest.Email, accountRequest.Phone, accountRequest.SpeakerID, accountRequest.ZoneID)
	if err != nil {
		w.Write(getGenericErrorResponse(err.Error()))
	} else {
		w.Write(getGenericSuccessResponse())
	}
}
