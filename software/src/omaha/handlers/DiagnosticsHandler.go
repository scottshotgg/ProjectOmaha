package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"omaha/system"
)

type diagnosticsRequest struct {
	FilterType int8 `json:"filterType"`
}

func DiagnosticsHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	diagnosticsRequest := &diagnosticsRequest{}
	err := json.NewDecoder(r.Body).Decode(diagnosticsRequest)
	if err != nil {
		if status.IsDebug() {
			log.Println(err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	//controller := status.GetController(0x6B)
	//err = controller.SetAveragingMode(averagingRequest.FilterType)
	//	controller.Set
	w.Write(getGenericSuccessResponse())
}
