package parameters

import (
	"encoding/json"
	"github.com/CJ-Jackson/shorty/src/common"
	"os"
)

func InitShortyParameters() {
	file, err := os.Open(getHomePath() + SETTING_FILE_LOCATION[1:])
	common.CheckError(err)

	err = json.NewDecoder(file).Decode(shortyParameters)
	file.Close()
	common.CheckError(err)
}
