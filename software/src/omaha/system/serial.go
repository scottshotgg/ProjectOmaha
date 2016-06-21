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

/*
	The startup process looks for speakers on the chains.
*/
func StartUpProcess(wg *sync.WaitGroup) {
	status := GetSystemStatus()
  done := make(chan bool)
  // set finding = true because we don't want the system to start until we have signaled that we have left the initial finding stage
	status.SetFinding(true)
  go func() {
      var x int8 = 0
      chain := 0
      // for every second, send out a keepAlive and see what you get
      for {
          second := time.After(1 * time.Second)
          select {
              case <- second:
                  found := KeepAlive(x)
                  
                  /*******
                  This is for debugging purposes only
                  vv
                  ********/

                  if(x == 5) {
                      wg.Done()
                      done <- true
                  }

                  /*******
                  ^^
                  This is for debugging purposes only
                  ********/

                  // if we got something back, append it to its respective chain
                  if(found != 0) {
                      log.Println("Found a speaker")
                      chains = append(chains, x - 1)
                      x++
                      if(chain == 4) {
                          wg.Done()
                          // when the chain is equal to 4 and we have already 
													status.SetFinding(false)
                          done <- true
                      } else {
                          chain++
                          SwitchTransceiver(chain)  
                          x = 0
                      }
                  } else {
                  	
                  }

              case <- done:
                  return
          }
      }
  }() 
}

/*
	This is essentially a getChainByID function that serves to resolve which chain we need to send on to contact the speaker that we want to talk to.
*/
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

/*
	HandleControllerMessages is called from main.go, this handlers the sending of the controller messages using a channel.
*/
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

/*
	HandleControllerMessages uses this to send data to the current port. This function SHOULD not use this function raw. Insert something into the channel and the handler will take care of it.
*/
func (status *SystemStatus) SendData(data byte) {
	log.Printf("port.Write: %v\n", data)
	_, err := status.Port.Write([]byte{data})
	if err != nil {
		log.Fatalf("port.Write: %v", err)
		panic("Failed on write")
	}
}

/*
	This is used to read data from the current port. This function may be used raw, but it needs to be wrapped in a timeout because it is a blocking function.
*/
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

/*
	This function sets up the ports
*/
func (status *SystemStatus) InitializePort() {
	// osx: /dev/cu.uart-34FF466E37414555
	portMap = append(portMap, &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600, ReadTimeout: time.Second})
	portMap = append(portMap, &serial.Config{Name: "/dev/ttyUSB1", Baud: 9600, ReadTimeout: time.Second})
	portMap = append(portMap, &serial.Config{Name: "/dev/ttyUSB2", Baud: 9600, ReadTimeout: time.Second})
	portMap = append(portMap, &serial.Config{Name: "/dev/ttyUSB3", Baud: 9600, ReadTimeout: time.Second})

	SwitchTransceiver(0)
}

/*
	This function switches the transiever based on the number sent to it. The resolver uses this function to change the chain that the handler is talking to.
*/
func SwitchTransceiver(id int) {
	status := GetSystemStatus()

	s, err := serial.OpenPort(portMap[id])                                                    
	if err != nil {                                                               
		log.Println(err)
	}

	log.Println("Hello " + portMap[id].Name)

	status.Port = s	
}
