package system

var Commands struct {
	SetVolume          byte
	GetVolume          byte
	TestAlive          byte
	SetAveragingFilter byte
	GetLEDStatus       byte
	TurnLEDOff         byte
	TurnLEDOn          byte
}

func init() {
	Commands.SetVolume = byte('S')		// Checking to see if this works
	Commands.GetVolume = byte('s') 
	Commands.TestAlive = 0x52 
	Commands.GetLEDStatus = 0x6C
	Commands.TurnLEDOff = 0x76
	Commands.TurnLEDOn = 0x56
	Commands.SetAveragingFilter = 0x61
}
