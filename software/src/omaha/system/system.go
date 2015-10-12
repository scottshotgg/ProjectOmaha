package system

var status SystemStatus

type SystemStatus struct {
	ledOn bool
	debug bool
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
		status.ledOn, _ = status.GetLEDStatusFromController()
	}
	return &status
}

func GetSystemStatus() *SystemStatus {
	return &status
}
