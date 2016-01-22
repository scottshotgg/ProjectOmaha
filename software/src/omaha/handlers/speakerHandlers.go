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
	Volume    	int8 	`json:"volume"`
	Music		int8 	`json:"musicVolume"`
	Averaging 	int8 	`json:"averaging"`
	LED       	bool 	`json:"led"` 
	Equalizer 	string	`json:"equalizer"`
	ZoneID 		int8 	`json:"zoneId"`
}

var speakerUpdateHandlers = map[string]func(*speakerAttributes, *database.ControllerStatus) error {
	"volume":		updateSpeakerVolume,
	"music":		updateSpeakerMusic,
	"averaging":	updateSpeakerAveragingMode,
	"led":       	updateSpeakerLED,
	"equalizer": 	updateSpeakerEqualizer,
	"zoneId": 		updateSpeakerZoneID,
}

func updateSpeakerVolume(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Volume >= 0 && attr.Volume <= 100 {
		log.Printf("Telling speaker %d to set volume to %d\n", speaker.ID, attr.Volume)
		system.SetVolume(speaker, attr.Volume) 
		speaker.VolumeLevel = attr.Volume
		database.SaveVolume(speaker)
	} else {
		return errors.New("Invalid volume")
	}
	return nil
}

func updateSpeakerMusic(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Music >= 0 && attr.Music <= 100 {
		log.Printf("Telling speaker %d to set music to %d\n", speaker.ID, attr.Music)
		//system.SetMusicVolume(speaker, attr.Music) 
	} else {
		return errors.New("Invalid music volume")
	}
	return nil
}

func updateSpeakerAveragingMode(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Averaging > 0 && attr.Averaging <= 20 {
		log.Printf("Telling speaker %d to set averaging mode to %d\n", speaker.ID, attr.Averaging)
		system.SetAveragingMode(speaker, attr.Averaging)
		speaker.AveragingMode = attr.Averaging
		database.SaveAveraging(speaker)
	} else {
		return errors.New("Invalid averaging mode")
	}
	return nil
}

func updateSpeakerEqualizer(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	constants := strings.Fields(attr.Equalizer)
	//log.Println(constants)

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

		if(intParse != speaker.Equalizer[k]) {			// change this to pull from the db, it might already do that
			//log.Println(speaker.Equalizer[k], speaker.VolumeLevel)
			system.SetEqualizerConstant(speaker, int8(intParse), int8(k))
			speaker.Equalizer[k] = intParse
			//log.Printf("You changed band %d to level %d", k, intParse)
		//	log.Println(speaker.Equalizer[k])		// see if this works, if it does then we know that it can be accessed as an array
			database.SaveBand(speaker, k, intParse)
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
		err = speakerUpdateHandlers[attr](&speakerRequest.AttributeValues, controller)		// change these names to be more goddamn consistent, this code is fucking hard to read

		if err != nil {
			w.Write(getGenericErrorResponse(err.Error()))
			return
		}
	}

	w.Write(getGenericSuccessResponse())
}
