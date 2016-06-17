package util

import (
	"github.com/jacobsa/go-serial/serial"
	"os"
)

var Options serial.OpenOptions = serial.OpenOptions{
	PortName:        "/dev/ttyUSB0",
	BaudRate:        9600,
	DataBits:        8,
	StopBits:        1,
	MinimumReadSize: 1,
}

/*
	GetOmahaPath returns the root of the Omaha application.
*/
func GetOmahaPath() string {
	return os.Getenv("SOFTWARE_DIR") + "/src/omaha"
}
