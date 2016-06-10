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
}

func SystemStatusHandler(w http.ResponseWriter, r *http.Request) {
	Zones 		:= database.GetAllZones()
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
