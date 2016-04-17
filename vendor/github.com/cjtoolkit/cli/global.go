package cli

import (
	"github.com/cjtoolkit/cli/help"
	"strings"
)

type Global struct {
	global  GlobalInterface
	options []*option
}

func newGlobal(global GlobalInterface) *Global {
	return &Global{
		global:  global,
		options: []*option{},
	}
}

func (g *Global) AddOption(name, description string, transformer OptionTransformerInterface) *Global {
	g.options = append(g.options, &option{
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		Transformer: transformer,
	})
	return g
}

func (g *Global) postCheck() {
	g.checkAllOptions()
}

func (g *Global) checkAllOptions() {
	defer handleErrorAndPanicAgain("Global Options: ")
	for _, option := range g.options {
		option.postCheck()
	}
}

func (g *Global) populateOptions(op *options) {
	for _, option := range g.options {
		option.populate(op)
	}
}

func (g *Global) collectGeneralHelp(helpData *help.GeneralHelp) {
	for _, op := range g.options {
		helpData.GlobalsOptions = append(helpData.GlobalsOptions, help.Option{
			Name:        op.Name,
			Description: op.Description,
			Constraint:  op.Constraint,
		})
	}
}
