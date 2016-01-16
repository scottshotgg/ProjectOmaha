package main

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

func main() {

	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600, ReadTimeout: time.Second} // Probably want a half a second. This is plenty for the microcontroller to have time to respond
	s, err := serial.OpenPort(c)                                                    // If it hasn't responded by now then it isn't going to most likely
	if err != nil {                                                                 // Actually, you can't do this, for some reason it has to be an int value and cant be a floating value
		log.Fatal(err)
	}

	log.Println("Hello " + c.Name)

	_, err = s.Write([]byte{0x00})
	_, err = s.Write([]byte{0x01})
	_, err = s.Write([]byte{0x53})
	_, err = s.Write([]byte{0x60})
}