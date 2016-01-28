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
	var serialPort = flag.Bool("s", false, "Something went wrong and this is a message from: main.go")
	flag.Parse()
	initializeLogger()
	go system.HandleControllerMessages()

	database.InitDB()
	defer database.DB.Close()

	system.InitializeSystemStatus(*debug, serialPort)
	if !(*debug) {
		status := system.GetSystemStatus()
		defer status.Port.Close()
	}

	handlers.InitializeHandlers()
	log.Println("Starting server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
}

func initializeLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
