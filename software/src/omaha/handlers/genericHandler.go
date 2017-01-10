package handlers

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"omaha/database"
	"sync"
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

type systemInformation struct {
	Hash 		string 	`json:"hash"`
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
	log.Println("ServeHTTP")
	log.Println(req, w)
	log.Println(req.RequestURI)

	/*
	systemInfo := &systemInformation{}
	err := json.NewDecoder(req.Body).Decode(systemInfo)
	log.Println(systemInfo.Hash, err)
	*/

	sessionCookie, _ := req.Cookie("session")
	log.Println(sessionCookie)
	host, _, _ := net.SplitHostPort(req.RemoteAddr)
	uri := req.RequestURI
	log.Println(host)
	/*
		If we're at the same IP address as the server, the host should be ::1 (ipv6 localhost)
	*/
	switch {
		case host == "::1":
		case host == "127.0.0.1":
		case uri == "/loadControllers/":
			/*
				The request is coming from the server or is a loadController request at the login state request. Let it through by default.
			*/
		case sessionCookie == nil || !database.IsSessionHashValid(sessionCookie.Value):
			log.Println("case sessionCookie is", sessionCookie)
			//log.Println("this is the hash thing", database.IsSessionHashValid(sessionCookie.Value))
			redirectToLoginHandler(w, req)
			log.Println("Redirected to login")
			return
	}


	//log.Println("this is the hash thing", database.IsSessionHashValid(sessionCookie.Value))

	/*
		Generic stuff goes here
	*/
	var dbmutex = &sync.Mutex{}
	
	if req.Method == "GET" && this.GET != nil {
		dbmutex.Lock()
		this.GET(w, req)
		dbmutex.Unlock()
	} else if req.Method == "POST" && this.POST != nil {
		dbmutex.Lock()
		this.POST(w, req)
		dbmutex.Unlock()
	} else if req.Method == "PUT" && this.PUT != nil {
		dbmutex.Lock()
		this.PUT(w, req)
		dbmutex.Unlock()
	} else {
		http.Error(w, "GenericHandler error", http.StatusInternalServerError)
		log.Fatalf("No handler specified for the request %s", req)
	}
}
