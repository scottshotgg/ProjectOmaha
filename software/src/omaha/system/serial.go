package system

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
)

var MessageChan chan *ControllerRequest = make(chan *ControllerRequest, 100)

type ControllerRequest struct {
	Data       []byte
	OnWrite    func() interface{}
	ResultChan chan interface{}
}

func HandleControllerMessages() {
	for {
		req := <-MessageChan
		status := GetSystemStatus()
		if !status.IsDebug() {
			status.Port.Write(req.Data)
		}
		if req.OnWrite != nil {
			req.ResultChan <- req.OnWrite()
		}
	}
}

func (status *SystemStatus) GetMessageHeader(size int) []byte {
	header := make([]byte, size)
	header[0] = 0x00 // section ID
	header[1] = 0x6B // speaker ID
	return header
}

func (status *SystemStatus) SendData(data byte) {
	fmt.Printf("port.Write: %v\n", data)
	_, err := status.Port.Write([]byte{data})
	if err != nil {
		log.Fatalf("port.Write: %v", err)
		panic("Failed on write")
	}
}

/*func (status *SystemStatus) SendMoreData(data []byte, amount int) {
	for int i := 0; i < amount; i++ {
		fmt.Printf("port.Write: %v\n", data)
		_, err := status.Port.Write(data)
		if err != nil {
			log.Fatalf("port.Write: %v", err)
			panic("Failed on write")
		}
	}
}*/

func (status *SystemStatus) ReadData(buffer []byte) bool {
	_, err := status.Port.Read(buffer)
	fmt.Printf("port.Read: %v\n", buffer)
	if err != nil {
		log.Fatalf("port.Read: %v", err)
		panic("Failed on read")

		return false
	}

	return true
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
