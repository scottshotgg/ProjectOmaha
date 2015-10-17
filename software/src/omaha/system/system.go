package system

import (
	"fmt"
	"io"
)

var status SystemStatus

type SystemStatus struct {
	ledOn       bool
	debug       bool
	Port        io.ReadWriteCloser
	volumeLevel int
}

func (status *SystemStatus) IsLEDOn() bool {
	return status.ledOn
}

func (status *SystemStatus) IsDebug() bool {
	return status.debug
}

func (status *SystemStatus) GetVolumeLevel() int {
	return status.volumeLevel
}

func InitializeSystemStatus(isDebug bool) *SystemStatus {
	var err error
	status.debug = isDebug
	status.ledOn, err = status.GetLEDStatusFromController()
	if err != nil {
		fmt.Println(err)
	}
	status.volumeLevel, err = status.GetVolumeFromController()
	if err != nil {
		fmt.Println(err)
	}

	return &status
}

func GetSystemStatus() *SystemStatus {
	return &status
}
