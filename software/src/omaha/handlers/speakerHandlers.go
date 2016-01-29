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
	Volume    	string 	`json:"volume"`
	//Music		int8 	`json:"musicVolume"`
	Averaging 	int8 	`json:"averaging"`
	LED       	bool 	`json:"led"` 
	Equalizer 	string	`json:"equalizer"`
	ZoneID 		int8 	`json:"zoneId"`
	Paging		string	`json:"paging"`
}

var speakerUpdateHandlers = map[string]func(*speakerAttributes, *database.ControllerStatus) error {
	"volume":		updateSpeakerVolume,
	//"music":		updateSpeakerMusic,
	"averaging":	updateSpeakerAveragingMode,
	"led":       	updateSpeakerLED,
	"equalizer": 	updateSpeakerEqualizer,
	//"fadelevel":	updateSpeakerFadeLevel,
	//"fadetime":		updateSpeakerFadeTime,
	"paging":		updateSpeakerPaging,
	"zoneId": 		updateSpeakerZoneID,
}

func updateSpeakerVolume(attr *speakerAttributes, speaker *database.ControllerStatus) error {

	volumes := strings.Fields(attr.Volume)
	log.Println(attr.Volume)
	log.Println(volumes)

	var volumeArray [4] int8

	if(len(volumes) < 4) {
		return errors.New("Invalid amount of volumes")
	}

	var k = 0
	for _, i := range volumes {
		intParse, err := strconv.Atoi(i)
		if err != nil {
			panic(err)		// test if this returns
		}

		volumeArray[k] = int8(intParse)
		// log.Printf("You changed band %d to level %d", k, intParse)
		// log.Println(speaker.Equalizer[k])		// see if this works, if it does then we know that it can be accessed as an array
		//}
		k++
	}

	if volumeArray[0] >= 0 && volumeArray[0] <= 100 && volumeArray[1] >= 0 && volumeArray[1] <= 100 && volumeArray[2] >= 0 && volumeArray[2] <= 100 && volumeArray[3] >= 0 && volumeArray[3] <= 100 {
		//log.Printf("Telling speaker %d to set volume to %d\n", speaker.ID, attr.Volume)
		
		log.Println(speaker.VolumeLevel)

		if(volumeArray[0] != speaker.VolumeLevel[0]) {
			speaker.VolumeLevel[0] = volumeArray[0]
			system.SetVolume(speaker) 	
		}
		if(volumeArray[1] != speaker.VolumeLevel[1]) {
			speaker.VolumeLevel[1] = volumeArray[1]
			system.SetMusicVolume(speaker)
		}
		if(volumeArray[2] != speaker.VolumeLevel[2]) {
			speaker.VolumeLevel[2] = volumeArray[2]
			system.SetPaging(speaker, volumeArray[2])		// 0 means volume adjustment
		}
		if(volumeArray[3] != speaker.VolumeLevel[3]) {
			speaker.VolumeLevel[3] = volumeArray[3]
			system.SetSoundMaskingVolume(speaker)		// 0 means volume adjustment
		}
		//if(l != 0) {
		defer database.SaveVolume(speaker)		// this may pose a security hole issue with injection, try incrementer mode or function by function if fails
		//}
		//log.Println(speaker.VolumeLevel)

	} else {
		return errors.New("Invalid volume")
	}

	return nil
}
/*
func updateSpeakerMusic(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Music >= 0 && attr.Music <= 100 {
		log.Printf("Telling speaker %d to set music to %d\n", speaker.ID, attr.Music)
		//system.SetMusicVolume(speaker, attr.Music) 
	} else {
		return errors.New("Invalid music volume")
	}
	return nil
}
*/

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
	constants := strings.Fields(attr.Equalizer)		// do not publish this function without checking for type/value errors
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
			speaker.Equalizer[k] = intParse	// this needs checking
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

/*
func updateSpeakerFadeLevel(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Paging >= 0 && attr.Paging <= 100 {
		log.Printf("Telling speaker %d to set paging to %d\n", speaker.ID, attr.Paging)
	//	system.SetPaging(speaker, attr.Paging) 
		//speaker.PagingLevel = attr.Paging
		//database.SaveVolume(speaker)
	} else {
		return errors.New("Invalid paging")
	}
	return nil
}

func updateSpeakerFadeTime(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Paging >= 0 && attr.Paging <= 100 {
		log.Printf("Telling speaker %d to set paging to %d\n", speaker.ID, attr.Paging)
	//	system.SetPaging(speaker, attr.Paging) 
		//speaker.PagingLevel = attr.Paging
		//database.SaveVolume(speaker)
	} else {
		return errors.New("Invalid paging")
	}
	return nil
}
*/

func updateSpeakerPaging(attr *speakerAttributes, speaker *database.ControllerStatus) error {

	pagings := strings.Fields(attr.Paging)
	//log.Println(attr.Volume)
	//log.Println(volumes)

	var pagingArray [3] int8

	if(len(pagings) < 3) {
		return errors.New("Invalid amount of pagings")
	}

	var k = 0
	for _, i := range pagings {
		intParse, err := strconv.Atoi(i)
		if err != nil {
			panic(err)		// test if this returns
		}

		pagingArray[k] = int8(intParse)
		// log.Printf("You changed band %d to level %d", k, intParse)
		// log.Println(speaker.Equalizer[k])		// see if this works, if it does then we know that it can be accessed as an array
		//}
		k++
	}

	if pagingArray[0] >= 0 && pagingArray[0] <= 100 && pagingArray[1] >= 0 && pagingArray[1] <= 100 && pagingArray[2] >= 0 && pagingArray[2] <= 100 {
		//log.Printf("Telling speaker %d to set volume to %d\n", speaker.ID, attr.Volume)
		
		//log.Println(speaker.VolumeLevel)

		if(pagingArray[0] != speaker.PagingLevel[0]) {
			speaker.PagingLevel[0] = pagingArray[0]
			system.SetPaging(speaker, pagingArray[0] + 101) 	
			database.SaveFade(speaker)		// we can optimize this by reducing the generality of the function
		}
		if(pagingArray[1] != speaker.PagingLevel[1] && pagingArray[1] != 0) {
			speaker.PagingLevel[1] = pagingArray[1]
			system.SetPaging(speaker, pagingArray[1])
			database.SaveFade(speaker)
		}
		if(pagingArray[2] != speaker.VolumeLevel[2]) {
			speaker.VolumeLevel[2] = pagingArray[2]
			system.SetPaging(speaker, pagingArray[2])		// 0 means volume adjustment
			database.SaveVolume(speaker)
		}
		//if(l != 0) {
		//database.SaveVolume(speaker)		// this may pose a security hole issue with injection, try incrementer mode or function by function if fails
		//}
		log.Println(pagingArray)

	} else {
		return errors.New("Invalid volume")
	}

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
