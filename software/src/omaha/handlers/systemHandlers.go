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


// remember about me if you are trying to change the speakerIDs
// change the zones and the speakerID to make them visible
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
