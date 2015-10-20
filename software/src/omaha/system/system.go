package system

import (
	"fmt"
	"io"
	"strconv"
)

var status SystemStatus

type SystemStatus struct {
	debug       bool                         `json:"-"`
	Port        io.ReadWriteCloser           `json:"-"`
	Controllers map[string]*ControllerStatus `json:"controllers"` // controllers mapped by their ID
}

// Do we want to store this stuff in a database since that way it will be written to disk?
// It may also be easier to access (programmatically, at least) that way
func (status *SystemStatus) IsDebug() bool {
	return status.debug
}

func (status *SystemStatus) GetController(ID int8) *ControllerStatus {
	return status.Controllers[strconv.Itoa(int(ID))]
}

func (status *SystemStatus) AddController(controller *ControllerStatus) {
	status.Controllers[strconv.Itoa(int(controller.ID))] = controller
}

func InitializeSystemStatus(isDebug bool) *SystemStatus {
	status.debug = isDebug
	if !isDebug {
		status.InitializePort()
	}
	status.Controllers = make(map[string]*ControllerStatus)

	controller := &ControllerStatus{ID: 0x6B, SectionID: 0x6B}
	var err error
	controller.LEDOn, err = controller.GetLEDStatusFromController()
	if err != nil {
		fmt.Println(err)
	}
	controller.VolumeLevel, err = controller.GetVolumeFromController()
	if err != nil {
		fmt.Println(err)
	}

	status.AddController(controller)
	return &status
}

func GetSystemStatus() *SystemStatus {
	return &status
}
