package cli

import (
	"fmt"
	"github.com/cjtoolkit/cli/help"
	"strings"
)

type Command struct {
	command     CommandInterface
	name        string
	description string
	options     []*option
	arguments   []*commandArgument
}

func newCommand(command CommandInterface) *Command {
	return &Command{
		command:   command,
		options:   []*option{},
		arguments: []*commandArgument{},
	}
}

func (c *Command) SetName(name string) *Command {
	c.name = name

	return c
}

func (c *Command) SetDescription(description string) *Command {
	c.description = description

	return c
}

func (c *Command) AddOption(name, description string, transformer OptionTransformerInterface) *Command {
	c.options = append(c.options, &option{
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		Transformer: transformer,
	})

	return c
}

func (c *Command) AddArgument(name, description string, transformer ArgumentTransformerInterface) *Command {
	c.arguments = append(c.arguments, &commandArgument{
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		Transformer: transformer,
	})

	return c
}

func (c *Command) postCheck() {
	switch {
	case "" == c.name:
		panic("Name cannot be left blank")
	case !commandNamePattern.MatchString(c.name):
		panic("Name does not match `" + commandNamePattern.String() + "`")
	default:
		c.checkAllOptions()
		c.checkAllArgument()
	}
}

func (c *Command) checkAllOptions() {
	defer handleErrorAndPanicAgain("Command Options: " + c.name + ": ")
	for _, option := range c.options {
		option.postCheck()
	}
}

func (c *Command) checkAllArgument() {
	defer handleErrorAndPanicAgain("Command Argument: " + c.name + ": ")
	for _, argument := range c.arguments {
		argument.postCheck()
	}
}

func (c *Command) execCommand(argOp argumentOptions) {
	c.populateOptions(argOp.options)
	argOp.options.checkForUnrecognisedOption()
	c.populateArgument(argOp.arguments[1:])
	c.command.CommandExecute()
}

func (c *Command) populateOptions(op *options) {
	for _, option := range c.options {
		option.populate(op)
	}
}

func (c *Command) populateArgument(arguments []string) {
	if userArgCount, argCount := len(arguments), len(c.arguments); userArgCount != argCount {
		panic(fmt.Sprintf("Expecting '%d' argument(s) got '%d'", argCount, userArgCount))
	}
	for key, argument := range c.arguments {
		argument.populate(arguments[key])
	}
}

func (c *Command) collectGeneralHelp(helpData *help.CommandHelp) {
	for _, op := range c.options {
		helpData.Options = append(helpData.Options, help.Option{
			Name:        op.Name,
			Description: op.Description,
			Constraint:  op.Constraint,
		})
	}
	for _, arg := range c.arguments {
		helpData.Arguments = append(helpData.Arguments, help.Argument{
			Name:        arg.Name,
			Description: arg.Description,
			Constraint:  arg.Constraint,
		})
	}
}
