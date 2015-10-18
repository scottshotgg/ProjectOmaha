package handlers

import (
	"encoding/json"
	"fmt"
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
		fmt.Fprint(w, "Err")
	} else {
		w.Write(responseData)
	}
}
