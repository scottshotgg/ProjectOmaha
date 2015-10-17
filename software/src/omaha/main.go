package main

import (
	"flag"
	"net/http"
	"omaha/handlers"
	"omaha/system"
)

func main() {
	// initialization
	var debug = flag.Bool("d", false, "help me!!!")
	flag.Parse()
	system.InitializeSystemStatus(*debug)
	if !(*debug) {
		status := system.GetSystemStatus()
		defer status.Port.Close()
	}
	http.HandleFunc("/", handlers.LoginHandler)
	http.HandleFunc("/app/", handlers.AppHandler)
	http.HandleFunc("/demo/start/", handlers.DemoStartHandler)
	http.HandleFunc("/demo/stop/", handlers.DemoStopHandler)
	http.HandleFunc("/demo/status/", handlers.DemoLEDHandler)
	http.HandleFunc("/scott/test/", handlers.DemoVolumeUpHandler)
	http.HandleFunc("/scott/test/", handlers.DemoVolumeVariableHandler)

	// file handlers
	http.Handle("/css/", handlers.CssHandler)
	http.Handle("/bower_components/", handlers.BowerHandler)
	http.Handle("/components/", handlers.ComponentsHandler)

	http.ListenAndServe(":8080", nil)
}
