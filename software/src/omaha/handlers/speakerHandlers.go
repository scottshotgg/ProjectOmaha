package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"omaha/system"
	"omaha/database"
)

type speakerPutRequest struct {
	UpdatedAttributes []string          `json:"updatedAttributes"`
	AttributeValues   speakerAttributes `json:"attributeValues"`
	Speaker           int8              `json:"speaker"`
}

type speakerAttributes struct {
	Volume    int8 `json:"volume"`
	Averaging int8 `json:"averaging"`
	LED       bool `json:"led"` // Experimenting, not sure if this is needed or not
}

var speakerUpdateHandlers = map[string]func(*speakerAttributes, *database.ControllerStatus) error{
	"volume":    updateSpeakerVolume,
	"averaging": updateSpeakerAveragingMode,
	"led":       updateSpeakerLED,
}

func updateSpeakerVolume(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Volume >= 0 && attr.Volume <= 100 {
		log.Printf("Telling speaker %d to set volume to %d\n", speaker, attr.Volume)
		system.SetVolume(speaker, attr.Volume) // Volume variable here)
	} else {
		return errors.New("Invalid volume")
	}
	return nil
}

func updateSpeakerAveragingMode(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Averaging > 0 && attr.Averaging <= 20 {
		log.Printf("Telling speaker %d to set averaging mode to %d\n", speaker, attr.Averaging)
		system.SetAveragingMode(speaker, attr.Averaging) // Volume variable here)
	} else {
		return errors.New("Invalid averaging mode")
	}
	return nil
}

func updateSpeakerLED(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.LED {
		log.Printf("Telling speaker %d to turn on the LED\n", speaker)
		system.TurnLEDOn(speaker) // Volume variable here)
	} else {
		log.Printf("Telling speaker %d to turn off the LED\n", speaker)
		system.TurnLEDOff(speaker)
	}
	return nil
}

/*
	SpeakerPutHandler updates speaker attribtues according to the request
*/
func SpeakerPutHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	speakerRequest := &speakerPutRequest{}
	err := json.NewDecoder(r.Body).Decode(speakerRequest)
	if err != nil {
		if status.IsDebug() {
			log.Printf("SpeakerPutHandler json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	
	controller := database.GetSpeaker(speakerRequest.Speaker)
	if controller == nil {
		w.Write(getGenericErrorResponse("Invalid speaker ID from volume"))
		return
	}
	
	for _, attr := range speakerRequest.UpdatedAttributes {
		err = speakerUpdateHandlers[attr](&speakerRequest.AttributeValues, controller)
		if err != nil {
			w.Write(getGenericErrorResponse(err.Error()))
			return
		}
	}

	w.Write(getGenericSuccessResponse())
}
