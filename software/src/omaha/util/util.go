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

func GetOmahaPath() string {
	return os.Getenv("SOFTWARE_DIR") + "/src/omaha"
}

/*func DecimalToHex(decimalValue int){		// Saving this for later
	b := [2]byte{}			// You can only go down to a 16-bit size so it split [8-bits][8-bits] 
	n := binary.LittleEndian.PutUint16(b[:], uint16(size))
	fmt.Printf("%d", b[1])		// We only need to send the first one, second will always be blank
}*/
