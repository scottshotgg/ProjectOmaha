package system

// these are abstracted so that the code is more readable. i.e, contoller.SetVolume(speakerID) vs controller.write(0x06, speakerID)
var Commands struct {
	KeepAlive							byte
	SetVolume          		byte
	GetVolume          		byte
	SetSoundMaskingVolume	byte
	GetSoundMaskingVolume	byte
	SetMusicVolume				byte
	SetPaging							byte
	SetFadeTime						byte
	SetFadeLevel					byte
	AreYouAlive        		byte
	ResetSpeaker       		byte
	ResetFIFO          		byte
	GetLEDStatus       		byte
	TurnLEDOn          		byte
	TurnLEDOff         		byte
	GetSPL             		byte
	SetFilter          		byte
	GetFilter          		byte
	SetAveragingFilter 		byte
	GetSpeakerVoltage  		byte
	GetMCUVoltage      		byte
	InitializeFilter   		byte
	SetID              		byte
	GetID              		byte
	SetEQMode							byte
}

/*
	init is a private function that is used when assigning the values to the variables for the serial commands.
*/
func init() {
	Commands.KeepAlive = 							byte('K')
	Commands.SetVolume = 							byte('S')
	Commands.GetVolume = 							byte('s')
	Commands.SetSoundMaskingVolume = 	byte('W')
	Commands.GetSoundMaskingVolume = 	byte('w')
	Commands.SetMusicVolume = 				byte('M')
	Commands.SetPaging = 							byte('p')
	Commands.AreYouAlive = 						byte('U')
	Commands.ResetSpeaker = 					byte('R')
	Commands.ResetFIFO = 							byte('r')
	Commands.GetLEDStatus = 					byte('l')
	Commands.GetSPL = 								byte('m')
	Commands.SetFilter = 							byte('I')
	Commands.GetFilter = 							byte('i')
	Commands.SetAveragingFilter = 		byte('a')
	Commands.GetSpeakerVoltage = 			byte('V')
	Commands.GetMCUVoltage = 					byte('v')
	Commands.InitializeFilter = 			byte('f')
	Commands.SetID = 									byte('W')
	Commands.GetID = 									byte('w')
	Commands.SetEQMode =							byte('P')
}
