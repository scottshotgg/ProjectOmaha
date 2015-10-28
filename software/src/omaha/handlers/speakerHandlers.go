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
	Volume    int8 `json:"volume"`
	Averaging int8 `json:"averaging"`
	LED       bool `json:"led"` // Experimenting, not sure if this is needed or not
}

var speakerUpdateHandlers = map[string]func(*speakerAttributes, int8) error{
	"volume":    updateSpeakerVolume,
	"averaging": updateSpeakerAveragingMode,
	"led":       updateSpeakerLED,
}

func updateSpeakerVolume(attr *speakerAttributes, speaker int8) error {
	status := system.GetSystemStatus()
	controller := status.GetController(speaker)
	if controller == nil {
		return errors.New("Invalid speaker ID from volume")
	}
	if attr.Volume >= 0 && attr.Volume <= 100 {
		fmt.Printf("Telling speaker %d to set volume to %d\n", speaker, attr.Volume)
		controller.SetVolume(attr.Volume) // Volume variable here)
	} else {
		return errors.New("Invalid volume")
	}
	return nil
}

func updateSpeakerAveragingMode(attr *speakerAttributes, speaker int8) error {
	status := system.GetSystemStatus()
	fmt.Println("This is the speaker that I'm getting in averaging: ", speaker)
	controller := status.GetController(speaker)
	if controller == nil {
		return errors.New("Invalid speaker ID from averaging")
	}
	if attr.Averaging > 0 && attr.Averaging <= 20 {
		fmt.Printf("Telling speaker %d to set averaging mode to %d\n", speaker, attr.Averaging)
		controller.SetAveragingMode(attr.Averaging) // Volume variable here)
	} else {
		return errors.New("Invalid averaging mode")
	}
	return nil
}

func updateSpeakerLED(attr *speakerAttributes, speaker int8) error {
	status := system.GetSystemStatus()
	fmt.Println("This is the speaker that I'm getting in LED: ", speaker)
	controller := status.GetController(speaker)
	if controller == nil {
		return errors.New("Invalid speaker ID from LED")
	}
	if attr.LED {
		fmt.Printf("Telling speaker %d to turn on the LED\n", speaker)
		controller.TurnLEDOn() // Volume variable here)
	} else {
		fmt.Printf("Telling speaker %d to turn off the LED\n", speaker)
		controller.TurnLEDOff()
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
			fmt.Printf("SpeakerPutHandler json decoding error: %s\n", err)
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
