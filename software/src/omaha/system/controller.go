package system

func (status *SystemStatus) TurnLEDOn() error {
	status.ledOn = true
	if status.debug {
		return nil
	}

	status.SendMessageHeader(1)
	status.SendData([]byte{0x56})
	return nil

}

func (status *SystemStatus) TurnLEDOff() error {
	status.ledOn = false
	if status.debug {
		return nil
	}

	status.SendMessageHeader(1)
	status.SendData([]byte{0x76})
	return nil
}

func (status *SystemStatus) GetLEDStatusFromController() (bool, error) {
	if status.debug {
		return true, nil
	}

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
	if status.debug {
		return nil
	}

	status.SendMessageHeader(1)
	status.SendData([]byte{0x56})

	return nil
}

func (status *SystemStatus) VolumeDown() error {
	status.volumeLevel--
	if status.debug {
		return nil
	}

	status.SendMessageHeader(1)
	status.SendData([]byte{0x76})

	return nil
}

func (status *SystemStatus) SetVolume(volumeLevel int) error {
	status.volumeLevel = volumeLevel
	if status.debug {
		return nil
	}
	status.SendMessageHeader(2)
	status.SendData([]byte{0x6D})
	status.SendData([]byte{byte(int8(volumeLevel))})

	return nil
}

func (status *SystemStatus) GetVolumeFromController() (int, error) {
	if status.debug {
		return 0, nil
	}

	return 0, nil
}
