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

/*
	IsDebug returns whether the system is acting in debug mode or not. The implications of this are in what kind of states the system can enter in to.
*/
func (status *SystemStatus) IsDebug() bool {
	return status.debug
}

/*
	IsFinding is used to mark whether the system is in the initial finding stage or not.
*/
func (status *SystemStatus) IsFinding() bool {
	return status.finding
}

/*
	SetFinding sets the finding variable when finding start and when it stops. This is used in main.go when starting the process to find the speakers.
*/
func (status *SystemStatus) SetFinding(finding bool) {
	log.Println("finding set")
	status.finding = finding
}

/*
	InitializeSystemStatus is used when initializing the systemm to the default status and processing extra command line arguments to determine whether it is debug mode.
*/
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

/*
	GetSystemStatus is used to access the status of the system from outside the system package.
*/
func GetSystemStatus() *SystemStatus {
	return &status
}
