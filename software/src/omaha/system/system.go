package system

import (
	"fmt"
	"io"
)

var status SystemStatus

type SystemStatus struct {
	ledOn bool
	debug bool
	Port  io.ReadWriteCloser
}

func (status *SystemStatus) IsLEDOn() bool {
	return status.ledOn
}

func (status *SystemStatus) IsDebug() bool {
	return status.debug
}

func InitializeSystemStatus(isDebug bool) *SystemStatus {
	if isDebug {
		status.ledOn = true
		status.debug = true
	} else {
		status.InitializePort()
		var err error
		status.ledOn, err = status.GetLEDStatusFromController()
		if err != nil {
			fmt.Println(err)
		}
	}
	return &status
}

func GetSystemStatus() *SystemStatus {
	return &status
}
