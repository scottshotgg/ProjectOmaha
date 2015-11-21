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
	AveragingMode int8 `json:"averagingMode"`
}

type Zone struct {
	VolumeLevel int8                `json:"-"`
	ID          int8                `json:"zoneID"`
	Name        string              `json:"name"`
	Speakers    []*ControllerStatus `json:"speakers"`
}
