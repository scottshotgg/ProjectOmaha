package system

func (status *SystemStatus) TurnLEDOn() error {
	status.ledOn = true
	if status.debug {
		return nil
	}
	status.SendMessageHeader()
	status.SendData([]byte{0x56})	// This is the command, outdated though
	status.SendData([]byte{0x00})	// This is the "data" for the command
	return nil

}

func (status *SystemStatus) TurnLEDOff() error {
	status.ledOn = false
	if status.debug {
		return nil
	}
	status.SendMessageHeader()
	status.SendData([]byte{0x76})
	status.SendData([]byte{0x00})
	return nil
}

func (status *SystemStatus) GetLEDStatusFromController() (bool, error) {
	if status.debug {
		return true, nil
	}
	// Write 4 bytes to the port.
	//var a = "a"
	status.SendMessageHeader()

	b := []byte{0x6c} //, 0x01, 0x02, 0x03}
	status.SendData(b)

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
	status.SendData([]byte{0x76})	// Uppercase V sets the volume
	status.SendData([]byte{int8(volumeLevel)})		// If compile fails, put back byte()

	return nil
}

func (status *SystemStatus) GetVolumeFromController() (int, error) {
	if status.debug {
		return 0, nil
	}

	status.SendMessageHeader()
	status.SendData([]byte{0x56})		// lowercase v gets the volume
	status.SendData([]byte{0x00})

	b := []byte{0x00}
	status.ReadData(b)

	return nil
}

func (status *SystemStatus) ResetFIFO() { 		// This function resets the FIFO on the micrcontrollers if one ever gets stuck
	if status.debug {
		return 0, nil
	}

	for i := 0; i < 8; i++{
		status.SendData([]byte{0x00})			// Send all zeros, this will reset their FIFO, [0][0][command][data] could also mean that everyone listens
	}

	return nil;
}

func (status *SystemStatus) ResetMicrocontroller() {
	if status.debug {
		return 0, nil
	}
	
	status.SendMessageHeader()
	status.SendData([]byte{0x00})
	status.SendData([]byte{0x00})
	
	return nil
	
}

func (status *SystemStatus) ChangeAveragingFilter(filter int) {
	if status.debug {
		return 0, nil
	}ssssssssss
	
	status.SendMessageHeader()
	 
	
	
	
	return nil
}