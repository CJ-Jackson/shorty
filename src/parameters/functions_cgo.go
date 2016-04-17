// +build cgo

package parameters

import (
	"github.com/CJ-Jackson/shorty/src/common"
	"os/user"
)

func getHomePath() string {
	usr, err := user.Current()
	common.CheckError(err)

	return usr.HomeDir
}
