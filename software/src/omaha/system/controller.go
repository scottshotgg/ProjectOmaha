package system

import (
	"strconv"
	//"fmt"
)

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func (status *SystemStatus) TurnLEDOn() error {
	status.LEDOn = true
	if status.debug {
		return nil
	}
	status.SendMessageHeader()

	status.SendData(TurnLEDOnID) // This is the command, outdated though
	status.SendData(0x00)        // This is the "data" for the command
	return nil

}

func (status *SystemStatus) TurnLEDOff() error {
	status.LEDOn = false
	if status.debug {
		return nil
	}
	status.SendMessageHeader()
	status.SendData(TurnLEDOffID)
	status.SendData(0x00)
	return nil
}

func (status *SystemStatus) GetLEDStatusFromController() (bool, error) {
	if status.debug {
		return true, nil
	}
	status.SendMessageHeader()
	status.SendData(GetLEDStatusID)
	status.SendData(0x00)
	var b []byte
	status.ReadData(b)

	if b[0] == 0x01 {
		return false, nil
	} else {
		return true, nil
	}
	return true, nil
}

func (status *SystemStatus) SetVolume(volumeLevel int8) error {
	status.VolumeLevel = volumeLevel
	if status.debug {
		return nil
	}
	status.SendMessageHeader()         // Later we will want to put the zone AND unit ID into this I think
	status.SendData(SetVolumeID)       // Uppercase V sets the volume
	status.SendData(byte(volumeLevel)) // If compile fails, put back byte()
	return nil
}

func (status *SystemStatus) GetVolumeFromController() (int8, error) {
	if status.debug {
		return 0, nil
	}

	status.SendMessageHeader()
	status.SendData(0x41)
	// status.SendData(filter) commented out until filter is created

	return 0, nil
}

func (status *SystemStatus) ResetFIFO() (int, error) { // This function resets the FIFO on the micrcontrollers if one ever gets stuck
	if status.debug {
		return 0, nil
	}

	for i := 0; i < 8; i++ {
		status.SendData(0x00) // Send all zeros, this will reset their FIFO, [0][0][command][data] could also mean that everyone listens
	}

	return 0, nil
}

func (status *SystemStatus) ResetMicrocontroller() (int, error) {
	if status.debug {
		return 0, nil
	}

	status.SendMessageHeader()
	status.SendData(0x00)
	status.SendData(0x00)

	return 0, nil

}

func (status *SystemStatus) ChangeAveragingFilter(filter int) (int, error) {
	return 0, nil
}

func (status *SystemStatus) GetAveragingFilter(filter int) (int, error) {
	if status.debug {
		return 0, nil
	}

	status.SendMessageHeader()
	status.SendData(0x61)
	status.SendData(0x00)

	b := []byte{0x00}
	status.ReadData(b)

	return 0, nil

}

func (status *SystemStatus) SendCoefficientInformation(equalizedGain int8, decibal int8) (int8, error) { // Equalized gain or just raw gain and equalization can be done in here
	if status.debug {
		return 0, nil
	}

	status.SendData(byte(decibal))
	status.SendData(byte(equalizedGain))

	return 0, nil
}

func (status *SystemStatus) AreYouAlive(n map[int]string) (m map[int]string) { // Equalized gain or just raw gain and equalization can be done in here
	if status.debug {
		return m
	}

	m = make(map[int]string)

	for i := 1; i <= len(n); i++ { // This should have something passed to it that tells it what IDs are still
		// available, maybe the map itself and we can get the length of the map

		status.SendMessageHeader() // i would go here
		status.SendData(0x52)
		status.SendData(0x00)

		b := []byte{0x00}
		alive := status.ReadData(b)

		if alive != true || b[0] != 0x72 { // This currently does the opposite of what we want it to do. Also include the data send back if we can
			m[i] = strconv.Itoa(Btoi(alive)) + strconv.Itoa(int(b[0])) // This should map the string consisting of 0l where 0 is the bool representing alive and
			// l is the variable recieved by the master controller if we need to send
			// maybe we should print thing to make sure its working
		} /*else if alive == true{

		}*/
		// Return the map with the microcontrollers that failed in it and maybe why they failed
	}

	return m
}
