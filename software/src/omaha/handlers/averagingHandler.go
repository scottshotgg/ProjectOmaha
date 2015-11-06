package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
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
			fmt.Println(err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	controller := status.GetController(0x6b)
	err = system.SetAveragingMode(controller, averagingRequest.FilterType)
	w.Write(getGenericSuccessResponse())
}
