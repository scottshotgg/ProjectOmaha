/*
 	The system package handles the interfacing of the system with the hardware that it supports. It supports the queuing messaging system, as well as the entire protocol that we use for sending messages.
	The protocol works by sending three bytes to communicate with the controllers/speakers.
	[0]:	address, 0-254
	[1]:	command, this is taken from the constants.go file
	[2]: 	data, this is where you would fill in the volume to set it to, or the amount to raise the band
*/
package system

import (
	"log"
	"omaha/database"
)

const NULLZone int = 0

/*
	KeepAlive is a function that is used by the scheduler in main.go to send a KeepAlive command to the speakers.
*/
func KeepAlive(ID int8) (error int8) {
	// apply the message header and then fill in the data and commands
	// KeepAlive sends a 1 and expects to get a capital 'A' back. 
	data := getMessageHeader(ID, 3)
	data[1] = Commands.KeepAlive
	data[2] = 0x01

    defer func() {
    	if r := recover(); r != nil {
        	log.Println("Error reading from the port for controller ", ID, "\n", r)
    		error = -ID
    	}
    }()

	log.Println("KeepAlive: ", data)
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			return nil
		}
		return nil
	}}
	MessageChan <- req

	if(!status.IsDebug()){
		b := []byte{0x00}			
		status.ReadData(b)

		log.Println("Return packet: ", b[0])

		if(int8(b[0]) == int8('A')) {
			log.Println(int8('A'))
			error = 0
		} else {
			error = int8(b[0])
		}
	}

	return 
}

/*
	SetVolume is used to send the set volume command to the specified speaker.
*/
func SetVolume(this *database.ControllerStatus) error {
	// fill in the header and then set the command and fill in the volume level that you would like to set it to
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetVolume
	data[2] = byte(this.VolumeLevel[0])

	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Printf("Set volume to %d\n", this.VolumeLevel[0])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

/*
	GetVolumeFromController is used to query the controller for its current volume variable should the controller ever need it.
*/
func GetVolumeFromController(this *database.ControllerStatus) (int8, error) {
	if status.debug {
		return 0, nil
	}

	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.GetVolume
	data[2] = 0x00

	return 0, nil
}

/*
	SetMusicVolume is used to set the music volume on the speaker.
*/
func SetMusicVolume(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetMusicVolume
	data[2] = byte(this.VolumeLevel[1])

	log.Println("Packet contents: ", data)

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Printf("Set music volume to %d\n", this.VolumeLevel[1])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

/*
	SetPaging is used to send the paging constants to the speaker.
*/
func SetPaging(this *database.ControllerStatus, pagingData int8) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetPaging
	data[2] = byte(pagingData)

	log.Println("Packet contents: ", data)

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {

		if status.IsDebug() {
			log.Printf("Set paging data to %d\n", pagingData)
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

/*
	SetSoundMaskingVolume sets how loud the speaker should be playing noise.
*/
func SetSoundMaskingVolume(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetSoundMaskingVolume
	data[2] = byte(this.VolumeLevel[3])

	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Printf("Set sound masking volume to %d\n", this.VolumeLevel[3])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

/*
	SetAveragingMode sets the averaging mode of the speaker, which consists of effectiveness and pleasantness.
*/
func SetAveragingMode(this *database.ControllerStatus, mode int8) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetAveragingFilter
	data[2] = byte(mode)

	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Printf("Set averaging mode to %d\n", mode)
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

/*
	ResetFIFO is being reserved if need be, but will most likely not be used and isn't used right now.
*/
func (status *SystemStatus) ResetFIFO() (int, error) { 
	if status.debug {
		return 0, nil
	}

	return 0, nil
}

/*
	ResetMicrocontroller is being reserved if we need it. This will most likely come to fruition when we get into doing more complicated diagnostics and on the fly automated debugging.
*/
func (status *SystemStatus) ResetMicrocontroller() (int, error) {
	if status.debug {
		return 0, nil
	}

	data := getMessageHeader(0, 4)
	data[1] = 0x00
	data[2] = 0x00

	return 0, nil

}

/*
	GetAveragingMode is used to query the specified speaker for its averaging mode.
*/
func GetAveragingMode(this *database.ControllerStatus) (int, error) {
	if status.debug {
		return 0, nil
	}

	data := getMessageHeader(0, 4)
	data[1] = 0x61
	data[2] = 0x00

	b := []byte{0x00}
	status.ReadData(b)

	return 0, nil

}

/*
	SetEqualizerConstant is used to send one equalizer constant for the specified band to a speaker.
*/
func SetEqualizerConstant(this *database.ControllerStatus, level int8, band int8, decimal bool) (int, error) {
	data := getMessageHeader(this.ID, 3)
	data[1] = byte(band)
	if(decimal) {
		data[2] = byte(level)
	} else {
		data[2] = byte(level + 40)
	}
	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{ Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Printf("Set band %-2d mode to %3ddB (-40)   |   Packet contents: %2d", band, level, data)
		}
		return nil
	}}
	MessageChan <- req	

	return 0, nil
}

/*
	SetEQMode is used to change the EQ mode on the speaker to music, paging, or back to masking.
*/
func SetEQMode(ID int8, mode int8) (int, error) {
	data := getMessageHeader(ID, 3)
	data[1] = Commands.SetEQMode
	data[2] = byte(mode)

	req := &ControllerRequest{ Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Println("Set EQ mode to ",  mode)
		}
		return nil
	}}
	MessageChan <- req

	return 0, nil
}

/*
	SendPagingRequest is a reserved / unimplemented function that will be used to tell the speaker that a paging signal is coming in and to prepare for that.
*/
func SendPagingRequest(ID int8) (error) {
	if status.debug {
		return nil
	}

	data := getMessageHeader(ID, 3)
	data[1] = Commands.SetPaging
	data[2] = 0

	return nil
}




