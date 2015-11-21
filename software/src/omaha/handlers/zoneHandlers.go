package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"omaha/system"
	"omaha/database"
)

type zonePostRequest struct {
	Name string          `json:"name"`
}

// ZonePostHandler handles an http request to create a new zone
func ZonePostHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	zoneRequest := &zonePostRequest{}
	err := json.NewDecoder(r.Body).Decode(zoneRequest)
	if err != nil {
		if status.IsDebug() {
			log.Printf("ZonePostHandler json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	
	database.AddZone(zoneRequest.Name)
	w.Write(getGenericSuccessResponse())
}
