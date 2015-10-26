package system

import (
	"fmt"
	"github.com/tarm/serial"
	"log"
	"time"
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
			//status.Port.Write(req.Data)			// I dont get what this is doing here....?

			//n, err := s.Write([]byte("test"))		// req.Data goes here?
        	//if err != nil {
            //	log.Fatal(err)
        	//}
		}
		if req.OnWrite != nil {
			result := req.OnWrite()					// I'm a bit confused on where the actual data is sent
			if req.ResultChan != nil {
				req.ResultChan <- result
			}
		}
	}
}

func getMessageHeader(section, id, size int8) []byte {
	header := make([]byte, size)
	header[0] = byte(section) // section ID
	header[1] = byte(id)      // speaker ID
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
/*	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 1,
	}*/

	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600, ReadTimeout: time.Second}	// Probably want a half a second. This is plenty for the microcontroller to have time to respond
	s, err := serial.OpenPort(c)															// If it hasn't responded by now then it isn't going to most likely
    if err != nil {																			// Actually, you can't do this, for some reason it has to be an int value and cant be a floating value
            log.Fatal(err)
    }

	fmt.Println("Hello " + c.Name)

	status.Port = s
}
