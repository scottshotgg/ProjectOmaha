package system

import (
	"fmt"
	"io"
)

var status SystemStatus

type SystemStatus struct {
	LEDOn         bool               `json:"ledOn"`
	debug         bool               `json:"-"`
	Port          io.ReadWriteCloser `json:"-"`
	volumeLevel   int8               `json:"volumeLevel"`
	numberOfNodes int
}

// Do we want to store this stuff in a database since that way it will be written to disk?
// It may also be easier to access (programmatically, at least) that way

func (status *SystemStatus) IsLEDOn() bool {
	return status.LEDOn
}

func (status *SystemStatus) IsDebug() bool {
	return status.debug
}

func (status *SystemStatus) GetVolumeLevel() int8 {
	return status.VolumeLevel
}

func InitializeSystemStatus(isDebug bool) *SystemStatus {
	var err error
	status.debug = isDebug
	status.LEDOn, err = status.GetLEDStatusFromController()
	if err != nil {
		fmt.Println(err)
	}
	status.VolumeLevel, err = status.GetVolumeFromController()
	if err != nil {
		fmt.Println(err)
	}

	return &status
}

func GetSystemStatus() *SystemStatus {
	return &status
}
