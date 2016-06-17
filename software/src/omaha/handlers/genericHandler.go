package handlers

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"omaha/database"
)

type GenericHandler struct {
	GET  func(http.ResponseWriter, *http.Request)
	PUT  func(http.ResponseWriter, *http.Request)
	POST func(http.ResponseWriter, *http.Request)
}

type genericResponse struct {
	Success bool   `json:"success"`
	Err     string `json:"err"`
}

/*
	getGenericSuccessResponse is a private function that returns a JSON object that contains a success boolean.
*/
func getGenericSuccessResponse() []byte {
	response, _ := json.Marshal(genericResponse{Success: true})
	return response
}

/*
	getGenericErrorResponse is a private function that returns a JSON object that contains an error message.
*/
func getGenericErrorResponse(err string) []byte {
	response, _ := json.Marshal(genericResponse{Err: err})
	return response
}

/*
	ServeHTTP is the function that starts the server.
*/
func (this GenericHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	sessionCookie, _ := req.Cookie("session")
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	/*
		If we're at the same IP address as the server, the host should be ::1 (ipv6 localhost)
	*/
	switch {
		case host == "::1":		// this currently only works on chrome for some reason
			/*
				The request is coming from the server. Let it through by default.
			*/
		case sessionCookie == nil || !database.IsSessionHashValid(sessionCookie.Value):
			log.Println("this is the hash thing", database.IsSessionHashValid(sessionCookie.Value))
			redirectToLoginHandler(w, req)
			log.Println("Redirected to login")
			return
	}
	log.Println("this is the hash thing", database.IsSessionHashValid(sessionCookie.Value))

	/*
		Generic stuff goes here
	*/
		
	if req.Method == "GET" && this.GET != nil {
		this.GET(w, req)
	} else if req.Method == "POST" && this.POST != nil {
		this.POST(w, req)
	} else if req.Method == "PUT" && this.PUT != nil {
		this.PUT(w, req)
	} else {
		http.Error(w, "GenericHandler error", http.StatusInternalServerError)
		log.Fatalf("No handler specified for the request %s", req)
	}
}
