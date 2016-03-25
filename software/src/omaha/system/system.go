package system

import (
	"io"
	"log"
	"omaha/database"
)

var status SystemStatus

type SystemStatus struct {
	debug		bool				`json:"-"`
	Port 		io.ReadWriteCloser	`json:"-"`
	ID 			int8				`json:"id"`
	BrokenLink	int8				`json:"brokenLink"`
	// damaged controllers in an array here
}

// Do we want to store this stuff in a database since that way it will be written to disk?
// It may also be easier to access (programmatically, at least) that way
func (status *SystemStatus) IsDebug() bool {
	return status.debug
}

func InitializeSystemStatus(isDebug bool) (*SystemStatus, int, []*database.ControllerStatus) {		// this might need to change to controllers and return the entire object/pointer 
	status.debug = isDebug
	controllers := database.GetAllSpeakers()
	length := len(controllers)

	if !isDebug {
		status.InitializePort()
	}

	/*for _, controller := range controllers {
		var err error

		controller.LEDOn, err = GetLEDStatusFromController(controller)
		if err != nil {
			log.Println(err)
		}

	//	controller.VolumeLevel, err = GetVolumeFromController(controller)		fix this at some later time but not now
		if err != nil {
			log.Println(err)
		}

		//	controller.AveragingMode, err = controller.GetAveragingMode()		// Something needs to be fixed
		if err != nil {
			log.Println(err)
		}
		//database.SaveVolume(controller)
	}*/

	log.Println("length", length)

	return &status, length, controllers
}

func GetSystemStatus() *SystemStatus {
	return &status
}
