package handlers

import (
	"encoding/json"
	"log"
	"net/http"
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

func getGenericSuccessResponse() []byte {
	response, _ := json.Marshal(genericResponse{Success: true})
	return response
}

func getGenericErrorResponse(err string) []byte {
	response, _ := json.Marshal(genericResponse{Err: err})
	return response
}

func (this GenericHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
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
