package system

import (
	"io"
	"log"
	"omaha/database"
	"strconv"
)

var status SystemStatus

type SystemStatus struct {
	debug       bool                                  `json:"-"`
	Port        io.ReadWriteCloser                    `json:"-"`
	Controllers map[string]*database.ControllerStatus `json:"controllers"` // controllers mapped by their ID
}

// Do we want to store this stuff in a database since that way it will be written to disk?
// It may also be easier to access (programmatically, at least) that way
func (status *SystemStatus) IsDebug() bool {
	return status.debug
}

func (status *SystemStatus) GetController(ID int8) *database.ControllerStatus {
	return database.GetSpeaker(ID)
}

func (status *SystemStatus) AddController(controller *database.ControllerStatus) {
	status.Controllers[strconv.Itoa(int(controller.ID))] = controller
}

func InitializeSystemStatus(isDebug bool) *SystemStatus {
	status.debug = isDebug
	if !isDebug {
		status.InitializePort()
	}
	status.Controllers = make(map[string]*database.ControllerStatus)

	controllers := database.GetAllSpeakers()

	for _, controller := range controllers {
		var err error

		controller.LEDOn, err = GetLEDStatusFromController(controller)
		if err != nil {
			log.Println(err)
		}

		controller.VolumeLevel, err = GetVolumeFromController(controller)
		if err != nil {
			log.Println(err)
		}

		//	controller.AveragingMode, err = controller.GetAveragingMode()		// Something needs to be fixed
		if err != nil {
			log.Println(err)
		}
		status.AddController(controller)
	}
	return &status
}

func GetSystemStatus() *SystemStatus {
	return &status
}