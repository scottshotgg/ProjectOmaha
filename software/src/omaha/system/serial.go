package system

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
)

func (status *SystemStatus) SendMessageHeader(size int) {
	status.SendData([]byte{0x6B})
	
	b := [2]byte{}			// You can only go down to a 16-bit size so it split [8-bits][8-bits] 
	n := binary.LittleEndian.PutUint16(b[:], uint16(size))
	fmt.Printf("%d", b[1])		// We only need to send the first one, second will always be blank
	status.SendData(b[1])		// This needs to be changed
}

func (status *SystemStatus) SendData(data []byte) {
	fmt.Printf("port.Write: %v\n", data)
	_, err := status.Port.Write(data)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
		panic("Failed on write")
	}
}

func (status *SystemStatus) ReadData(buffer []byte) {
	_, err := status.Port.Read(buffer)
	fmt.Printf("port.Read: %v\n", buffer)
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
