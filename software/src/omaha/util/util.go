package util

import (
	"os"
)

func GetOmahaPath() string {
	return os.Getenv("SOFTWARE_DIR") + "/src/omaha"
}
