package system

import (
	"log"
	"omaha/database"
	"strconv"
)

const NULLZone int = 0

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func IsLEDOn(this *database.ControllerStatus) bool {
	return this.LEDOn
}

func TurnLEDOn(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.TurnLEDOn
	data[2] = 0x01

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		this.LEDOn = true
		if status.IsDebug() {
			log.Println("LED turned on")
		}
		return nil
	}}
	MessageChan <- req

	return nil

}

func TurnLEDOff(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.TurnLEDOff
	data[2] = 0x00

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		this.LEDOn = false
		if status.IsDebug() {
			log.Println("LED turned off")
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

type LEDStatusResponse struct {
	ledOn bool
}

func GetLEDStatusFromController(this *database.ControllerStatus) (bool, error) {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.GetLEDStatus
	data[2] = 0x00

	ch := make(chan interface{})
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		response := &LEDStatusResponse{}
		if status.IsDebug() {
			log.Println("Got LED Status From Controller")
			response.ledOn = true
		} else {
			b := make([]byte, 1)
			status.ReadData(b)
			response.ledOn = b[0] != 0x01
		}
		return response
	}, ResultChan: ch}

	MessageChan <- req
	response := (<-ch).(*LEDStatusResponse)

	return response.ledOn, nil
}

func SetVolume(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetVolume
	data[2] = byte(this.VolumeLevel[0])

	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		//this.VolumeLevel[0] = volumeLevel		this did not make sense at the time and was not useful
		if status.IsDebug() {
			log.Printf("Set volume to %d\n", this.VolumeLevel[0])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func GetVolumeFromController(this *database.ControllerStatus) (int8, error) {
	if status.debug {
		return 0, nil
	}

	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.GetVolume
	data[2] = 0x00

	// status.SendData(filter) commented out until filter is created

	return 0, nil
}

func SetMusicVolume(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetMusicVolume
	data[2] = byte(this.VolumeLevel[1])

	log.Println("Packet contents: ", data)

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		//this.VolumeLevel[1] = musicVolumeLevel
		if status.IsDebug() {
			log.Printf("Set music volume to %d\n", this.VolumeLevel[1])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

// put a getter here

func SetPaging(this *database.ControllerStatus, pagingData int8) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetPaging

	// we could either check here, but i dont think this is the appropriate place for that
	// make the check where this is called 
	data[2] = byte(pagingData)

	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		//this.VolumeLevel[2] = pagingVolumeLevel	

		if status.IsDebug() {
			log.Printf("Set paging volume to %d\n", this.VolumeLevel[2])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func SetSoundMaskingVolume(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetSoundMaskingVolume
	data[2] = byte(this.VolumeLevel[3])

	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		//this.VolumeLevel[1] = musicVolumeLevel
		if status.IsDebug() {
			log.Printf("Set sound masking volume to %d\n", this.VolumeLevel[3])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

/*

func SetFadeTime(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetFadeTime

	// we could either check here, but i dont think this is the appropriate place for that
	// make the check where this is called 
	data[2] = byte(this.PagingLevel[0])

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		//this.PagingLevel[0] = pagingFadeTime
		if status.IsDebug() {
			log.Printf("Set paging fade time to %d\n", this.PagingLevel[0])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func SetFadeLevel(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetFadeLevel

	// we could either check here, but i dont think this is the appropriate place for that
	// make the check where this is called 
	data[2] = byte(this.PagingLevel[1])

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		//this.PagingLevel[1] = pagingFadeLevel
		if status.IsDebug() {
			log.Printf("Set paging fade level to %d\n", this.PagingLevel[1])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}
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

func (status *SystemStatus) ResetFIFO() (int, error) { // This function resets the FIFO on the micrcontrollers if one ever gets stuck
	if status.debug {		// may not be implemented
		return 0, nil
	}

	//data := make([]byte, 8)

	// for i := 0; i < 8; i++ {
	// 	status.SendData(0x00) // Send all zeros, this will reset their FIFO, [0][0][command][data] could also mean that everyone listens
	// }

	return 0, nil
}

func (status *SystemStatus) ResetMicrocontroller() (int, error) {
	if status.debug {
		return 0, nil
	}

	data := getMessageHeader(0, 4)
	data[1] = 0x00
	data[2] = 0x00

	return 0, nil

}

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

// Deprecated
func (status *SystemStatus) SendCoefficientInformation(equalizedGain int8, decibal int8) (int8, error) { // Equalized gain or just raw gain and equalization can be done in here
	if status.debug {
		return 0, nil
	}

	// status.SendData(byte(decibal))
	// status.SendData(byte(equalizedGain))

	return 0, nil
}

func SetEqualizerConstant(this *database.ControllerStatus, level int8, band int8) (int, error) {

	/*if status.debug {
		return 0, nil
	}*/

	data := getMessageHeader(this.ID, 3)	// zone, id
	data[1] = byte(band)		// Cannot use a command to do this
	data[2] = byte(level * 2 + 80)

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
func (status *SystemStatus) SetPaging(this *database.ControllerStatus, pagingLevel int8) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetPaging
	data[2] = byte(pagingLevel)

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		this.PagingLevel = pagingLevel
		if status.IsDebug() {
			log.Printf("Set volume to %d\n", pagingLevel)
		}
		return nil
	}}
	MessageChan <- req

	return nil
}
*/

func (status *SystemStatus) AreYouAlive(n map[int]string) (m map[int]string) { // Equalized gain or just raw gain and equalization can be done in here
	if status.debug {
		return m
	}

	m = make(map[int]string)

	for i := 1; i <= len(n); i++ { // This should have something passed to it that tells it what IDs are still
		// available, maybe the map itself and we can get the length of the map

		// status.SendMessageHeader() // i would go here
		// status.SendData(Commands.TestAlive)
		// status.SendData(0x00)

		b := []byte{0x00}
		alive := status.ReadData(b)

		if alive != true || b[0] != 0x72 {
			m[i] = strconv.Itoa(Btoi(alive)) + strconv.Itoa(int(b[0])) // This should map the string consisting of 0l where 0 is the bool representing alive and
			// l is the variable recieved by the master controller if we need to send
			// maybe we should print thing to make sure its working
			// In here if it isn't r then we also need to check if it is 'v', 'm', other crap
		} /*else if alive == true{

		}*/
		// Return the map with the microcontrollers that failed in it and maybe why they failed
	}

	return m
}
