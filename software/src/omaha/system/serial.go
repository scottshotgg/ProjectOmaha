package system

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
)

func (status *SystemStatus) SendMessageHeader(size int) {
	status.SendData([]byte{0x6C})
	status.SendData([]byte{0x01})
}

func (status *SystemStatus) SendData(data []byte) {
	_, err := status.Port.Write(data)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
		panic("Failed on write")
	}
}

func (status *SystemStatus) ReadData(buffer []byte) {
	_, err := status.Port.Read(buffer)

	if err != nil {
		log.Fatalf("port.Read: %v", err)
		panic("Failed on read")
	}
}

func (status *SystemStatus) InitializePort() {
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
	status.Port = port
}
