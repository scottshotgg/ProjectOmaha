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

var filterTypes struct {
	Test int8
}

func init() {
	filterTypes.Test = 21 // just an example
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
	controller := status.GetController(0x65)
	err = controller.SetAveragingMode(averagingRequest.FilterType)
	w.Write(getGenericSuccessResponse())
}
