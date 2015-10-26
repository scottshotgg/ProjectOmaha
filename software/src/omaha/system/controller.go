package system

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
)

func (status *SystemStatus) TurnLEDOn() error {
	if status.debug {
		status.ledOn = true
		return nil
	}
	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
	}

	fmt.Println("Hello " + options.PortName)

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.

	// Write 4 bytes to the port.
	//	var a = "a"
	b := []byte{0x01} //, 0x01, 0x02, 0x03}
	//stop := []byte(0x00)
	//i, err := strconv.Atoi("a")
	port.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
		panic("Error in write")
	}
	//write("a", port)

	// := []byte("k")
	fmt.Println("Wrote: ", b)

	_, err = port.Read(b)

	if err != nil {
		log.Fatalf("port.Read: %v", err)
		panic("Error in read")
	}

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)
	defer port.Close()

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
	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
	}

	fmt.Println("Hello " + options.PortName)

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.

	// Write 4 bytes to the port.
	//var a = "a"
	b := []byte{0x00} //, 0x01, 0x02, 0x03}
	//stop := []byte(0x00)
	//i, err := strconv.Atoi("a")
	port.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
		panic("Failed in write")
	}
	//write("a", port)

	// := []byte("k")

	fmt.Println("Wrote: ", b)

	_, err = port.Read(b)

	if err != nil {
		log.Fatalf("port.Read: %v", err)
		panic("Failed in read")
	}

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)
	defer port.Close()
	if s == "1" {
		return nil
	} else {
		return nil
	}

}

func (status *SystemStatus) GetLEDStatusFromController() (bool, error) {
	if status.debug {
		status.ledOn = true
		return true, nil
	}
	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
	}

	fmt.Println("Hello " + options.PortName)

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.

	// Write 4 bytes to the port.
	//var a = "a"
	b := []byte{0x6c} //, 0x01, 0x02, 0x03}
	//stop := []byte(0x00)
	//i, err := strconv.Atoi("a")
	port.Write(b)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
		panic("Failed on write")
	}
	//write("a", port)

	// := []byte("k")

	fmt.Println("Wrote: ", b)

	_, err = port.Read(b)

	if err != nil {
		log.Fatalf("port.Read: %v", err)
		panic("Failed on read")
	}

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)
	defer port.Close()
	if s == "1" {
		return true, nil
	} else {
		return false, nil
	}
}
