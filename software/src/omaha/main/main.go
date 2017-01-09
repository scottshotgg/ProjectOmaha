package main

import (
	"flag"
	"log"
	"time"
	"net/http"
	"omaha/database"
	"omaha/handlers"
	"omaha/system"
    //"sync"
)

/*
    schedule is a function that is used to schedule the keepAlive function that will send every second to a new microcontroller
*/
func schedule(keepAlive func(ID int8) (int8), delay time.Duration, amountOfControllers int8, controllers []*database.ControllerStatus) chan bool {
    stop := make(chan bool)
    var controller int8 = 0
    status := system.GetSystemStatus()
    go func() {
        for {
            second := time.After(delay)
            select {
                case <- second:
                    status.ID = controller + 1

                    status.BrokenLink = keepAlive(controller)

                    if(status.BrokenLink != int8(controllers[controller].Status)) {
                        database.UpdateStatus(status.ID, status.BrokenLink)
                    }

                    if(status.BrokenLink == 0) {
                        controller++            
                    } else {
                        // something is wrong with the controller 
                    }

                    if(controller > amountOfControllers - 1) {
                        controller = 1
                        status.ID = 1
                    }

                case <- stop:
                    return
            }
        }
    }()

    return stop
}

/*
    StartKeepAlive is the calling function that calls the scheduling after finding all the information.
*/
func startKeepAlive(amountOfControllers int, controllers []*database.ControllerStatus) {
	keepAlive := func(ID int8) (int8) { 
		var returnValue int8 = 0
		returnValue = system.KeepAlive(ID)
		return returnValue
	}
    schedule(keepAlive, 1 * time.Second, int8(amountOfControllers), controllers)
}

/*
    This is the main function. This is where it all starts. Main is responsible for parsing the command line flags, spawning threads to do different tasks, initializing the DB, and initializing the handlers, and starting the keelAlive. 
*/
func main() {
	log.Println("starting main function in golang")

	var controllerAmount = flag.Int("c", -1, "help me!!!")  
    var debug = flag.Bool("d", false, "help me!!!")
	flag.Parse()
	initializeLogger()
	go system.HandleControllerMessages()

	database.InitDB()
	defer database.DB.Close()

    SystemStatus, amountOfControllers, controllers := system.InitializeSystemStatus(*debug)
    if(!*debug) {
        amountOfControllers = *controllerAmount
    }

    log.Println("this is me", amountOfControllers)

	if !(*debug) {
		status := system.GetSystemStatus()
		defer status.Port.Close()
	}

	log.Println(SystemStatus, amountOfControllers)

	handlers.InitializeHandlers()
	
    /*if(*debug == false) {
        var wg sync.WaitGroup
        wg.Add(1)
        system.StartUpProcess(&wg)
        wg.Wait()
    }*/
	
	log.Println("Starting KeepAlive sequence")
	startKeepAlive(amountOfControllers, controllers)

	log.Println("Starting server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
}

func initializeLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
