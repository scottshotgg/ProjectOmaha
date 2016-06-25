package util

import "os"

/*
	GetOmahaPath returns the root of the Omaha application.
*/
func GetOmahaPath() string {
	return os.Getenv("SOFTWARE_DIR") + "/software/src/omaha"
}
