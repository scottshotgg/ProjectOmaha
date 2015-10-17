package handlers

import (
	"fmt"
	"net/http"
	"omaha/system"
)

func DemoStartHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LED on handler called")
	status := system.GetSystemStatus()
	if !status.IsLEDOn() {
		fmt.Println("Telling the controller to turn on")
		status.TurnLEDOn()
		if status.IsDebug() {
			fmt.Println("Turned the LED on")
		}
	}

	fmt.Fprint(w, "1")
}

func DemoStopHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LED off handler called")
	status := system.GetSystemStatus()
	if status.IsLEDOn() {
		fmt.Println("Telling the controller to turn off")
		status.TurnLEDOff()
		if status.IsDebug() {
			fmt.Println("Turned the LED off")
		}
	}
	fmt.Fprint(w, "1")
}

func DemoLEDHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("LED status handler called")
	status := system.GetSystemStatus()
	if status.IsLEDOn() {
		fmt.Fprint(w, "1")
	} else {
		fmt.Fprint(w, "0")
	}
}

/*func DemoVolumeUpHandler(w http.ResponseWriter, r *http.Request) {
	// status := system.GetSystemStatus()
	// if status.GetVolumeLevel() < 100 {
	// 	fmt.Println("Telling the controller to turn up on a Tuesday") // Print volume level
	// 	status.VolumeUp()
	// 	if status.IsDebug() {
	// 		fmt.Println("Turned the volume up") // Print out volume level too
	// 	}
	// }
	fmt.Fprint(w, "1")
}

func DemoVolumeDownHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("VOLUME DOWN")
	status := system.GetSystemStatus()
	if status.GetVolumeLevel() > 0 {
		fmt.Println("Telling the controller to turn down for what") // Print volume level
		status.VolumeDown()
		if status.IsDebug() {
			fmt.Println("Turned the volume down") // Print out volume level too
		}
	}
	fmt.Fprint(w, "1")
}*/

func DemoVolumeVariableHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var volumeLevel = r.PostForm["test"][0]	// This is a string for some reason
	fmt.Println("VOLUME VARIABLE")			// Print volume variable too
	status := system.GetSystemStatus()
		fmt.Println("Telling the controller to turn to whatever I want: " + volumeLevel)	// Print volume level
		status.VolumeVariable(1)	// need int value not string, hardcoded as 1 for now
		if status.IsDebug() {
			fmt.Println("Turned the volume to " + volumeLevel)	// Print out volume level too
		}
	fmt.Fprint(w, "1")
}
