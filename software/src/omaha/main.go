package main

import (
	"flag"
	"log"
	"time"
	"net/http"
	"omaha/database"
	"omaha/handlers"
	"omaha/system"
)

func schedule(keepAlive func(ID int8) (int8), delay time.Duration, amountOfControllers int8) chan bool {
    stop := make(chan bool)
    var controller int8 = 1

    status := system.GetSystemStatus()

    go func() {
        for {
        	status.ID = controller
           
            status.BrokenLink = keepAlive(controller)

            // repeat the keepAlive until the controller responds, we can also set a bounding here so that it will not be tested regularly after awhile
        	if(status.BrokenLink == 0) {
            	controller++			// if we get a 0 then move onto the next controller
        	} else {
        		// at this point we should have a function jump that will use a case and figure out wtf is wrong with the controller, and then we can
        		// add it to a skipped array (or make a new data structure to support this) if there is UI consent/acknowledgment 
        	}

            if(controller > amountOfControllers) {
            	controller = 1
            }

            select {
            case <-time.After(delay):
            case <-stop:
                return
            }
        }
    }()

    return stop
}

func startKeepAlive(amountOfControllers int) {
	keepAlive := func(ID int8) (int8) { 
		var returnValue int8 = 0
		returnValue = system.KeepAlive(ID)
		return returnValue
	}
    schedule(keepAlive, 1 * time.Second, int8(amountOfControllers))
}

func main() {
	// initialization
	var debug = flag.Bool("d", false, "help me!!!")
	flag.Parse()
	initializeLogger()
	go system.HandleControllerMessages()

	database.InitDB()
	defer database.DB.Close()

	SystemStatus, amountOfControllers := system.InitializeSystemStatus(*debug)
	if !(*debug) {
		status := system.GetSystemStatus()
		defer status.Port.Close()
	}

	log.Println(SystemStatus, amountOfControllers)

	handlers.InitializeHandlers()

	log.Println("Starting KeepAlive sequence")
	startKeepAlive(amountOfControllers)

	log.Println("Starting server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
}

func initializeLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
