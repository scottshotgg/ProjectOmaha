package system

import (
	"io"
	"log"
	"omaha/database"
)

var status SystemStatus

type SystemStatus struct {
	debug				bool								`json:"-"`
	finding			bool								`json:"-"`
	Port 				io.ReadWriteCloser	`json:"-"`
	ID 					int8								`json:"id"`
	BrokenLink	int8								`json:"brokenLink"`
}

func (status *SystemStatus) IsDebug() bool {
	return status.debug
}

func (status *SystemStatus) IsFinding() bool {
	return status.finding
}

func (status *SystemStatus) SetFinding(finding bool) {
	status.finding = finding
}

func InitializeSystemStatus(isDebug bool) (*SystemStatus, int, []*database.ControllerStatus) {
	status.debug = isDebug
	controllers := database.GetAllSpeakers()
	length := len(controllers)

	if !isDebug {
		status.InitializePort()
	}

	log.Println("length", length)

	return &status, length, controllers
}

func GetSystemStatus() *SystemStatus {
	return &status
}
