// +build !cgo

package parameters

import "os"

func getHomePath() string {
	return os.Getenv("HOME")
}
