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
		return
	}
	status := system.GetSystemStatus()
	if volumeRequest.Volume >= 0 && volumeRequest.Volume <= 100 {
		fmt.Println("Telling the controller to turn to whatever I want", volumeRequest.Volume) // Print volume level
		status.SetVolume(volumeRequest.Volume)                                                 // Volume variable here)
		fmt.Fprint(w, "1")
	} else {
		fmt.Fprint(w, "0")
	}
}
