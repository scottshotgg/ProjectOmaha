package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"omaha/database"
	"omaha/system"
	"strings"
	"strconv"
)

type speakerPutRequest struct {
	UpdatedAttributes []string          `json:"updatedAttributes"`
	AttributeValues   speakerAttributes `json:"attributeValues"`
	Speaker           int8              `json:"speaker"`
}

type speakerAttributes struct {
	Volume    int8 `json:"volume"`
	Averaging int8 `json:"averaging"`
	LED       bool `json:"led"` 
	Equalizer string `json:"equalizer"`
	ZoneID int8 `json:"zoneId"`
}

var speakerUpdateHandlers = map[string]func(*speakerAttributes, *database.ControllerStatus) error {
	"volume":    updateSpeakerVolume,
	"averaging": updateSpeakerAveragingMode,
	"led":       updateSpeakerLED,
	"equalizer": updateSpeakerEqualizer,
	"zoneId": updateSpeakerZoneID,
}

func updateSpeakerVolume(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Volume >= 0 && attr.Volume <= 100 {
		log.Printf("Telling speaker %d to set volume to %d\n", speaker.ID, attr.Volume)
		system.SetVolume(speaker, attr.Volume) 
	} else {
		return errors.New("Invalid volume")
	}
	return nil
}

func updateSpeakerAveragingMode(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Averaging > 0 && attr.Averaging <= 20 {
		log.Printf("Telling speaker %d to set averaging mode to %d\n", speaker.ID, attr.Averaging)
		system.SetAveragingMode(speaker, attr.Averaging)
	} else {
		return errors.New("Invalid averaging mode")
	}
	return nil
}

func updateSpeakerEqualizer(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	constants := strings.Fields(attr.Equalizer)

	if(len(constants) < 21) {
		return errors.New("Invalid amount of constants")
	}

	var k = 0
	for _, i := range constants {
		intParse, err := strconv.Atoi(i)
		if err != nil {
			panic(err)		// test if this returns
		}
		//constantsInts = append(constantsInts, int8(intParse))

		if(int8(intParse) != speaker.Equalizer[k]) {
			log.Println(speaker.Equalizer[k], speaker.VolumeLevel)
			system.SetEqualizerConstant(speaker, int8(intParse), int8(k))
			speaker.Equalizer[k] = int8(intParse)
			log.Printf("You changed band %d to level %d", k, intParse)
			log.Println(speaker.Equalizer[k])
		}
		k++

		//log.Println("constantsInts: ", constantsInts)
	}

	log.Printf("Telling speaker %d to change equalizer to %s", speaker.ID, constants)

	return nil
}

func updateSpeakerLED(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.LED {	// i think this may be wrong but idc
		log.Printf("Telling speaker %d to turn on the LED\n", speaker.ID)
		system.TurnLEDOn(speaker) // Volume variable here)
	} else {
		log.Printf("Telling speaker %d to turn off the LED\n", speaker.ID)
		system.TurnLEDOff(speaker)
	}
	return nil
}

func updateSpeakerZoneID(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	database.SetSpeakerToZoneByID(speaker, attr.ZoneID)
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
	log.Println(controller)
	if controller == nil {
		w.Write(getGenericErrorResponse("Invalid speaker ID from request"))		// print out the request identifier here
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
