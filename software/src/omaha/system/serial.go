package system

import (
	"github.com/tarm/serial"
	"io"
	"log"
	"time"
  "sync"
)

var MessageChan chan *ControllerRequest = make(chan *ControllerRequest, 100)

var portMap [] *serial.Config
var chains[] int8
type ControllerRequest struct {
	Data       []byte
	OnWrite    func() interface{}
	ResultChan chan interface{}
}


func StartUpProcess(wg *sync.WaitGroup) {
	status := GetSystemStatus()
  done := make(chan bool)
	status.SetFinding(true)
  go func() {
      var x int8 = 0
      chain := 0
      for {
          second := time.After(1 * time.Second)
          select {
              case <- second:
                  found := KeepAlive(x)

                  if(x == 5) {
                      wg.Done()
                      done <- true
                  }

                  if(found != 0) {
                      chains = append(chains, x - 1)
                      if(chain == 3) {
                          wg.Done()
													status.SetFinding(false)
                          done <- true
                      } else {
                          chain++
                          SwitchTransceiver(chain)  
                          x = 0
                      }
                  } else {
                      log.Println("Found a speaker")
                      x++
                  }

              case <- done:
                  return
          }
      }
  }() 
}

func resolveChain(id int8) {
	transceiver := 0
	switch {
		case id > chains[0]:
			transceiver = 0
			break
		case id > chains[1]:
			transceiver = 1
			break
		case id > chains[2]:
			transceiver = 2
			break
		case id > chains[3]:
			transceiver = 3
			break
		default:
			log.Println("SPEAKER WITH ID:", id, " COULD NOT BE RESOLVED")
	}

	SwitchTransceiver(transceiver)

}

func HandleControllerMessages() {
	for {
		req := <-MessageChan
		status := GetSystemStatus()

		if !status.IsDebug() {
			if !status.IsFinding() {
				resolveChain(int8(req.Data[0]))
			}
			_, err := status.Port.Write(req.Data)

			if err != nil {
				log.Fatal(err)
			}
		}
		if req.OnWrite != nil {
			result := req.OnWrite()
			if req.ResultChan != nil {
				req.ResultChan <- result
			}
		}
	}
}

func getMessageHeader(id, size int8) []byte {
	header := make([]byte, size)
	header[0] = byte(id)

	return header
}

func (status *SystemStatus) SendData(data byte) {
	log.Printf("port.Write: %v\n", data)
	_, err := status.Port.Write([]byte{data})
	if err != nil {
		log.Fatalf("port.Write: %v", err)
		panic("Failed on write")
	}
}

func (status *SystemStatus) ReadData(buffer []byte) bool {

	buffer[0] = 0x01
	_, err := status.Port.Read(buffer)
	log.Printf("port.Read: %v\n", buffer)
	if err != nil && err != io.EOF {
		log.Fatalf("port.Read: %v", err)
		panic("Failed on read")

		return false
	}

	return true
}

func (status *SystemStatus) InitializePort() {
	// osx: /dev/cu.uart-34FF466E37414555
	portMap = append(portMap, &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600, ReadTimeout: time.Second})
	portMap = append(portMap, &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600, ReadTimeout: time.Second})
	portMap = append(portMap, &serial.Config{Name: "/dev/ttyUSB2", Baud: 9600, ReadTimeout: time.Second})
	portMap = append(portMap, &serial.Config{Name: "/dev/ttyUSB3", Baud: 9600, ReadTimeout: time.Second})

	SwitchTransceiver(0)
}

func SwitchTransceiver(id int) {
	status := GetSystemStatus()

	s, err := serial.OpenPort(portMap[id])                                                    
	if err != nil {                                                                 
		log.Fatal(err)
	}

	log.Println("Hello " + portMap[id].Name)

	status.Port = s	
}
