package system

import (
	"fmt"
	"strconv"
)

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

type ControllerStatus struct {
	LEDOn         bool `json:"ledOn"`
	VolumeLevel   int8 `json:"volumeLevel"`
	ID            int8 `json:"id"`
	SectionID     int8 `json:"sectionId"`
	AveragingMode int8 `json:"averagingMode"`
}

func (this *ControllerStatus) IsLEDOn() bool {
	return this.LEDOn
}

func (this *ControllerStatus) TurnLEDOn() error {
	data := getMessageHeader(this.SectionID, this.ID, 4)
	data[2] = Commands.TurnLEDOn
	data[3] = 0x00

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		this.LEDOn = true
		if status.IsDebug() {
			fmt.Println("LED turned on")
		}
		return nil
	}}
	MessageChan <- req

	return nil

}

func (this *ControllerStatus) TurnLEDOff() error {
	data := getMessageHeader(this.SectionID, this.ID, 4)
	data[2] = Commands.TurnLEDOff
	data[3] = 0x01

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		this.LEDOn = false
		if status.IsDebug() {
			fmt.Println("LED turned off")
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

type LEDStatusResponse struct {
	ledOn bool
}

func (this *ControllerStatus) GetLEDStatusFromController() (bool, error) {
	data := getMessageHeader(this.SectionID, this.ID, 4)
	data[2] = Commands.GetLEDStatus
	data[3] = 0x00

	ch := make(chan interface{})
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		response := &LEDStatusResponse{}
		if status.IsDebug() {
			fmt.Println("Got LED Status From Controller")
			response.ledOn = true
		} else {
			var b []byte
			status.ReadData(b)
			response.ledOn = b[0] != 0x01
		}
		return response
	}, ResultChan: ch}

	MessageChan <- req
	response := (<-ch).(*LEDStatusResponse)

	return response.ledOn, nil
}

func (this *ControllerStatus) SetVolume(volumeLevel int8) error {
	data := getMessageHeader(this.SectionID, this.ID, 4)
	data[2] = Commands.SetVolume
	data[3] = byte(volumeLevel)

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		this.VolumeLevel = volumeLevel
		if status.IsDebug() {
			fmt.Printf("Set volume to %d\n", volumeLevel)
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func (this *ControllerStatus) GetVolumeFromController() (int8, error) {
	if status.debug {
		return 0, nil
	}

	data := getMessageHeader(this.SectionID, this.ID, 4)
	data[2] = Commands.GetVolume
	data[3] = 0x00

	// status.SendData(filter) commented out until filter is created

	return 0, nil
}

func (this *ControllerStatus) SetAveragingMode(mode int8) error {
	data := getMessageHeader(this.SectionID, this.ID, 4)
	data[2] = Commands.SetAveragingFilter
	data[3] = byte(mode)

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			fmt.Printf("Set averaging mode to %d\n", mode)
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func (status *SystemStatus) ResetFIFO() (int, error) { // This function resets the FIFO on the micrcontrollers if one ever gets stuck
	if status.debug {
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

	data := getMessageHeader(0, 0, 4)
	data[2] = 0x00
	data[3] = 0x00

	return 0, nil

}

func (this *ControllerStatus) GetAveragingMode() (int, error) {
	if status.debug {
		return 0, nil
	}

	data := getMessageHeader(0, 0, 4)
	data[2] = 0x61
	data[3] = 0x00

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
