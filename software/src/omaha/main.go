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

	database.InitDB()
	defer database.DB.Close()

	system.InitializeSystemStatus(*debug)
	if !(*debug) {
		status := system.GetSystemStatus()
		defer status.Port.Close()
	}

	http.HandleFunc("/login/", handlers.LoginPostHandler)
	http.HandleFunc("/", handlers.LoginPageHandler)

	http.Handle("/user/", handlers.GenericHandler{POST: handlers.AccountCreationHandler})

	http.Handle("/app/", handlers.GenericHandler{GET: handlers.AppHandler})

	http.Handle("/demo/start/", handlers.GenericHandler{PUT: handlers.SpeakerPutHandler})
	http.Handle("/demo/stop/", handlers.GenericHandler{PUT: handlers.SpeakerPutHandler})

	http.Handle("/system/", handlers.GenericHandler{GET: handlers.SystemStatusHandler})
	http.Handle("/system/speaker/", handlers.GenericHandler{PUT: handlers.SpeakerPutHandler})

	// file handlers
	http.Handle("/css/", handlers.CssHandler)
	http.Handle("/bower_components/", handlers.BowerHandler)
	http.Handle("/components/", handlers.ComponentsHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
}
