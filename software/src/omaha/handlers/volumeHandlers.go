package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"omaha/system"
)

type VolumeQuery struct {
	Volume int8 `json:"volume"`
}

/*
	Sets the volume of the control. Expects a post request with variables:
	- volume (integer)
*/
func SetVolumeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SET VOLUME")
	volumeRequest := &VolumeQuery{}
	err := json.NewDecoder(r.Body).Decode(volumeRequest)
	if err != nil {
		fmt.Fprint(w, "0")
	}
	fmt.Println(volumeRequest.Volume)
	status := system.GetSystemStatus()
	//fmt.Println("Telling the controller to turn to whatever I want: ", volumeLevel)	// Print volume level
	//status.VolumeVariable(volumeLevel)	// need int value not string, hardcoded as 1 for now
	if status.GetVolumeLevel() > 0 { // Change to comapre incoming volume variable, also > 100
		fmt.Println("Telling the controller to turn to whatever I want", volumeRequest.Volume) // Print volume level
		status.SetVolume(volumeRequest.Volume)                                                 // Volume variable here)
		if status.IsDebug() {
			fmt.Println("Turned the volume to", volumeRequest.Volume) // Print out volume level too
		}
	}
	fmt.Fprint(w, "1")
}
