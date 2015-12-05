package system

var Commands struct {
	SetVolume          byte
	GetVolume          byte
	AreYouAlive        byte // Nto sure how to implement this, could split off another thread and make a ticker, http://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals-in-golang
	ResetSpeaker       byte
	ResetFIFO          byte
	GetLEDStatus       byte
	TurnLEDOn          byte
	TurnLEDOff         byte
	GetSPL             byte
	SetFilter          byte
	GetFilter          byte
	SetAveragingFilter byte
	GetSpeakerVoltage  byte
	GetMCUVoltage      byte
	InitializeFilter   byte
	SetID              byte
	GetID              byte
}

func init() { // Can we not just set these up ^ there so that we don't
	// have to write all the commands out twice?
	Commands.SetVolume = byte('S')
	Commands.GetVolume = byte('s')
	Commands.AreYouAlive = byte('U')
	Commands.ResetSpeaker = byte('R')
	Commands.ResetFIFO = byte('r')
	Commands.GetLEDStatus = byte('l')
	Commands.TurnLEDOn = byte('L')           // This needs to be changed to SetLED
	Commands.TurnLEDOff = Commands.TurnLEDOn // This command changes based on the data that you send it, 0 and 1 for on and off
	Commands.GetSPL = byte('m')              // get microphone level, essentially
	Commands.SetFilter = byte('I')           // This needs to have a O or o to indicate the filter sent with it
	Commands.GetFilter = byte('i')
	Commands.SetAveragingFilter = byte('a')
	Commands.GetSpeakerVoltage = byte('V') // This may or may not be implemented
	Commands.GetMCUVoltage = byte('v')     // This may or may not be implemented
	Commands.InitializeFilter = byte('f')
	Commands.SetID = byte('W')
	Commands.GetID = byte('w')
}
