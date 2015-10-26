package main

import (
	"flag"
	"log"
	"net/http"
	"omaha/database"
	"omaha/handlers"
	"omaha/system"
)

func main() {
	// initialization
	var debug = flag.Bool("d", false, "help me!!!")
	flag.Parse()
	go system.HandleControllerMessages()
	system.InitializeSystemStatus(*debug)
	if !(*debug) {
		status := system.GetSystemStatus()
		defer status.Port.Close()
	}
	database.InitDB()
	defer database.DB.Close()
	database.CreateAccount("hello", "world", "again")
	http.HandleFunc("/", handlers.LoginHandler)
	http.Handle("/app/", handlers.GenericHandler{GET: handlers.AppHandler})

	http.Handle("/demo/start/", handlers.GenericHandler{GET: handlers.DemoStartHandler})
	http.Handle("/demo/stop/", handlers.GenericHandler{GET: handlers.DemoStopHandler})

	http.Handle("/demo/averaging/", handlers.GenericHandler{PUT: handlers.AveragingHandler})

	http.Handle("/system/", handlers.GenericHandler{GET: handlers.SystemStatusHandler})
	http.Handle("/system/speaker/", handlers.GenericHandler{PUT: handlers.SpeakerPutHandler})

	//http.Handle("demo/diagnostics/", handlers.GenericHandler{PUT: handlers.DiagnosticsHandler})

	// file handlers
	http.Handle("/css/", handlers.CssHandler)
	http.Handle("/bower_components/", handlers.BowerHandler)
	http.Handle("/components/", handlers.ComponentsHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
}
