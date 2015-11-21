package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"omaha/database"
)

type systemStatusResponse struct {
	Zones []*database.Zone `json:"zones"`
}

func SystemStatusHandler(w http.ResponseWriter, r *http.Request) {
	response := &systemStatusResponse{Zones: database.GetAllZones()}
	responseData, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		w.Write(getGenericErrorResponse(err.Error()))
	} else {
		w.Write(responseData)
	}
}
