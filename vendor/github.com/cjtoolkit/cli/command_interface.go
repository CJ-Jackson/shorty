package cli

type CommandInterface interface {
	CommandConfigure(c *Command)
	CommandExecute()
}
