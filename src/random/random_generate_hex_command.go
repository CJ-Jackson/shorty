package random

import (
	"fmt"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/arguments"
)

type randomGenerateHexCommand struct {
	numOfBytes int64
}

func (rGHC *randomGenerateHexCommand) CommandConfigure(c *cli.Command) {
	c.
		SetName("random:generate:hex").
		SetDescription("Generate Random Hexadecimal").
		AddArgument("numOfBytes", "Number Of Bytes", arguments.Int{
			Ptr: &rGHC.numOfBytes,
			Min: 1,
		})
}

func (rGHC *randomGenerateHexCommand) CommandExecute() {
	fmt.Println((Random{}).GenerateHex(int(rGHC.numOfBytes)))
}

func init() {
	cli.RegisterCommand(&randomGenerateHexCommand{})
}
