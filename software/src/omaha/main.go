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

    // split off a thread for an anonymous function that loops forever (for "ever" {}) and try sending and receiving from the controllers to get status updates. 
    go func() {
        for {
            second := time.After(delay)
            select {
                // After a second has passed by, this case will activate
                case <- second:
                    status.ID = controller + 1

                    // try sending the keep alive, this function is in controller.go
                    status.BrokenLink = keepAlive(controller)

                    // if the status that we have received back is different than the status that the controller had before, then update the database.
                    if(status.BrokenLink != int8(controllers[controller].Status)) {
                        database.UpdateStatus(status.ID, status.BrokenLink)
                    }

                    // if the controller is still functioning, increment like normal
                    if(status.BrokenLink == 0) {
                        controller++            
                    } else {
                        log.Println(controller, "has something wrong with it")
                        // something is wrong with the controller, this is where we would do debugging to see if something useful was sent back
                        // if nothing useful, we can try restarting the controller 
                    }

                    if(controller > amountOfControllers - 1) {
                        // if we loop past the amount of controllers, reset the values back to the start
                        controller = 1
                        status.ID = 1
                    }

                // This would activate if we ever put stop into the channel, but we aren't using this
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
    // Parse controller amount flag if there is one. This flag is used to force the amount of controllers. The effects of this are implemented with regards to the process of finding new controllers.
	var controllerAmount = flag.Int("c", -1, "help me!!!") 
    // Parse the debug flag 
    var debug = flag.Bool("d", false, "help me!!!")
	flag.Parse()
	initializeLogger()

    // Split off a process to handle the controller messages
	go system.HandleControllerMessages()

    // init the DB and defer closing it until the end of the program
	database.InitDB()
	defer database.DB.Close()

    SystemStatus, amountOfControllers, controllers := system.InitializeSystemStatus(*debug)
    if(!*debug) {
        amountOfControllers = *controllerAmount
    }

	if !(*debug) {
		status := system.GetSystemStatus()
		defer status.Port.Close()
	}

	log.Println(SystemStatus, amountOfControllers)

	handlers.InitializeHandlers()

    //go system.ControllerSearch()

    /*if(*debug == false) {
        var wg sync.WaitGroup
        wg.Add(1)
        system.StartUpProcess(&wg)
        wg.Wait()
    }*/

    // Start the keep alive sequence
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
