package system

var Commands struct {
	SetVolume			byte
	GetVolume			byte
	TestAlive			byte
	ResetSpeaker 		byte
	ResetFIFO			byte
	GetLEDStatus		byte
	TurnLEDOn			byte
	TurnLEDOff			byte
	GetSPL				byte
	SetFilter			byte
	SetAveragingFilter	byte
	GetSpeakerVoltage	byte
	GetMCUVoltage		byte
	InitializeFilter	byte
}

func init() {											// Can we not just set these up ^ there so that we don't 
														// have to write all the commands out twice?
	Commands.SetVolume 			= byte('S')
	Commands.GetVolume 			= byte('s') 
	Commands.TestAlive 			= byte('U') 
	Commands.ResetSpeaker 		= byte('R')
	Commands.ResetFIFO 			= byte('r') 
	Commands.GetLEDStatus		= byte('l')
	Commands.TurnLEDOn 			= byte('L')
	Commands.TurnLEDOff 		= Commands.TurnLEDOn	// This command changes based on the data that you send it, 0 and 1 for on and off
	Commands.GetSPL 			= byte('m') 			// get microphone level, essentially
	Commands.SetFilter 			= byte('i')				// This needs to have a O or o to indicate the filter sent with it
	Commands.SetAveragingFilter = byte('a')
	Commands.GetSpeakerVoltage 	= byte('V')			 	// This may or may not be implemented
	Commands.GetMCUVoltage 		= byte('v') 			// This may or may not be implemented
	Commands.InitializeFilter 	= byte('f')

}
