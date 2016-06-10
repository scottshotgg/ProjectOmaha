package system

import (
	"log"
	"omaha/database"
)

const NULLZone int = 0

func KeepAlive(ID int8) (error int8) {
	data := getMessageHeader(ID, 3)
	data[1] = Commands.KeepAlive
	data[2] = 0x01

    defer func() {
    	if r := recover(); r != nil {
        	log.Println("Error reading from the port for controller ", ID, "\n", r)
    		error = -ID
    	}
    }()

	log.Println("KeepAlive: ", data)
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			return nil
		}
		return nil
	}}
	MessageChan <- req

	if(!status.IsDebug()){
		b := []byte{0x00}			
		status.ReadData(b)

		log.Println("Return packet: ", b[0])

		if(int8(b[0]) == int8('A')) {
			log.Println(int8('A'))
			error = 0
		} else {
			error = int8(b[0])
		}
	}

	return 
}

func IsLEDOn(this *database.ControllerStatus) bool {
	return this.LEDOn
}

func TurnLEDOn(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.TurnLEDOn
	data[2] = 0x01

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		this.LEDOn = true
		if status.IsDebug() {
			log.Println("LED turned on")
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func TurnLEDOff(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.TurnLEDOff
	data[2] = 0x00

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		this.LEDOn = false
		if status.IsDebug() {
			log.Println("LED turned off")
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

type LEDStatusResponse struct {
	ledOn bool
}

func GetLEDStatusFromController(this *database.ControllerStatus) (bool, error) {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.GetLEDStatus
	data[2] = 0x00

	ch := make(chan interface{})
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		response := &LEDStatusResponse{}
		if status.IsDebug() {
			log.Println("Got LED Status From Controller")
			response.ledOn = true
		} else {
			b := make([]byte, 1)
			status.ReadData(b)
			response.ledOn = b[0] != 0x01
		}
		return response
	}, ResultChan: ch}

	MessageChan <- req
	response := (<-ch).(*LEDStatusResponse)

	return response.ledOn, nil
}

func SetVolume(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetVolume
	data[2] = byte(this.VolumeLevel[0])

	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Printf("Set volume to %d\n", this.VolumeLevel[0])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func GetVolumeFromController(this *database.ControllerStatus) (int8, error) {
	if status.debug {
		return 0, nil
	}

	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.GetVolume
	data[2] = 0x00

	return 0, nil
}

func SetMusicVolume(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetMusicVolume
	data[2] = byte(this.VolumeLevel[1])

	log.Println("Packet contents: ", data)

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Printf("Set music volume to %d\n", this.VolumeLevel[1])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func SetPaging(this *database.ControllerStatus, pagingData int8) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetPaging
	data[2] = byte(pagingData)

	log.Println("Packet contents: ", data)

	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {

		if status.IsDebug() {
			log.Printf("Set paging data to %d\n", pagingData)
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func SetSoundMaskingVolume(this *database.ControllerStatus) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetSoundMaskingVolume
	data[2] = byte(this.VolumeLevel[3])

	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Printf("Set sound masking volume to %d\n", this.VolumeLevel[3])
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func SetAveragingMode(this *database.ControllerStatus, mode int8) error {
	data := getMessageHeader(this.ID, 3)
	data[1] = Commands.SetAveragingFilter
	data[2] = byte(mode)

	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Printf("Set averaging mode to %d\n", mode)
		}
		return nil
	}}
	MessageChan <- req

	return nil
}

func (status *SystemStatus) ResetFIFO() (int, error) { 
	if status.debug {
		return 0, nil
	}

	return 0, nil
}

func (status *SystemStatus) ResetMicrocontroller() (int, error) {
	if status.debug {
		return 0, nil
	}

	data := getMessageHeader(0, 4)
	data[1] = 0x00
	data[2] = 0x00

	return 0, nil

}

func GetAveragingMode(this *database.ControllerStatus) (int, error) {
	if status.debug {
		return 0, nil
	}

	data := getMessageHeader(0, 4)
	data[1] = 0x61
	data[2] = 0x00

	b := []byte{0x00}
	status.ReadData(b)

	return 0, nil

}

func (status *SystemStatus) SendCoefficientInformation(equalizedGain int8, decibal int8) (int8, error) {
	if status.debug {
		return 0, nil
	}

	return 0, nil
}

func SetEqualizerConstant(this *database.ControllerStatus, level int8, band int8, decimal bool) (int, error) {
	data := getMessageHeader(this.ID, 3)
	data[1] = byte(band)
	if(decimal) {
		data[2] = byte(level)
	} else {
		data[2] = byte(level + 40)
	}
	log.Println("Packet contents: ", data)
	
	req := &ControllerRequest{ Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Printf("Set band %-2d mode to %3ddB (-40)   |   Packet contents: %2d", band, level, data)
		}
		return nil
	}}
	MessageChan <- req	

	return 0, nil
}

func SetEQMode(ID int8, mode int8) (int, error) {
	data := getMessageHeader(ID, 3)
	data[1] = Commands.SetEQMode
	data[2] = byte(mode)

	req := &ControllerRequest{ Data: data, OnWrite: func() interface{} {
		if status.IsDebug() {
			log.Println("Set EQ mode to ",  mode)
		}
		return nil
	}}
	MessageChan <- req

	return 0, nil
}

func SendPagingRequest(ID int8) (error) {
	if status.debug {
		return nil
	}

	data := getMessageHeader(ID, 3)
	data[1] = Commands.SetPaging
	data[2] = 0

	return nil
}

func (status *SystemStatus) AreYouAlive(n map[int]string) (m map[int]string) {
	if status.debug {
		return m
	}

	m = make(map[int]string)

	for i := 1; i <= len(n); i++ {
		b := []byte{0x00}
		alive := status.ReadData(b)

		if alive != true || b[0] != 0x72 {
		}
	}

	return m
}