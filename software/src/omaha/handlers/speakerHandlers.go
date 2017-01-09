/*
	The handlers package includes the systems backend handlers that are used to serve requests that are submitted by the client. It includes many structs and functions to supports these features as well.
 */
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
	//"time"
	"sync"
)

/*
	All of these structs are used in AJAX handling.
*/
type speakerPutRequest struct {
	UpdatedAttributes []string          `json:"updatedAttributes"`
	AttributeValues   speakerAttributes `json:"attributeValues"`
	Speaker           int8              `json:"speaker"`
}

type speakerResponse struct {
	Volume									int8			`json:"volume"`
	Music 									int8			`json:"music"`
	Paging 									int8			`json:"paging"`
	Masking 								int8			`json:"masking"`

	LowerMusicThreshold			int8			`json:"lowerMusicThreshold"`
	UpperMusicThreshold			int8			`json:"upperMusicThreshold"`
	LowerPagingThreshold		int8			`json:"lowerPagingThreshold"`
	UpperPagingThreshold		int8			`json:"upperPagingThreshold"`
	LowerMaskingThreshold		int8			`json:"lowerMaskingThreshold"`
	UpperMaskingThreshold		int8			`json:"upperMaskingThreshold"`

	Effectiveness 					int8			`json:"effectiveness"`
	Pleasantness 						int8			`json:"pleasantness"`
	FadeTime								int8			`json:"fadetime"`
	FadeLevel								int8			`json:"fadelevel"`

	Target[][]							float64		`json:["target"]`
	TargetNames[]						string		`json:["targetNames"]`
	CurrentTarget[21]				float64		`json:["currentTarget"]` 

	Equalizer[][]						float64		`json:["equalizer"]`
	PresetNames[]						string		`json:["presetNames"]`
	CurrentPreset[21]				float64		`json:["currentPreset"]`

	MusicEqualizer[][]			float64		`json:["musicEqualizer"]`
	MusicPresetNames[]			string		`json:["musicPresetNames"]`
	CurrentMusicPreset[21]	float64		`json:["currentMusicPreset"]` 

	PagingEqualizer[][]			float64		`json:["pagingEqualizer"]`
	PagingPresetNames[]			string		`json:["pagingPresetNames"]`
	CurrentPagingPreset[21]	float64		`json:["currentPagingPreset"]` 

	Err     								string 		`json:"err"`
	Speaker									int8			`json:"speaker"`
	Name										string		`json:"name"`
	EqualizerMode						int8			`json:"equalizerMode"`
	SchedulingMode					int8			`json:"schedulingMode"`

	Schedule[][]						int				`json:"schedule"'`
}

type speakerGetRequest struct {
	Speaker	int8		`json:"speaker"`
}

type keepAlive struct {
	ID			int8		`json:"id"`
	Status	int8		`json:"status"`
}

type addPresetData struct {
	Speaker		int8		`json:"speaker"`
	Name			string	`json:"name"`
	Type 			int			`json:"type"`
	Constants	string	`json:"constants"`
	Target		string	`json:"target"`
	Update 		bool		`json:"update"`
}


type addPresetZoneData struct {
	Zone			int8		`json:"zone"`
	Name			string	`json:"name"`
	Type 			int			`json:"type"`
	Constants	string	`json:"constants"`
	Target		string	`json:"target"`
	Update 		bool		`json:"update"`
}

type changeEQModeData struct {
	Speaker int8	`json:"speaker"`
	Mode		int8	`json:"mode"`
}

type pagingRequest struct {
	Speaker	int8	`json:"speaker"`
}

type zoneData struct {
	ZoneName		string	`json:"name"`
	Speakers[] 	int8		`json:["speakers"]`
}

type addController struct {
	IP		string		`json:"ip"`
	Name 	string		`json:"name"`
}

type controllerResponse struct {
	Controllerids[]		int			`json:["controllerids"]`
	Ips[]							string	`json:["ips"]`
	Names[]						string	`json:["names"]`
}

type speakerAttributes struct {
	Volume    			string 		`json:"volume"`
	Threshold				string		`json:"threshold"`
	//Music					int8 			`json:"musicVolume"`
	Pleasantness 		int8 			`json:"effectiveness"`
	Effectiveness 	int8 			`json:"pleasantness"`
	LED       			bool 			`json:"led"` 
	Equalizer 			string		`json:"equalizer"`
	MusicEqualizer 	string		`json:"musicEqualizer"`
	PagingEqualizer string		`json:"pagingEqualizer"`
	ZoneID 					int8 			`json:"zoneId"`
	Paging					string		`json:"paging"`
	Target					string		`json:"target"`
}	

type timeSchedule struct {
	Speaker			int8	`json:"speaker"`
	Mode				int		`json:"mode"`
	Interval		int		`json:"interval"`
	Start 			int		`json:"start"`
	End 				int		`json:"end"`
	Day					int		`json:"day"`
	Times[24] 	int		`json:["times"]` 
}

type thresholdRequest struct {
	Speaker									int8	`json:"speaker"`
	LowerMusicThreshold			int8	`json:"musicMin"`
	UpperMusicThreshold			int8	`json:"musicMax"`
	LowerPagingThreshold		int8	`json:"pagingMin"`
	UpperPagingThreshold		int8	`json:"pagingMax"`
	LowerMaskingThreshold		int8	`json:"maskingMin"`
	UpperMaskingThreshold		int8	`json:"maskingMax"`
}

type volumesZoneData struct {
	Zone		int8	`json:"zone"`
	Volume	int8	`json:"volume"`
	Music		int8	`json:"music"`
	Paging	int8	`json:"paging"`
	Masking	int8	`json:"masking"`
}

type thresholdZoneRequest struct {
	Zone										int8	`json:"zone"`
	LowerMusicThreshold			int8	`json:"musicMin"`
	UpperMusicThreshold			int8	`json:"musicMax"`
	LowerPagingThreshold		int8	`json:"pagingMin"`
	UpperPagingThreshold		int8	`json:"pagingMax"`
	LowerMaskingThreshold		int8	`json:"maskingMin"`
	UpperMaskingThreshold		int8	`json:"maskingMax"`
}

type averagingZoneData struct {
	Zone					int8	`json:"zone"`
	Pleasantness	int8	`json:"effectiveness"`
	Effectiveness	int8 	`json:"pleasantness"`
}

type targetZoneData struct {
	Zone			int8		`json:"zone"`
	Constants string 	`json:"constants"`
	Mode			int8		`json:"mode"`
}

type updatePagingValuesData struct {
	Zone					int8	`json:"zone"`
	FadeTime			int8	`json:"fadeTime"`
	FadeLevel			int8	`json:"fadeLevel"`
	PagingVolume	int8	`json:"pagingVolume"`
}	

type timeScheduleZone struct {
	Zone				int8	`json:"zone"`
	Mode				int		`json:"mode"`
	Interval		int		`json:"interval"`
	Start 			int		`json:"start"`
	End 				int		`json:"end"`
	Day					int		`json:"day"`
	Times[24] 	int		`json:["times"]` 
}

type changeEQModeZoneData struct {
	Zone		int8		`json:"zone"`
	Mode		int8	`json:"mode"`
}

// This is the map that the SpeakerPutHandler uses to figure out which attributes need to be updated and what to use to update them
var speakerUpdateHandlers = map[string]func(*speakerAttributes, *database.ControllerStatus) error {
	"volume":			updateSpeakerVolume,
	//"music":		updateSpeakerMusic,
	"averaging":	updateSpeakerAveragingMode,
	"equalizer": 	updateSpeakerEqualizer,
	//"fadelevel":	updateSpeakerFadeLevel,
	//"fadetime":	updateSpeakerFadeTime,
	"paging":			updateSpeakerPaging,
	"zoneId": 		updateSpeakerZoneID,
	"target":			updateSpeakerTarget,
}


var dbmutex = &sync.Mutex{}

/*
	UpdateVolumesZone: This is the zone version of the UpdateVolumes function. As such, this function updates the volumes for the specified zones.
*/
func UpdateVolumesZone(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	//addPresetRequest := &addPresetData{}		// this is not a preset, but it is the same data
	volumesZoneRequest := &volumesZoneData{}
	err := json.NewDecoder(r.Body).Decode(volumesZoneRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("volumesZoneRequest json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	zone := database.GetZone(volumesZoneRequest.Zone)

	dbmutex.Lock()

	// for all the speakers in the zone	
	database.SaveVolumeZone(zone)
	for x := range zone.Speakers {
		if volumesZoneRequest.Volume >= 0 && volumesZoneRequest.Volume <= 100 && volumesZoneRequest.Music >= 0 && volumesZoneRequest.Music <= 100 && volumesZoneRequest.Paging >= 0 && volumesZoneRequest.Paging <= 100 && volumesZoneRequest.Masking >= 0 && volumesZoneRequest.Masking <= 100 {
			//log.Printf("Telling speaker %d to set volume to %d\n", speaker.ID, attr.Volume)
			
			log.Println(zone.Speakers[x].VolumeLevel)

			// the VolumeLevel[0-3] is because the volume levels that were received from the server were put into an array
			//if(volumesZoneRequest.Volume != zone.Speakers[x].VolumeLevel[0]) {
				zone.Speakers[x].VolumeLevel[0] = volumesZoneRequest.Volume
				zone.VolumeLevel[0] = volumesZoneRequest.Volume
				system.SetVolume(zone.Speakers[x])
				// will have to do a SetVolumeZone maybe 	
			//}
			//if(volumesZoneRequest.Music != zone.Speakers[x].VolumeLevel[1]) {
				zone.Speakers[x].VolumeLevel[1] = volumesZoneRequest.Music
				zone.VolumeLevel[1] = volumesZoneRequest.Music
				system.SetMusicVolume(zone.Speakers[x])
			//}
			//if(volumesZoneRequest.Paging != zone.Speakers[x].VolumeLevel[2]) {
				zone.Speakers[x].VolumeLevel[2] = volumesZoneRequest.Paging
				zone.VolumeLevel[2] = volumesZoneRequest.Paging
				system.SetPaging(zone.Speakers[x], volumesZoneRequest.Paging)		// 0 means volume adjustment
			//}
			//if(volumesZoneRequest.Masking != zone.Speakers[x].VolumeLevel[3]) {
				zone.Speakers[x].VolumeLevel[3] = volumesZoneRequest.Masking
				zone.VolumeLevel[3] = volumesZoneRequest.Masking
				system.SetSoundMaskingVolume(zone.Speakers[x])		// 0 means volume adjustment
			//}
			//if(l != 0) {
			database.SaveVolume(zone.Speakers[x])		// this may pose a security hole issue with injection, try incrementer mode or function by function if fails
			//defer database.SaveVolumeZone(zone)		// this may pose a security hole issue with injection, try incrementer mode or function by function if fails
		}
		w.Write(getGenericSuccessResponse())
	}
    	dbmutex.Unlock()
}

/*
	ThresholdHandlerZone: This is the zone version of the ThresholdHandler function. As such, this sets the volume thresholds for an entire zone.
*/
func ThresholdHandlerZone(w http.ResponseWriter, r *http.Request) {		// could merge this with AddPresetHandler
	status := system.GetSystemStatus()
	log.Println(status)
	thresholdZoneData := &thresholdZoneRequest{}		// this is not a preset, but it is the same data
	err := json.NewDecoder(r.Body).Decode(thresholdZoneData)

	if err != nil {
		if status.IsDebug() {
			log.Printf("AddTargetHandler json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	//log.Println("I got you packet dude, sendin one back")
	log.Println(thresholdZoneData)

	session, _ := r.Cookie("session")
	log.Println(session.Value)
	log.Println(thresholdZoneData.Zone)

	
	dbmutex.Lock()

	// if you have permission to perform this action. This is the backend part of the security. The frontend is also "secured" but people could easily manipulate the data on that side.
	if(int(thresholdZoneData.Zone) == database.AuthenticateZoneFromHash(session.Value) || database.AuthenticatePermissionFromHash(session.Value) > 0) {
		zone := database.GetZone(thresholdZoneData.Zone)
		log.Println("authenticated")
		for x := range zone.Speakers {
			database.UpdateThreshold(zone.Speakers[x].ID, thresholdZoneData.LowerMusicThreshold, thresholdZoneData.UpperMusicThreshold, thresholdZoneData.LowerPagingThreshold,	thresholdZoneData.UpperPagingThreshold,	thresholdZoneData.LowerMaskingThreshold, thresholdZoneData.UpperMaskingThreshold)
		} 
		database.UpdateThresholdZone(zone.ZoneID, thresholdZoneData.LowerMusicThreshold, thresholdZoneData.UpperMusicThreshold, thresholdZoneData.LowerPagingThreshold,	thresholdZoneData.UpperPagingThreshold,	thresholdZoneData.LowerMaskingThreshold, thresholdZoneData.UpperMaskingThreshold)
	}

	dbmutex.Unlock()
	// else don't write anything to let anyone that is trying to intrude know whether they are failing or not

	w.Write(getGenericSuccessResponse()) // this needs to be adapted to take into account the error from the database
}

/*
	UpdateAveragingZone: This is the zone version of the UpdateAveraging function. As such, it updates the averaging, which consists of the effectiveness and pleseantness of the speaker.
*/
func UpdateAveragingZone(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	//addPresetRequest := &addPresetData{}		// this is not a preset, but it is the same data
	averagingZoneRequest := &averagingZoneData{}
	err := json.NewDecoder(r.Body).Decode(averagingZoneRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("averagingZoneRequest json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}


	dbmutex.Lock()

	if averagingZoneRequest.Effectiveness > 0 && averagingZoneRequest.Effectiveness < 11 && averagingZoneRequest.Pleasantness > 0 && averagingZoneRequest.Pleasantness < 11 {	
		zone := database.GetZone(averagingZoneRequest.Zone)
		for x := range zone.Speakers {
			log.Printf("Telling speaker %d to set effectiveness to %d and pleasantness to %d\n", zone.Speakers[x].ID, averagingZoneRequest.Effectiveness, averagingZoneRequest.Pleasantness)
			system.SetAveragingMode(zone.Speakers[x], averagingZoneRequest.Effectiveness + averagingZoneRequest.Pleasantness)
			zone.Speakers[x].Effectiveness = averagingZoneRequest.Effectiveness 
			zone.Speakers[x].Pleasantness = averagingZoneRequest.Pleasantness
			zone.Effectiveness = averagingZoneRequest.Effectiveness
			zone.Pleasantness = averagingZoneRequest.Pleasantness
			database.SaveAveraging(zone.Speakers[x])
		}
		database.SaveAveragingZone(zone)
	}

	dbmutex.Unlock()

	w.Write(getGenericSuccessResponse())
}

/*
	UpdateSpeakerTargetZone: This is the zone version of the UpdateSpeakerTarget function. As such, this function updates the speaker target for all speaker within the specified zone.
*/
func UpdateSpeakerTargetZone(w http.ResponseWriter, r *http.Request) {		// could merge this with AddPresetHandler
	status := system.GetSystemStatus()
	targetZoneRequest := &targetZoneData{}
	err := json.NewDecoder(r.Body).Decode(targetZoneRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("updateSpeakerTargetZone json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	zone := database.GetZone(targetZoneRequest.Zone)

	constants := strings.Fields(targetZoneRequest.Constants)		// do not publish this function without checking for type/value errors
	log.Println(constants)

	if(len(constants) < 21) {
		// make a better error
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	var k = 0

	// while we are still in range of the constants
	// 
	dbmutex.Lock()

	for _, i := range constants {
		// try to parse a float value
		floatParse, err := strconv.ParseFloat(i, 64)
		if err != nil {
			panic(err)		// test if this returns
		}
		
		// for in range of the speakers that are in the zone
		for j := range zone.Speakers {
			// if the value is not the same that we already have in the zone
			if(floatParse != zone.Speakers[j].CurrentTarget[k]) {
				// seperate the decimal into the characteristic and the mantissa, before and after the decimal
				// floor to get the decimal
				whole := math.Floor(floatParse)
				// subtract the floatparse from the whole and then multiply by 100 to get the value
				decimal := math.Abs(whole - floatParse) * 100

				//log.Println(speaker.Equalizer[k], speaker.VolumeLevel)
				//system.SetEqualizerConstant(speaker, int8(whole), int8(k))
				log.Println("\n\n", whole, "\n\n")
				//system.SetEqualizerConstant(speaker, int8(decimal), int8(k))
				log.Println("\n\n", decimal, "\n\n")
				zone.Speakers[j].CurrentTarget[k] = floatParse	// this needs checking

				database.SaveBand(zone.Speakers[j], k, floatParse, 3)
				log.Println("band", k, "value", floatParse)
			}
			log.Printf("Telling speaker %d to change target to %s", zone.Speakers[j].ID, constants)
		}
		database.SaveBandZone(zone, k, floatParse, 3)
		k++

	}

	dbmutex.Unlock()

	w.Write(getGenericSuccessResponse()) // this needs to be adapted to take into account the error from the database
}

/*
	UpdateSpeakerEqualizerZone: This is the zone version of the UpdateSpeakerEqualizer function. As such, it updates the speaker equalizer of all speakers within the specified zone.
*/
func UpdateSpeakerEqualizerZone(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	
	// parse the constants out the string
	equalizerZoneRequest := &targetZoneData{}
	err := json.NewDecoder(r.Body).Decode(equalizerZoneRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("updateSpeakerEqualizerZone json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	// split the string
	constants := strings.Fields(equalizerZoneRequest.Constants)		// do not publish this function without checking for type/value errors

	if(len(constants) < 21) {
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	var k = 0
	// the equalizer has modes that it is sent that tell it what kind of equalizer it is
	// case 0 is a masking equalizer
	switch(equalizerZoneRequest.Mode) {
		case 0: {
			zone := database.GetZone(equalizerZoneRequest.Zone)

			dbmutex.Lock()
			for  i := 0; i < len(constants); i++ {
				for j := range zone.Speakers {
					speaker := database.GetSpeaker(zone.Speakers[j].ID)
					floatParse, err := strconv.ParseFloat(constants[i], 64)
					if err != nil {
						panic(err)
					}
					if(floatParse != speaker.CurrentPreset[i]) {
						whole := math.Floor(floatParse)

						decimal := math.Abs(whole - floatParse) * 100

						//log.Println(speaker.Equalizer[k], speaker.VolumeLevel)
						system.SetEqualizerConstant(speaker, int8(whole), int8(i), false)
						log.Println("Whole value:", int8(whole))
						system.SetEqualizerConstant(speaker, int8(decimal), 21, true)
						log.Println("Decimal value:", int8(decimal))
						speaker.CurrentPreset[i] = floatParse
						database.SaveBand(speaker, i, floatParse, 0)
						database.SaveBandZone(zone, i, floatParse, 0)
					}
				}
			}

		dbmutex.Unlock()
			k++

			break;
		}

		// case 1 is a music equalizer
		case 1: {
			zone := database.GetZone(equalizerZoneRequest.Zone)

			dbmutex.Lock()
			for  i := 0; i < len(constants); i++ {
				for j := range zone.Speakers {
					speaker := database.GetSpeaker(zone.Speakers[j].ID)
					floatParse, err := strconv.ParseFloat(constants[i], 64)
					if err != nil {
						panic(err)
					}
					if(floatParse != speaker.CurrentMusicPreset[i]) {
						whole := math.Floor(floatParse)

						decimal := math.Abs(whole - floatParse) * 100

						log.Println("Whole value:", int8(whole))
						log.Println("Decimal value:", int8(decimal))
						speaker.CurrentMusicPreset[i] = floatParse
						database.SaveBand(speaker, i, floatParse, 1)
						database.SaveBandZone(zone, i, floatParse, 1)
					}
				}
				k++
			}

			dbmutex.Unlock()

			break;
		}

		// case 2 is a paging equalizer
		case 2: {
			zone := database.GetZone(equalizerZoneRequest.Zone)

			dbmutex.Lock()

			for  i := 0; i < len(constants); i++ {
				for j := range zone.Speakers {
					speaker := database.GetSpeaker(zone.Speakers[j].ID)
					floatParse, err := strconv.ParseFloat(constants[i], 64)
					if err != nil {
						panic(err)
					}
					if(floatParse != speaker.CurrentMusicPreset[i]) {
						whole := math.Floor(floatParse)

						decimal := math.Abs(whole - floatParse) * 100

						log.Println("Whole value:", int8(whole))
						log.Println("Decimal value:", int8(decimal))
						speaker.CurrentPagingPreset[i] = floatParse
						database.SaveBand(speaker, i, floatParse, 2)
						database.SaveBandZone(zone, i, floatParse, 2)
					}
					log.Printf("Telling speaker %d to change equalizer to %s", speaker, constants)
				}
				k++
			}

			dbmutex.Unlock()

			break;
		}

		default: 
			log.Println("updateSpeakerEqualizer is getting the wrong value....")
	}

	w.Write(getGenericSuccessResponse())
}

/*
	AddPresetHandlerZone: This is the zone version of the AddPresteHandler. As such, it updates the speaker target for all targets within the specified zone.
*/
func AddPresetHandlerZone(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	addPresetRequestZone := &addPresetZoneData{}
	err := json.NewDecoder(r.Body).Decode(addPresetRequestZone)
	//controller := database.GetSpeaker(speakerRequest.Speaker)

	if err != nil {
		if status.IsDebug() {
			log.Printf("AddPresetHandler json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	log.Println(addPresetRequestZone)

	// these are the same cases as on the equalizer page
	switch addPresetRequestZone.Type {
		case 0:
			zone := database.GetZone(addPresetRequestZone.Zone)

			dbmutex.Lock()
			for x := range zone.Speakers {
				database.SavePreset(zone.Speakers[x].ID, addPresetRequestZone.Name, strings.Fields(addPresetRequestZone.Constants))
			}
			database.SavePresetZone(zone.ZoneID, addPresetRequestZone.Name, strings.Fields(addPresetRequestZone.Constants))
			dbmutex.Unlock()

		case 1:
			zone := database.GetZone(addPresetRequestZone.Zone)
			
			dbmutex.Lock()
			for x := range zone.Speakers {
				database.SaveMusicPreset(zone.Speakers[x].ID, addPresetRequestZone.Name, strings.Fields(addPresetRequestZone.Constants))
			}
			database.SaveMusicPresetZone(zone.ZoneID, addPresetRequestZone.Name, strings.Fields(addPresetRequestZone.Constants))
			dbmutex.Unlock()

		case 2:		
			zone := database.GetZone(addPresetRequestZone.Zone)
			dbmutex.Lock()
			for x := range zone.Speakers {
				database.SavePagingPreset(zone.Speakers[x].ID, addPresetRequestZone.Name, strings.Fields(addPresetRequestZone.Constants))
			}
			database.SavePagingPresetZone(zone.ZoneID, addPresetRequestZone.Name, strings.Fields(addPresetRequestZone.Constants))
			dbmutex.Unlock()
		default: 
			log.Println("AddPresetHandler MESSED UP SOMEHOW")
	}

	w.Write(getGenericSuccessResponse()) // this needs to be adapted to take into account the error from the database
}

/*
	UpdatePagingValuesZone: This is the zone version of the UpdatePagingValues table. As such, it updates the paging values for the entire specified entire zone.
*/
func UpdatePagingValuesZone(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	updatePagingValuesRequest := &updatePagingValuesData{}
	err := json.NewDecoder(r.Body).Decode(updatePagingValuesRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("updateSpeakerEqualizerZone json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	if updatePagingValuesRequest.FadeTime > -1 && updatePagingValuesRequest.FadeTime < 101 && updatePagingValuesRequest.FadeLevel > -1 && updatePagingValuesRequest.FadeLevel < 101 && updatePagingValuesRequest.PagingVolume > -1 && updatePagingValuesRequest.PagingVolume < 101 {	
		zone := database.GetZone(updatePagingValuesRequest.Zone)
		

		dbmutex.Lock()
		for x := range zone.Speakers {
			if(updatePagingValuesRequest.FadeTime != zone.Speakers[x].PagingLevel[0]) {
					zone.Speakers[x].PagingLevel[0] = updatePagingValuesRequest.FadeTime
					zone.PagingLevel[0] = updatePagingValuesRequest.FadeTime
					system.SetPaging(zone.Speakers[x], updatePagingValuesRequest.FadeTime + 101) 	// offset tells the mcu to set fade time	
					database.SaveFade(zone.Speakers[x])		// we can optimize this by reducing the generality of the functio}
			}
			if(updatePagingValuesRequest.FadeLevel != zone.Speakers[x].PagingLevel[1] && updatePagingValuesRequest.FadeLevel != 0) {
					zone.Speakers[x].PagingLevel[1] = updatePagingValuesRequest.FadeLevel
					system.SetPaging(zone.Speakers[x], updatePagingValuesRequest.FadeLevel)
				  database.SaveFade(zone.Speakers[x])
			}
			if(updatePagingValuesRequest.PagingVolume != zone.Speakers[x].VolumeLevel[2]) {
				zone.Speakers[x].VolumeLevel[2] = updatePagingValuesRequest.PagingVolume
				system.SetPaging(zone.Speakers[x], updatePagingValuesRequest.PagingVolume)
				database.SaveVolume(zone.Speakers[x])
			}
		}
		zone.PagingLevel[0] = updatePagingValuesRequest.FadeTime
		zone.PagingLevel[1] = updatePagingValuesRequest.FadeLevel
		zone.VolumeLevel[2] = updatePagingValuesRequest.PagingVolume
		database.SaveFadeZone(zone)
		database.SaveVolumeZone(zone)
		dbmutex.Unlock()

	} else {
		return
	}

	return
}

/*
	AddTargetHandlerZone: This is the zone version of the AddTargetHandler function. As such, this function adds a target in the database for an entire zone.
*/
func AddTargetHandlerZone(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	addPresetRequestZone := &addPresetZoneData{}		// this is not a preset, but it is the same data so why reinvent the wheel
	err := json.NewDecoder(r.Body).Decode(addPresetRequestZone)

	if err != nil {
		if status.IsDebug() {
			log.Printf("AddTargetHandler json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	zone := database.GetZone(addPresetRequestZone.Zone)

	dbmutex.Lock()
	for x := range zone.Speakers {
		database.SaveTarget(zone.Speakers[x].ID, addPresetRequestZone.Name, strings.Fields(addPresetRequestZone.Constants))
	}
	database.SaveTargetZone(addPresetRequestZone.Zone, addPresetRequestZone.Name, strings.Fields(addPresetRequestZone.Constants))
	dbmutex.Unlock()

	w.Write(getGenericSuccessResponse())
}

/*
	ScheduleTimeZone: This is the zone version of the function ScheduleTime. As such, it saves the schedule for the specified zone.
*/
func ScheduleTimeZone(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	timeGivenZone := &timeScheduleZone{}
	err := json.NewDecoder(r.Body).Decode(timeGivenZone)

	if err != nil {
		if status.IsDebug() {
			log.Printf("ScheduleTimeZone json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	log.Println(timeGivenZone)
	zone := database.GetZone(timeGivenZone.Zone)

	dbmutex.Lock()
	for x := range zone.Speakers {
		database.UpdateSchedule(zone.Speakers[x].ID, timeGivenZone.Day, timeGivenZone.Interval, timeGivenZone.Start, timeGivenZone.End, timeGivenZone.Times)
		err = database.ChangeSchedulingMode(zone.Speakers[x].ID, timeGivenZone.Mode)
  }
	database.UpdateScheduleZone(zone.ZoneID, timeGivenZone.Day, timeGivenZone.Interval, timeGivenZone.Start, timeGivenZone.End, timeGivenZone.Times)
 	err = database.ChangeSchedulingModeZone(zone.ZoneID, timeGivenZone.Mode)
	dbmutex.Unlock()

	log.Println("THIS IS THE TIME GIVEN", timeGivenZone)

  w.Write(getGenericSuccessResponse())
}

/*
	ChangeEQModeZone: This is the zone version of the function ChangeEQMode. As such, this function changes the EQ mode for each speaker within the specified zone.
*/
func ChangeEQModeZone(w http.ResponseWriter, r *http.Request) {		// could merge this with AddPresetHandler
	status := system.GetSystemStatus()

	changeEQModeZoneRequest := &changeEQModeZoneData{}
	err := json.NewDecoder(r.Body).Decode(changeEQModeZoneRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("changeEQModeZone json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	zone := database.GetZone(changeEQModeZoneRequest.Zone)

	dbmutex.Lock()
	for j := range zone.Speakers {
		err = database.ChangeEQMode(zone.Speakers[j].ID, changeEQModeZoneRequest.Mode)
		_, err = system.SetEQMode(zone.Speakers[j].ID, changeEQModeZoneRequest.Mode)
	}
	err = database.ChangeEQModeZone(zone.ZoneID, changeEQModeZoneRequest.Mode)
	dbmutex.Unlock()
	// this should tell the controller to send out the packets for the music constants
	// should also call to change the status in the database to music mode

	log.Println("I got you packet dude, sendin one back", changeEQModeZoneRequest)
	w.Write(getGenericSuccessResponse())
}

/*
	updateSpeakerVolume is a private function that is used by the SpeakerUpdateHandlers map when it comes across volume.
*/
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
			panic(err)
		}

		volumeArray[k] = int8(intParse)
		k++
	}

	if volumeArray[0] >= 0 && volumeArray[0] <= 100 && volumeArray[1] >= 0 && volumeArray[1] <= 100 && volumeArray[2] >= 0 && volumeArray[2] <= 100 && volumeArray[3] >= 0 && volumeArray[3] <= 100 {
		//log.Printf("Telling speaker %d to set volume to %d\n", speaker.ID, attr.Volume)
		
		log.Println(speaker.VolumeLevel)

		//if(volumeArray[0] != speaker.VolumeLevel[0]) {
			speaker.VolumeLevel[0] = volumeArray[0]
			system.SetVolume(speaker) 	
		//}
		//if(volumeArray[1] != speaker.VolumeLevel[1]) {
			speaker.VolumeLevel[1] = volumeArray[1]
			system.SetMusicVolume(speaker)
		//}
		//if(volumeArray[2] != speaker.VolumeLevel[2]) {
			speaker.VolumeLevel[2] = volumeArray[2]
			system.SetPaging(speaker, volumeArray[2])
		//}
		//if(volumeArray[3] != speaker.VolumeLevel[3]) {
			speaker.VolumeLevel[3] = volumeArray[3]
			system.SetSoundMaskingVolume(speaker)
		//}
		
		dbmutex.Lock()
		database.SaveVolume(speaker)		// this may pose a security hole issue with injection, try incrementer mode or function by function if fails
		dbmutex.Unlock()

	} else {
		return errors.New("Invalid volume")
	}

	return nil
}

/*
	updateSpeakerAveragingMode is a private function that is used by the SpeakerUpdateHandlers map when it comes across averaging.
*/
func updateSpeakerAveragingMode(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	if attr.Effectiveness > 0 && attr.Effectiveness < 11 && attr.Pleasantness > 0 && attr.Pleasantness < 11 {
		log.Printf("Telling speaker %d to set effectiveness to %d and pleasantness to %d\n", speaker.ID, attr.Effectiveness, attr.Pleasantness)
		system.SetAveragingMode(speaker, attr.Effectiveness + attr.Pleasantness)
		speaker.Effectiveness = attr.Effectiveness 
		speaker.Pleasantness = attr.Pleasantness
		dbmutex.Lock()
		database.SaveAveraging(speaker)
		dbmutex.Unlock()
	} else {
		return errors.New("Invalid averaging mode")
	}
	return nil
}

/*
	updateSpeakerEqualizer is a private function that is used by the SpeakerUpdateHandler map when it comes across equalizer.
*/
func updateSpeakerEqualizer(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	log.Println(attr.Equalizer)
	constants := strings.Fields(attr.Equalizer)
	log.Println(constants, len(constants))

	for i := range speaker.Equalizer {
		log.Println(speaker.Equalizer[i])
	}

	if(len(constants) < 21) {
		return errors.New("Invalid amount of constants")
	}

	var k = 0

	switch(speaker.EqualizerMode) {
		case 0: {
			dbmutex.Lock()
			for  i := 0; i < len(constants); i++ {
				floatParse, err := strconv.ParseFloat(constants[i], 64)
				if err != nil {
					panic(err)
				}
				if(floatParse != speaker.CurrentPreset[i]) {
					whole := math.Floor(floatParse)

					decimal := math.Abs(whole - floatParse) * 100

					system.SetEqualizerConstant(speaker, int8(whole), int8(i), false)
					log.Println("Whole value:", int8(whole))
					system.SetEqualizerConstant(speaker, int8(decimal), 21, true)
					log.Println("Decimal value:", int8(decimal))
					speaker.CurrentPreset[i] = floatParse
					database.SaveBand(speaker, i, floatParse, 0)
				}

				floatParse, err = strconv.ParseFloat(constants[i], 64)
				if err != nil {
					panic(err)
				}
				k++
			}
			dbmutex.Lock()

			break;
		}

		case 1: {
			dbmutex.Lock()
			for  i := 0; i < len(constants); i++ {
				floatParse, err := strconv.ParseFloat(constants[i], 64)
				if err != nil {
					panic(err)
				}
				if(floatParse != speaker.CurrentMusicPreset[i]) {
					whole := math.Floor(floatParse)

					decimal := math.Abs(whole - floatParse) * 100

					log.Println("Whole value:", int8(whole))
					log.Println("Decimal value:", int8(decimal))
					speaker.CurrentMusicPreset[i] = floatParse
					database.SaveBand(speaker, i, floatParse, 1)
				}

				floatParse, err = strconv.ParseFloat(constants[i], 64)
				if err != nil {
					panic(err)
				}
				k++
			}
			dbmutex.Unlock()

			break;
		}

		case 2: {
			dbmutex.Lock()
			for  i := 0; i < len(constants); i++ {
				floatParse, err := strconv.ParseFloat(constants[i], 64)
				if err != nil {
					panic(err)
				}
				if(floatParse != speaker.CurrentMusicPreset[i]) {
					whole := math.Floor(floatParse)
					decimal := math.Abs(whole - floatParse) * 100
					log.Println("Whole value:", int8(whole))
					log.Println("Decimal value:", int8(decimal))
					speaker.CurrentPagingPreset[i] = floatParse
					database.SaveBand(speaker, i, floatParse, 2)
				}

				floatParse, err = strconv.ParseFloat(constants[i], 64)
				if err != nil {
					panic(err)
				}
				k++
			}
			dbmutex.Unlock()

			break;
		}

		default: 
			log.Println("updateSpeakerEqualizer is getting the wrong value....")
	}

	log.Printf("Telling speaker %d to change equalizer to %s", speaker.ID, constants)

	return nil
}

/*
	updateSpeakerTarget is a function that is used by the SpeakerUpdateHandlers map when it comes across target.
*/
func updateSpeakerTarget(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	constants := strings.Fields(attr.Target)

	for i := range speaker.Target {
		log.Println(speaker.Target[i])
	}

	if(len(constants) < 21) {
		return errors.New("Invalid amount of constants")
	}

	var k = 0
	dbmutex.Lock()
	for _, i := range constants {
		floatParse, err := strconv.ParseFloat(i, 64)
		if err != nil {
			panic(err)		
		}

		if(floatParse != speaker.CurrentTarget[k]) {
			whole := math.Floor(floatParse)
			decimal := math.Abs(whole - floatParse) * 100
			log.Println("\n\n", whole, "\n\n")
			log.Println("\n\n", decimal, "\n\n")
			speaker.CurrentTarget[k] = floatParse
			database.SaveBand(speaker, k, floatParse, 3)
		}
			k++
	}
	dbmutex.Unlock()

	log.Printf("Telling speaker %d to change target to %s", speaker.ID, constants)

	return nil
}

/*
	updateSpeakerPaging is a private function that is used when the SpeakerUpdateHandlers map comes across paging.
*/
func updateSpeakerPaging(attr *speakerAttributes, speaker *database.ControllerStatus) error {

	pagings := strings.Fields(attr.Paging)
	var pagingArray [3] int8

	if(len(pagings) < 3) {
		return errors.New("Invalid amount of pagings")
	}

	var k = 0
	for _, i := range pagings {
		intParse, err := strconv.Atoi(i)
		if err != nil {
			panic(err)		
		}

		pagingArray[k] = int8(intParse)
		k++
	}

	if pagingArray[0] >= 0 && pagingArray[0] <= 100 && pagingArray[1] >= 0 && pagingArray[1] <= 100 && pagingArray[2] >= 0 && pagingArray[2] <= 100 {

		if(pagingArray[0] != speaker.PagingLevel[0]) {
			speaker.PagingLevel[0] = pagingArray[0]
			system.SetPaging(speaker, pagingArray[0] + 101) 	// offset tells the mcu to set fade time	
			dbmutex.Lock()
			database.SaveFade(speaker)
			dbmutex.Unlock()
		}
		if(pagingArray[1] != speaker.PagingLevel[1] && pagingArray[1] != 0) {
			speaker.PagingLevel[1] = pagingArray[1]
			system.SetPaging(speaker, pagingArray[1])
			dbmutex.Lock()
			database.SaveFade(speaker)
			dbmutex.Unlock()
		}
		if(pagingArray[2] != speaker.VolumeLevel[2]) {
			speaker.VolumeLevel[2] = pagingArray[2]
			system.SetPaging(speaker, pagingArray[2])
			dbmutex.Lock()
			database.SaveVolume(speaker)
			dbmutex.Unlock()
		}

		log.Println(pagingArray)

	} else {
		return errors.New("Invalid volume")
	}

	return nil
}

/*
	updateSpeakerZoneID is a private function that is used when reassigning a speaker to another zone.
*/
func updateSpeakerZoneID(attr *speakerAttributes, speaker *database.ControllerStatus) error {
	dbmutex.Lock()
	database.SetSpeakerToZoneByID(speaker, attr.ZoneID)
	dbmutex.Unlock()
	return nil
}

/*
	SpeakerPutHandler is used when the server recieves a PUT request and then updates speaker attribtues according to the request and how the SpeakerUpdateHandlers map is defined.
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

	session, _ := r.Cookie("session")
	log.Println(session.Value)
	if(int(speakerRequest.Speaker) == database.AuthenticateSpeakerFromHash(session.Value) || database.AuthenticatePermissionFromHash(session.Value) > 0) {
		log.Println("authenticated bro")
	} else {
		log.Println("invalid request")
		w.Write(getGenericErrorResponse("invalid speaker request: permission"))
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

/*
	SpeakerGetHandler is used when the server recieves a GET request, which is mostly for the processing the speaker requests. This is done so by filling the response JSON struct.
*/
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

	response, _ := json.Marshal(fillSpeakerResponse(controller))

	w.Write(response)
}

/*
	fillSpeakerResponse is a private function used for filling the response struct for the SpeakerGetHandler
*/
func fillSpeakerResponse(controller *database.ControllerStatus) speakerResponse {

	speakerResponse := speakerResponse {
		Name: 			controller.Name, 
		Volume: 		controller.VolumeLevel[0], 
		Music: 			controller.VolumeLevel[1], 
		Paging: 		controller.VolumeLevel[2], 
		Masking: 		controller.VolumeLevel[3], 


		LowerMusicThreshold:		controller.LowerThreshold[0],
		UpperMusicThreshold:		controller.UpperThreshold[0],
		LowerPagingThreshold:		controller.LowerThreshold[1],
		UpperPagingThreshold:		controller.UpperThreshold[1],
		LowerMaskingThreshold:	controller.LowerThreshold[2],
		UpperMaskingThreshold:	controller.UpperThreshold[2],


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
		CurrentMusicPreset: [21] float64 {
			controller.CurrentMusicPreset[0], controller.CurrentMusicPreset[1], controller.CurrentMusicPreset[2], controller.CurrentMusicPreset[3], 
			controller.CurrentMusicPreset[4], controller.CurrentMusicPreset[5], controller.CurrentMusicPreset[6], controller.CurrentMusicPreset[7], 
			controller.CurrentMusicPreset[8], controller.CurrentMusicPreset[9], controller.CurrentMusicPreset[10], controller.CurrentMusicPreset[11], 
			controller.CurrentMusicPreset[12], controller.CurrentMusicPreset[13], controller.CurrentMusicPreset[14], controller.CurrentMusicPreset[15], 
			controller.CurrentMusicPreset[16], controller.CurrentMusicPreset[17], controller.CurrentMusicPreset[18], controller.CurrentMusicPreset[19], 
			controller.CurrentMusicPreset[20] }, 
		CurrentPagingPreset: [21] float64 {
			controller.CurrentPagingPreset[0], controller.CurrentPagingPreset[1], controller.CurrentPagingPreset[2], controller.CurrentPagingPreset[3], 
			controller.CurrentPagingPreset[4], controller.CurrentPagingPreset[5], controller.CurrentPagingPreset[6], controller.CurrentPagingPreset[7], 
			controller.CurrentPagingPreset[8], controller.CurrentPagingPreset[9], controller.CurrentPagingPreset[10], controller.CurrentPagingPreset[11], 
			controller.CurrentPagingPreset[12], controller.CurrentPagingPreset[13], controller.CurrentPagingPreset[14], controller.CurrentPagingPreset[15], 
			controller.CurrentPagingPreset[16], controller.CurrentPagingPreset[17], controller.CurrentPagingPreset[18], controller.CurrentPagingPreset[19], 
			controller.CurrentPagingPreset[20] }, 
		CurrentTarget: [21] float64 {
			controller.CurrentTarget[0], controller.CurrentTarget[1], controller.CurrentTarget[2], controller.CurrentTarget[3],
			controller.CurrentTarget[4], controller.CurrentTarget[5], controller.CurrentTarget[6], controller.CurrentTarget[7], 
			controller.CurrentTarget[8], controller.CurrentTarget[9], controller.CurrentTarget[10], controller.CurrentTarget[11], 
			controller.CurrentTarget[12], controller.CurrentTarget[13], controller.CurrentTarget[14], controller.CurrentTarget[15], 
			controller.CurrentTarget[16], controller.CurrentTarget[17], controller.CurrentTarget[18], controller.CurrentTarget[19], 
			controller.CurrentTarget[20] }, 
		Target: 			controller.Target, 
		TargetNames: 		controller.TargetNames, 
		Equalizer: 			controller.Equalizer,
		PresetNames: 		controller.PresetNames, 
		MusicEqualizer:		controller.MusicEqualizer,
		MusicPresetNames:	controller.MusicPresetNames, 
		EqualizerMode:		controller.EqualizerMode, 
		SchedulingMode:		controller.SchedulingMode,
		Schedule:					controller.Schedule }
	return speakerResponse
}

/*
	AddPresetHandler is used to add a preset to the database for the specified speaker.
*/
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

	switch addPresetRequest.Type {
		case 0:
			dbmutex.Lock()
			database.SavePreset(addPresetRequest.Speaker, addPresetRequest.Name, strings.Fields(addPresetRequest.Constants))
			dbmutex.Unlock()
		case 1:
			dbmutex.Lock()
			database.SaveMusicPreset(addPresetRequest.Speaker, addPresetRequest.Name, strings.Fields(addPresetRequest.Constants))
			dbmutex.Unlock()
		case 2:
			dbmutex.Lock()
			database.SavePagingPreset(addPresetRequest.Speaker, addPresetRequest.Name, strings.Fields(addPresetRequest.Constants))
			dbmutex.Unlock()
		default: 
			log.Println("AddPresetHandler MESSED UP SOMEHOW")
	}

	w.Write(getGenericSuccessResponse()) // this needs to be adapted to take into account the error from the database
}

/*
	AddTargetHandler adds a target for the specified speaker.
*/
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
	dbmutex.Lock()
	database.SaveTarget(addPresetRequest.Speaker, addPresetRequest.Name, strings.Fields(addPresetRequest.Constants))
	dbmutex.Unlock()
	w.Write(getGenericSuccessResponse()) // this needs to be adapted to take into account the error from the database
}

/*
	PagingRequestHandler is not used currently, but is reserved for connecting the paging system and making calls.
*/
func PagingRequestHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	// It is NOT a GET request, but the struct that is used is identical to the one that would have been used for this, so why reinvent the wheel
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

/*
	ControllerUpdateHandler is used for handling the KeepAlive requests that the webpage sends out. 
*/
func ControllerUpdateHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	log.Println(status)

	log.Println(status.ID, status.BrokenLink)

	keepAliveResponse := keepAlive {
		ID: status.ID, 
		Status: status.BrokenLink }

	log.Println("Controller Request", keepAliveResponse)

	response, _ := json.Marshal(keepAliveResponse)

	w.Write(response)
}

/*
	CreateZoneHandler is used for handling zone creation.
*/
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

	dbmutex.Lock()
	err = database.CreateZone(createZone.Speakers, createZone.ZoneName)
	dbmutex.Unlock()
	
	if err != nil {
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	w.Write(getGenericSuccessResponse())
}

/*
	CreatePagingZoneHandler is the paging version of the CreateZoneHandler and is used for paging zone creation. 
*/
func CreatePagingZoneHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	createZone := &zoneData{}		//  this is not the same thing as a zone but it uses the same data type
	err := json.NewDecoder(r.Body).Decode(createZone)

	log.Println("Creating a paging zone", createZone)

	if err != nil {
		if status.IsDebug() {
			log.Printf("createPagingZone json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	dbmutex.Lock()
	err = database.CreatePagingZone(createZone.Speakers, createZone.ZoneName)
	dbmutex.Unlock()

	if err != nil {
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	w.Write(getGenericSuccessResponse())
}

/*
	ChangeEQMode is used for changing the EQ mode of the specified speaker
*/
func ChangeEQMode(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	changeEQModeRequest := &changeEQModeData{}
	err := json.NewDecoder(r.Body).Decode(changeEQModeRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("ChangeEQMode json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	dbmutex.Lock()
	err = database.ChangeEQMode(changeEQModeRequest.Speaker, changeEQModeRequest.Mode)
	_, err = system.SetEQMode(changeEQModeRequest.Speaker, changeEQModeRequest.Mode)
	dbmutex.Unlock()

	log.Println("I got you packet dude, sendin one back", changeEQModeRequest)
	w.Write(getGenericSuccessResponse())
}

/*
	ScheduleTime is used for updating the schedule for the specified speaker.
*/
func ScheduleTime(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()

	timeGiven := &timeSchedule{}
	err := json.NewDecoder(r.Body).Decode(timeGiven)

	if err != nil {
		if status.IsDebug() {
			log.Printf("ScheduleTime json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	log.Println(timeGiven)

	dbmutex.Lock()
	database.UpdateSchedule(timeGiven.Speaker, timeGiven.Day, timeGiven.Interval, timeGiven.Start, timeGiven.End, timeGiven.Times)
	err = database.ChangeSchedulingMode(timeGiven.Speaker, timeGiven.Mode)
	dbmutex.Unlock()

	log.Println("THIS IS THE TIME GIVEN", timeGiven)

  w.Write(getGenericSuccessResponse())
}

/*
	ThresholdHandler is used for updating the volume thresholds for the specified speaker.
*/
func ThresholdHandler(w http.ResponseWriter, r *http.Request) {
	status := system.GetSystemStatus()
	log.Println(status)
	thresholdData := &thresholdRequest{}		// this is not a preset, but it is the same data
	err := json.NewDecoder(r.Body).Decode(thresholdData)

	if err != nil {
		if status.IsDebug() {
			log.Printf("AddTargetHandler json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	log.Println(thresholdData)

	session, _ := r.Cookie("session")
	log.Println(session.Value)
	log.Println(thresholdData.Speaker)


	dbmutex.Lock()
	if(int(thresholdData.Speaker) == database.AuthenticateSpeakerFromHash(session.Value) || database.AuthenticatePermissionFromHash(session.Value) > 0) {
		database.UpdateThreshold(thresholdData.Speaker, thresholdData.LowerMusicThreshold, thresholdData.UpperMusicThreshold, thresholdData.LowerPagingThreshold,	thresholdData.UpperPagingThreshold,	thresholdData.LowerMaskingThreshold, thresholdData.UpperMaskingThreshold)
	}
	dbmutex.Unlock()

	w.Write(getGenericSuccessResponse())
}

/*
	AddController is used to handle the add controller call on the login screen that is needed to link the controllers.
*/
func AddController(w http.ResponseWriter, r *http.Request) {		// could merge this with AddPresetHandler
	status := system.GetSystemStatus()
	controllerInformation := &addController{}
	err := json.NewDecoder(r.Body).Decode(controllerInformation)
	log.Println(controllerInformation);

	if err != nil {
		if status.IsDebug() {
			log.Printf("AddTargetHandler json decoding error: %s\n", err)
		}
		log.Println(err)
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}

	database.InsertController(controllerInformation.IP, controllerInformation.Name)
	log.Println("inserterino")
	w.Write(getGenericSuccessResponse()) // this needs to be adapted to take into account the error from the database
}

/*
	LoadControllers is used when the login screen calls to see which controllers have been added and adds them to the drop down so that the user can select the desired controller.
*/
func LoadControllers(w http.ResponseWriter, r *http.Request) {		// could merge this with AddPresetHandler
	//status := system.GetSystemStatus()
	log.Println("we are here")
	controllerIDs, ips, names := database.RetrieveControllers()

	LoadControllersResponse := &controllerResponse{
		Controllerids: controllerIDs,
		Ips: ips,
		Names: names}

	response, _ := json.Marshal(LoadControllersResponse)

	w.Write(response)
}


// use this when creating a new handler. It will give you a skeleton of one so that you can craft more
// Example handler
/*
func ChangeEQMode(w http.ResponseWriter, r *http.Request) {		// could merge this with AddPresetHandler
	status := system.GetSystemStatus()

	//addPresetRequest := &addPresetData{}		// this is not a preset, but it is the same data
	//err := json.NewDecoder(r.Body).Decode(addPresetRequest)

	if err != nil {
		if status.IsDebug() {
			log.Printf("AddTargetHandler json decoding error: %s\n", err)
		}
		w.Write(getGenericErrorResponse(err.Error()))
		return
	}
	log.Println("I got you packet dude, sendin one back")
	//log.Println(addPresetRequest)
	//database.SaveTarget(addPresetRequest.Speaker, addPresetRequest.Name, strings.Fields(addPresetRequest.Constants))
	w.Write(getGenericSuccessResponse()) // this needs to be adapted to take into account the error from the database
}
*/