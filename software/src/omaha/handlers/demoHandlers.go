package handlers

import (
	"fmt"
	"github.com/jacobsa/go-serial/serial"
	"log"
	"net/http"
)

func DemoStartHandler(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprint(w, "An error occurred on write")
		return
	}
	//write("a", port)

	// := []byte("k")
	fmt.Println("Wrote: ", b)

	_, err = port.Read(b)

	if err != nil {
		log.Fatalf("port.Read: %v", err)
		fmt.Fprint(w, "An error occurred on read")
		return
	}

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)
	fmt.Fprint(w, s)
	defer port.Close()
}

func DemoStopHandler(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprint(w, "An error occurred on write")
		return
	}
	//write("a", port)

	// := []byte("k")

	fmt.Println("Wrote: ", b)

	_, err = port.Read(b)

	if err != nil {
		log.Fatalf("port.Read: %v", err)
		fmt.Fprint(w, "An error occurred on read")
		return
	}

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)
	fmt.Fprint(w, s)
	defer port.Close()
}

func DemoLEDHandler(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprint(w, "An error occurred on write")
		return
	}
	//write("a", port)

	// := []byte("k")

	fmt.Println("Wrote: ", b)

	_, err = port.Read(b)

	if err != nil {
		log.Fatalf("port.Read: %v", err)
		fmt.Fprint(w, "An error occurred on read")
		return
	}

	s := string(b[:len(b)])
	fmt.Println("Read:  ", b)
	fmt.Fprint(w, s)
	defer port.Close()
}
