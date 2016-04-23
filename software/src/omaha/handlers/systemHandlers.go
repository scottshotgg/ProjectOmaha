package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"omaha/database"
)

type systemStatusResponse struct {
	Zones 			[]*database.Zone	`json:"zones"`
	PagingZones []*database.Zone	`json:"pagingZones"`
	//Status	[]int				`json:"status"`

	// this is where we could have some IP analysis to see whether or not we should preload certain data or not
}


// remember about me if you are trying to change the speakerIDs
// change the zones and the speakerID to make them visible
func SystemStatusHandler(w http.ResponseWriter, r *http.Request) {
	//StatusArray[] = database.GetAllStatus()
	Zones 		:= database.GetAllZones()			// maybe we should just use this to transfer all of the data\
	PagingZones := database.GetAllPagingZones()

	response := &systemStatusResponse{Zones: Zones, PagingZones: PagingZones}
	responseData, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		w.Write(getGenericErrorResponse(err.Error()))
	} else {
		w.Write(responseData)
	}
}
