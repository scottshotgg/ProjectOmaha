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
