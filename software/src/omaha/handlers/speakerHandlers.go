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
	"math"
)

type speakerPutRequest struct {
	UpdatedAttributes []string          `json:"updatedAttributes"`
	AttributeValues   speakerAttributes `json:"attributeValues"`
	Speaker           int8              `json:"speaker"`
}

type speakerResponse struct {
	Volume				int8		`json:"volume"`
	Music 				int8		`json:"music"`
	Paging 				int8		`json:"paging"`
	Masking 			int8		`json:"masking"`
	Effectiveness 		int8		`json:"effectiveness"`
	Pleasantness 		int8		`json:"pleasantness"`
	FadeTime			int8		`json:"fadetime"`
	FadeLevel			int8		`json:"fadelevel"`
	Target[][]			float64			`json:["target"]`
	TargetNames[]		string		`json:["targetNames"]`
	Equalizer[][]		float64			`json:["equalizer"]`
	PresetNames[]		string		`json:["presetNames"]`
	CurrentPreset[21]	float64			`json:["currentPreset"]` 
	CurrentTarget[21]	float64			`json:["currentTarget"]` 
	Err     			string 		`json:"err"`
	Speaker				int8		`json:"speaker"`
	Name				string		`json:"name"`
}

type speakerGetRequest struct {
	Speaker		int8		`json:"speaker"`
}

type keepAlive struct {
	ID			int8		`json:"id"`
	Status		int8		`json:"status"`
}

type addPresetData struct {
	Speaker		int8		`json:"speaker"`
	Name		string		`json:"name"`
	Constants	string		`json:"constants"`
	Target		string		`json:"target"`
	Update 		bool		`json:"update"`

}

type pagingRequest struct {
	Speaker		int8	`json:"speaker"`
}

type zoneData struct {
	ZoneName	string	`json:"name"`
	Speakers[] 	int8	`json:["speakers"]`
}

type speakerAttributes struct {
	Volume    		string 		`json:"volume"`
	//Music		int8 		`json:"musicVolume"`
	Pleasantness 	int8 		`json:"effectiveness"`
	Effectiveness 	int8 		`json:"pleasantness"`
	LED       		bool 		`json:"led"` 
	Equalizer 		string		`json:"equalizer"`
	ZoneID 			int8 		`json:"zoneId"`
	Paging			string		`json:"paging"`
	Target			string		`json:"target"`
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
	"target":		updateSpeakerTarget,
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
	if attr.Effectiveness > 0 && attr.Effectiveness < 11 && attr.Pleasantness > 0 && attr.Pleasantness < 11 {
		log.Printf("Telling speaker %d to set effectiveness to %d and pleasantness to %d\n", speaker.ID, attr.Effectiveness, attr.Pleasantness)
		system.SetAveragingMode(speaker, attr.Effectiveness + attr.Pleasantness)
		speaker.Effectiveness = attr.Effectiveness 
		speaker.Pleasantness = attr.Pleasantness
		database.SaveAveraging(speaker)
	} else {
		return errors.New("Invalid averaging mode")
	}
	return nil
}

func updateSpeakerEqualizer(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	log.Println(attr, speaker)

	constants := strings.Fields(attr.Equalizer)		// do not publish this function without checking for type/value errors
	//log.Println(constants)

	for i := range speaker.Equalizer {
		log.Println(speaker.Equalizer[i])
	}

	if(len(constants) < 21) {
		return errors.New("Invalid amount of constants")
	}

	var k = 0
	for _, i := range constants {
		floatParse, err := strconv.ParseFloat(i, 64)
		if err != nil {
			panic(err)		// test if this returns
		}
		//constantsInts = append(constantsInts, int8(floatParse))

		if(floatParse != speaker.CurrentPreset[k]) {			// change this to pull from the db, it might already do that
			whole := math.Floor(floatParse)

			decimal := math.Abs(whole - floatParse) * 100

			//log.Println(speaker.Equalizer[k], speaker.VolumeLevel)
			system.SetEqualizerConstant(speaker, int8(whole), int8(k))
			log.Println("\n\n", whole, "\n\n")
			system.SetEqualizerConstant(speaker, int8(decimal), 21)
			log.Println("\n\n", decimal, "\n\n")
			speaker.CurrentPreset[k] = floatParse	// this needs checking
			//log.Printf("You changed band %d to level %d", k, floatParse)
		//	log.Println(speaker.Equalizer[k])		// see if this works, if it does then we know that it can be accessed as an array
			database.SaveBand(speaker, k, floatParse, false)
		}
			k++
		//log.Println("constantsInts: ", constantsInts)
	}

	log.Printf("Telling speaker %d to change equalizer to %s", speaker.ID, constants)

	return nil
}

func updateSpeakerTarget(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	constants := strings.Fields(attr.Target)		// do not publish this function without checking for type/value errors
	//log.Println(constants)

	for i := range speaker.Target {
		log.Println(speaker.Target[i])
	}

	if(len(constants) < 21) {
		return errors.New("Invalid amount of constants")
	}

	var k = 0
	for _, i := range constants {
		floatParse, err := strconv.ParseFloat(i, 64)
		if err != nil {
			panic(err)		// test if this returns
		}
		//constantsInts = append(constantsInts, int8(floatParse))

		if(floatParse != speaker.CurrentTarget[k]) {			// change this to pull from the db, it might already do that
			whole := math.Floor(floatParse)

			decimal := math.Abs(whole - floatParse) * 100

			//log.Println(speaker.Equalizer[k], speaker.VolumeLevel)
			//system.SetEqualizerConstant(speaker, int8(whole), int8(k))
			log.Println("\n\n", whole, "\n\n")
			//system.SetEqualizerConstant(speaker, int8(decimal), int8(k))
			log.Println("\n\n", decimal, "\n\n")
			speaker.CurrentPreset[k] = floatParse	// this needs checking
			//log.Printf("You changed band %d to level %d", k, floatParse)
		//	log.Println(speaker.Equalizer[k])		// see if this works, if it does then we know that it can be accessed as an array
			database.SaveBand(speaker, k, floatParse, true)
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
			system.SetPaging(speaker, pagingArray[0] + 101) 	// offset tells the mcu to set fade time	
			database.SaveFade(speaker)		// we can optimize this by reducing the generality of the function
		}
		if(pagingArray[1] != speaker.PagingLevel[1] && pagingArray[1] != 0) {
			speaker.PagingLevel[1] = pagingArray[1]
			system.SetPaging(speaker, pagingArray[1])
			database.SaveFade(speaker)
		}
		if(pagingArray[2] != speaker.VolumeLevel[2]) {
			speaker.VolumeLevel[2] = pagingArray[2]
			system.SetPaging(speaker, pagingArray[2])
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

	log.Println(r)
	log.Println(r.Body)

	log.Println(speakerRequest)
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

func SpeakerGetHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	speakerRequest := &speakerGetRequest{}
	err := json.NewDecoder(r.Body).Decode(speakerRequest)
	controller := database.GetSpeaker(speakerRequest.Speaker)
	log.Println(controller)

	log.Println(speakerRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("SpeakerGetHandler json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	//speakerRequest.Volume = 
	//log.Println(Volume)
	//response, _ := json.Marshal()
	//speakerResponse := fillSpeakerResponse(controller)
	//speakerResponse := speakerResponse{Volume: controller.VolumeLevel[0]}
	response, _ := json.Marshal(fillSpeakerResponse(controller))

	w.Write(response)
}

func fillSpeakerResponse(controller *database.ControllerStatus) speakerResponse {

	speakerResponse := speakerResponse {
		Name: 			controller.Name, 
		Volume: 		controller.VolumeLevel[0], 
		Music: 			controller.VolumeLevel[1], 
		Paging: 		controller.VolumeLevel[2], 
		Masking: 		controller.VolumeLevel[3], 
		Effectiveness: 	controller.Effectiveness,
		Pleasantness:	controller.Pleasantness,
		FadeTime: 		controller.PagingLevel[0], 
		FadeLevel: 		controller.PagingLevel[1], 
		CurrentPreset: [21] float64 {
			controller.CurrentPreset[0], controller.CurrentPreset[1], controller.CurrentPreset[2], controller.CurrentPreset[3], 
			controller.CurrentPreset[4], controller.CurrentPreset[5], controller.CurrentPreset[6], controller.CurrentPreset[7], 
			controller.CurrentPreset[8], controller.CurrentPreset[9], controller.CurrentPreset[10], controller.CurrentPreset[11], 
			controller.CurrentPreset[12], controller.CurrentPreset[13], controller.CurrentPreset[14], controller.CurrentPreset[15], 
			controller.CurrentPreset[16], controller.CurrentPreset[17], controller.CurrentPreset[18], controller.CurrentPreset[19], 
			controller.CurrentPreset[20] }, 
		CurrentTarget: [21] float64 {
			controller.CurrentTarget[0], controller.CurrentTarget[1], controller.CurrentTarget[2], controller.CurrentTarget[3],
			controller.CurrentTarget[4], controller.CurrentTarget[5], controller.CurrentTarget[6], controller.CurrentTarget[7], 
			controller.CurrentTarget[8], controller.CurrentTarget[9], controller.CurrentTarget[10], controller.CurrentTarget[11], 
			controller.CurrentTarget[12], controller.CurrentTarget[13], controller.CurrentTarget[14], controller.CurrentTarget[15], 
			controller.CurrentTarget[16], controller.CurrentTarget[17], controller.CurrentTarget[18], controller.CurrentTarget[19], 
			controller.CurrentTarget[20] }, 
		Target: 		controller.Target, 
		TargetNames: 	controller.TargetNames, 
		Equalizer: 		controller.Equalizer, 
		PresetNames: 	controller.PresetNames }
	//log.Println(controller)
	return speakerResponse
}


func AddPresetHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	addPresetRequest := &addPresetData{}
	err := json.NewDecoder(r.Body).Decode(addPresetRequest)
	//controller := database.GetSpeaker(speakerRequest.Speaker)

	if err != nil {
		if status.IsDebug() {
			log.Printf("AddPresetHandler json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	log.Println(addPresetRequest)

	database.SavePreset(addPresetRequest.Speaker, addPresetRequest.Name, strings.Fields(addPresetRequest.Constants))

	w.Write(getGenericSuccessResponse()) // this needs to be adapted to take into account the error form the database shit
}

func AddTargetHandler(w http.ResponseWriter, r *http.Request) {		// could merge this with AddPresetHandler
	status := system.GetSystemStatus()

	addPresetRequest := &addPresetData{}		// this is not a preset, but it is the same data
	err := json.NewDecoder(r.Body).Decode(addPresetRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("AddTargetHandler json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	log.Println(addPresetRequest)
	database.SaveTarget(addPresetRequest.Speaker, addPresetRequest.Name, strings.Fields(addPresetRequest.Constants))
	w.Write(getGenericSuccessResponse()) // this needs to be adapted to take into account the error form the database shit
}

func PagingRequestHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	// It is NOT a GET request, but the struct that is used is identical to the one that would have been used for 
	// this, so why reinvent the wheel
	pagingRequest := &speakerGetRequest{}
	err := json.NewDecoder(r.Body).Decode(pagingRequest)

	log.Println("Making a paging request", pagingRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("pagingRequest json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	err = system.SendPagingRequest(pagingRequest.Speaker)
	
	if err != nil {
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	w.Write(getGenericSuccessResponse())
}

func ControllerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	log.Println(status)

	// It is NOT a GET request, but the struct that is used is identical to the one that would have been used for 
	// this, so why reinvent the wheel
	//pagingRequest := &speakerGetRequest{}
	//err := json.NewDecoder(r.Body).Decode(pagingRequest)

	//log.Println("Making a paging request", pagingRequest)
	
	/*if err != nil {
		if status.IsDebug() {
			//log.Printf("pagingRequest json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	err = system.SendPagingRequest(pagingRequest.Speaker)
	
	if err != nil {
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}*/

	keepAliveResponse := keepAlive {
		ID: status.ID, 
		Status: status.BrokenLink }

	//response, _ := json.Marshal(fillSpeakerResponse(controller))

	log.Println("Controller Request", keepAliveResponse)

	response, _ := json.Marshal(keepAliveResponse)

	w.Write(response)

	//w.Write(getGenericSuccessResponse())
}

func CreateZoneHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	createZone := &zoneData{}
	err := json.NewDecoder(r.Body).Decode(createZone)

	log.Println("Creating a zone", createZone)

	if err != nil {
		if status.IsDebug() {
			log.Printf("createZone json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	err = database.CreateZone(createZone.Speakers, createZone.ZoneName)
	
	if err != nil {
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	w.Write(getGenericSuccessResponse())
}

func CreatePagingZoneHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	//  this is not the same thing as a zone but it uses the same data type
	createZone := &zoneData{}
	err := json.NewDecoder(r.Body).Decode(createZone)

	log.Println("Creating a paging zone", createZone)

	if err != nil {
		if status.IsDebug() {
			log.Printf("createPagingZone json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	err = database.CreatePagingZone(createZone.Speakers, createZone.ZoneName)
	
	if err != nil {
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	w.Write(getGenericSuccessResponse())
}

	// make a split string fucntion when you fee like it

	//constants := strings.Fields(addPresetRequest.Constants)		// do not publish this function without checking for type/value errors
	//log.Println(constants)
	/*var constantsInts [21] int8

	if(len(constants) < 21) {
		return //errors.New("Invalid amount of constants")
	}

	var k = 0
	for _, i := range constants {
		intParse, err := strconv.Atoi(i)
		if err != nil {
			panic(err)		// test if this returns
		}

			constantsInts[k] = int8(intParse)	// this needs checking
		}
			k++*/

		//log.Println("constantsInts: ", constantsInts)