package cli

import "regexp"

var (
	main_cli = newCli()
	blankFn  = func() {}

	commandNamePattern = regexp.MustCompile(`^[^\s:-]{1}[^\s:]+[:]{1}[^\s]+$`)
	otherNamePattern   = regexp.MustCompile(`^[^\s-=]{1}[^\s=]*$`)
)
