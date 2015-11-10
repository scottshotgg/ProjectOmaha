package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"omaha/system"
)

type SystemStatusResponse struct {
	Status *system.SystemStatus `json:"status"`
}

func SystemStatusHandler(w http.ResponseWriter, r *http.Request) {
	response := &SystemStatusResponse{Status: system.GetSystemStatus()}
	responseData, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		w.Write(getGenericErrorResponse(err.Error()))
	} else {
		w.Write(responseData)
	}
}
