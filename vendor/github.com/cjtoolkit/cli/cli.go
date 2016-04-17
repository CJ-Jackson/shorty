package cli

import (
	"github.com/cjtoolkit/cli/help"
	"io"
	"os"
)

type cli struct {
	args     argumentOptions
	global   *Global
	commands map[string]*Command
	stdErr   io.Writer
	stdOut   io.Writer
	exitFn   func()
}

func newCli() *cli {
	return &cli{
		args:     newArgumentOptions(os.Args[1:]),
		commands: map[string]*Command{},
		stdErr:   os.Stderr,
		stdOut:   os.Stdout,
		exitFn: func() {
			os.Stderr.Write([]byte(NEW_LINE))
			os.Exit(1)
		},
	}
}

func (c *cli) registerGlobal(global GlobalInterface) {
	defer handleErrorAndPanicAgain("Register Global: ")

	execFunctionIfNotNil(c.global, func() {
		panic("cannot be called more than once")
	})

	execFunctionIfNil(global, func() {
		panic("global cannot be nil")
	})

	_global := newGlobal(global)

	global.GlobalConfigure(_global)

	_global.postCheck()

	c.global = _global
}

func (c *cli) registerCommand(command CommandInterface) {
	handleErrorAndPanicAgain("Register Command: ")

	execFunctionIfNil(command, func() {
		panic("command cannot be nil")
	})

	_command := newCommand(command)

	command.CommandConfigure(_command)

	_command.postCheck()

	execFunctionIfNotNil(c.commands[_command.name], func() {
		panic("'" + _command.name + "' has already been taken")
	})

	c.commands[_command.name] = _command
}

func (c *cli) run() {
	execTrueFalse(c.args.help || 0 == len(c.args.arguments), func() {
		c.help()
	}, func() {
		c.execute()
	})
}

func (c *cli) help() {
	defer handleErrorAndExit("Help: ", c.exitFn, c.stdErr)
	execTrueFalse(len(c.args.arguments) >= 1, func() {
		c.helpWithCommand()
	}, func() {
		c.helpInGeneral()
	})
}

func (c *cli) helpInGeneral() {
	helpData := help.NewGeneralHelp()

	execFunctionIfNotNil(c.global, func() {
		c.global.collectGeneralHelp(helpData)
	})

	for name, command := range c.commands {
		commandData := &help.Command{
			Name:        name,
			Description: command.description,
		}

		execTrueFalse(nil == helpData.Commands, func() {
			helpData.Commands = help.Commands{
				name: commandData,
			}
		}, func() {
			helpData.Commands[name] = commandData
		})
	}

	helpData.Finalise()
	helpData.Render(c.stdOut)
}

func (c *cli) helpWithCommand() {
	defer handleErrorAndPanicAgain("Command: ")

	command := c.commands[c.args.arguments[0]]

	execFunctionIfNil(command, func() {
		c.helpInGeneral()
		panic("'" + c.args.arguments[0] + "' was not found.")
	})

	helpData := help.NewCommandHelp(command.name)
	command.collectGeneralHelp(helpData)

	helpData.Finalise()
	helpData.Render(c.stdOut)
}

func (c *cli) execute() {
	defer handleErrorAndExit("Execute Command: ", c.exitFn, c.stdErr)
	execTrueFalse(len(c.args.arguments) >= 1, func() {
		c.executeCommand()
	}, func() {
		c.helpInGeneral()
		panic("name of command was not specified")
	})
}

func (c *cli) executeCommand() {
	command := c.commands[c.args.arguments[0]]

	execFunctionIfNil(command, func() {
		c.helpInGeneral()
		panic("'" + c.args.arguments[0] + "' was not found.")
	})

	c.commands = nil
	execFunctionIfNotNil(c.global, func() {
		c.global.populateOptions(c.args.options)
	})

	c.global = nil

	command.execCommand(c.args)
}
