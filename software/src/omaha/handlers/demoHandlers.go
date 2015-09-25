package handlers

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"net/http"
)

func DemoWriteHandler(w http.ResponseWriter, r *http.Request) {
	options := serial.OpenOptions{
		PortName:        "/dev/ttyUSB0",
		BaudRate:        9600,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	// var id = -1

	fmt.Println("Hello " + options.PortName)

	// Open the port.
	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}

	// Make sure to close it later.
	//defer port.Close()

	// Write 4 bytes to the port.
	var a = "a"
	start := []byte{0x01} //, 0x01, 0x02, 0x03}
	//stop := []byte(0x00)
	//i, err := strconv.Atoi("a")
	n, err := port.Write(start)
	if err != nil {
		log.Fatalf("port.Write: %v", err)
	}
	//write("a", port)

	// := []byte("k")

	var _ = n
	fmt.Println("Wrote: ", a)

	// n, err = port.Read(b)

	// //var bSize = len(b)

	// s := string(b[:len(b)])
	// fmt.Println("Read:  ", s)
	fmt.Fprint(w, "Started")
	defer port.Close()
}
