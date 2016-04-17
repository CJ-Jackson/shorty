package globals

import (
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/options"
)

type globalsGlobal struct{}

func (_ globalsGlobal) GlobalConfigure(g *cli.Global) {
	g.
		AddOption("prod", "Set to Production Mode", options.Bool{Ptr: &shortyGlobals.Production})
}

func init() {
	cli.RegisterGlobal(globalsGlobal{})
}
