package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"omaha/database"
	"omaha/system"
)

type averageRequest struct {
	FilterType int8 `json:"filterType"`
}

func AveragingHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	averagingRequest := &averageRequest{}
	err := json.NewDecoder(r.Body).Decode(averagingRequest)
	if err != nil {
		if status.IsDebug() {
			log.Println(err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	controller := database.GetSpeaker(0x6b)
	err = system.SetAveragingMode(controller, averagingRequest.FilterType)
	w.Write(getGenericSuccessResponse())
}
