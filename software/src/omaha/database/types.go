package database

type Account struct {
	UID      int
	UserName string
	Hash     string
}

type ControllerStatus struct {
	LEDOn         bool `json:"ledOn"`
	VolumeLevel   int8 `json:"volumeLevel"`
	ID            int8 `json:"id"`
	X             int  `json:"x"`
	Y             int  `json:"y"`
	SectionID     int8 `json:"sectionId"`
	AveragingMode int8 `json:"averagingMode"`
}
