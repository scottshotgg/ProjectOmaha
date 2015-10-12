package system

import (
	"fmt"
)

func (status *SystemStatus) TurnLEDOn() error {
	if status.debug {
		status.ledOn = true
		return nil
	}

	b := []byte{0x01} //, 0x01, 0x02, 0x03}
	status.SendData(b)

	status.ReadData(b)

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)

	if s == "1" {
		return nil
	} else {
		return nil
	}

}

func (status *SystemStatus) TurnLEDOff() error {
	if status.debug {
		status.ledOn = false
		return nil
	}

	b := []byte{0x00} //, 0x01, 0x02, 0x03}
	status.SendData(b)

	status.ReadData(b)

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)
	if s == "1" {
		return nil
	} else {
		return nil
	}

}

func (status *SystemStatus) GetLEDStatusFromController() (bool, error) {
	// Make sure to close it later.

	// Write 4 bytes to the port.
	//var a = "a"
	b := []byte{0x6c} //, 0x01, 0x02, 0x03}
	status.SendData(b)

	status.ReadData(b)

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)
	if s == "1" {
		return true, nil
	} else {
		return false, nil
	}
}
