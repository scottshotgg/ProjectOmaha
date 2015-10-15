package system

func (status *SystemStatus) TurnLEDOn() error {
	status.ledOn = true
	if status.debug {
		return nil
	}
	status.SendMessageHeader(1)
	b := []byte{0x56} //, 0x01, 0x02, 0x03}		// Using V and v from volume now
	status.SendData(b)
	return nil

}

func (status *SystemStatus) TurnLEDOff() error {
	status.ledOn = false
	if status.debug {
		return nil
	}
	status.SendMessageHeader(1)
	b := []byte{0x76} //, 0x01, 0x02, 0x03}		// Change this
	status.SendData(b)
	return nil
}

func (status *SystemStatus) GetLEDStatusFromController() (bool, error) {
	// Make sure to close it later.

	// Write 4 bytes to the port.
	//var a = "a"
	status.SendMessageHeader(1)

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

func (status *SystemStatus) VolumeUp() error {
	status.volumeLevel++
	status.SendMessageHeader(1)
	b := []byte{0x56}
	status.SendData(b)

	return nil
}

func (status *SystemStatus) VolumeDown() error {
	status.volumeLevel--
	status.SendMessageHeader(1)
	b := []byte{0x76}
	status.SendData(b)

	return nil
}

func (status *SystemStatus) VolumeVariable(volumeLevel int) error {
	status.volumeLevel = volumeLevel
	status.SendMessageHeader(2)
	b := []byte{0x6D}
	status.SendData(b)
	b := []byte{int8(volumeLevel)} // We need to get the volume level integer and send that
	status.SendData(b)

	return nil
}
