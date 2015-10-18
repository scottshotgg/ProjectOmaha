package system

func (status *SystemStatus) TurnLEDOn() error {
	status.ledOn = true
	if status.debug {
		return nil
	}
	status.SendMessageHeader()
	status.SendData(0x56)	// This is the command, outdated though
	status.SendData(0x00)	// This is the "data" for the command
	return nil

}

func (status *SystemStatus) TurnLEDOff() error {
	status.ledOn = false
	if status.debug {
		return nil
	}
	status.SendMessageHeader()
	status.SendData(0x76)
	status.SendData(0x00)
	return nil
}

func (status *SystemStatus) GetLEDStatusFromController() (bool, error) {
	if status.debug {
		return true, nil
	}
	status.SendMessageHeader()
	status.SendData(0x6C)
	status.SendData(0x00)

	b := []byte{0x6c}
	status.ReadData(b)

	if b[0] == 0x01 {
		return false, nil
	} else {
		return true, nil
	}
	return true, nil
}

/*func (status *SystemStatus) VolumeUp() error {			// These won't be used
	status.volumeLevel++
	if status.debug {
		return nil
	}

	status.SendMessageHeader(1)
	status.SendData([]byte{0x56})

	return nil
}

func (status *SystemStatus) VolumeDown() error {		// These won't be used
	status.volumeLevel--
	if status.debug {
		return nil
	}

	status.SendMessageHeader(1)
	status.SendData([]byte{0x76})

	return nil
}*/

func (status *SystemStatus) SetVolume(volumeLevel int) error {
	status.volumeLevel = volumeLevel
	if status.debug {
		return nil
	}
	status.SendMessageHeader()		// Later we will want to put the zone AND unit ID into this I think
	status.SendData(0x56)	// Uppercase V sets the volume
	status.SendData(volumeLevel)		// If compile fails, put back byte()

	return nil
}

func (status *SystemStatus) GetVolume() (int, error) {
	if status.debug {
		return 0, nil
	}

	status.SendMessageHeader()
	status.SendData(0x76)		// lowercase v gets the volume
	status.SendData(0x00)

	b := []byte{0x00}
	status.ReadData(b)		// need to do something with this

	return 0, nil
}

func (status *SystemStatus) ResetFIFO() (int, error) { 		// This function resets the FIFO on the micrcontrollers if one ever gets stuck
	if status.debug {
		return 0, nil
	}

	for i := 0; i < 8; i++ {
		status.SendData(0x00)			// Send all zeros, this will reset their FIFO, [0][0][command][data] could also mean that everyone listens
	}

	return 0, nil;
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
	if status.debug {
		return 0, nil
	}
	
	status.SendMessageHeader()
	status.SendData(0x41)
	status.SendData(filter)
	
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

func (status *SystemStatus) SendCoefficientInformation(equalizedGain int, decibal int) (int, error) {		// Equalized gain or just raw gain and equalization can be done in here
	if status.debug {
		return 0, nil
	}

	status.SendData(decibal)
	status.SendData(equalizedGain)

	return 0, nil
}

func (status *SystemStatus) AreYouAlive(numberOfNodes int) (m map[int]bool) {		// Equalized gain or just raw gain and equalization can be done in here
	if status.debug {
		return m
	}

	for i := 1; i <= numberOfNodes; i++ {
		var m map[int]bool
		m = make(map[int]bool)

		status.SendMessageHeader()	// i would go here
		status.SendData(0x52)
		status.SendData(0x00)

		b := []byte{0x00}
		error := status.ReadData(b)

		if error == true && b[0] == 0x72 {
			m[i] = error		// Map error to key [i]
		} 
	}

	return m		// Return the map with the microcontrollers that failed in it and maybe why they failed

}

