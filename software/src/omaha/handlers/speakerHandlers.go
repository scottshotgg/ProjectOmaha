package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"omaha/system"
)

type speakerPutRequest struct {
	UpdatedAttributes []string          `json:"updatedAttributes"`
	AttributeValues   speakerAttributes `json:"attributeValues"`
	Speaker           int8              `json:"speaker"`
}

type speakerAttributes struct {
	Volume int8 `json:"volume"`
}

var speakerUpdateHandlers = map[string]func(*speakerAttributes, int8) error{
	"volume": updateSpeakerVolume,
}

func updateSpeakerVolume(attr *speakerAttributes, speaker int8) error {
	status := system.GetSystemStatus()
	controller := status.GetController(speaker)
	if controller == nil {
		return errors.New("Invalid speaker ID")
	}
	if attr.Volume >= 0 && attr.Volume <= 100 {
		fmt.Printf("Telling speaker %d to set volume to %d\n", speaker, attr.Volume)
		controller.SetVolume(attr.Volume) // Volume variable here)
	} else {
		return errors.New("Invalid volume")
	}
	return nil
}

/*
	update speaker attribtues
*/
func SpeakerPutHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	speakerRequest := &speakerPutRequest{}
	err := json.NewDecoder(r.Body).Decode(speakerRequest)
	if err != nil {
		if status.IsDebug() {
			fmt.Println(err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	for _, attr := range speakerRequest.UpdatedAttributes {
		err = speakerUpdateHandlers[attr](&speakerRequest.AttributeValues, speakerRequest.Speaker)
		if err != nil {
			w.Write(getGenericErrorResponse(err.Error()))
			return
		}
	}

	w.Write(getGenericSuccessResponse())
}
