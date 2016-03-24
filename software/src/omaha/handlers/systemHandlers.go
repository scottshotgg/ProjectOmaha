package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"omaha/database"
)

type systemStatusResponse struct {
	Zones 	[]*database.Zone	`json:"zones"`
	//Status	[]int				`json:"status"`
}


// remember about me if you are trying to change the speakerIDs
// change the zones and the speakerID to make them visible
func SystemStatusHandler(w http.ResponseWriter, r *http.Request) {
	//StatusArray[] = database.GetAllStatus()
	Zones := database.GetAllZones()			// maybe we should just use this to transfer all of the data

	response := &systemStatusResponse{Zones: Zones}
	responseData, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		w.Write(getGenericErrorResponse(err.Error()))
	} else {
		w.Write(responseData)
	}
}
