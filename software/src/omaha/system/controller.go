package system

import (
	"fmt"
)

func (status *SystemStatus) TurnLEDOn() error {
	if status.debug {
		status.ledOn = true
		return nil
	}
	status.SendMessageHeader(1)
	b := []byte{0x56} //, 0x01, 0x02, 0x03}
	status.SendData(b)

	status.ReadData(b)

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)

	if s == "1" {
		return nil
	} else {
		return nil
	}
	return nil

}

func (status *SystemStatus) TurnLEDOff() error {
	if status.debug {
		status.ledOn = false
		return nil
	}
	status.SendMessageHeader(1)
	b := []byte{0x76} //, 0x01, 0x02, 0x03}
	status.SendData(b)

	status.ReadData(b)

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)
	if s == "1" {
		return nil
	} else {
		return nil
	}
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

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)
	if s == "1" {
		return false, nil
	} else {
		return true, nil
	}
	return true, nil
}
