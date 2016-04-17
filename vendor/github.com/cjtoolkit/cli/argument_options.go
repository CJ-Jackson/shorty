package cli

import (
	"unicode/utf8"
)

type argumentOptions struct {
	arguments []string
	options   *options
	help      bool
}

func newArgumentOptions(values []string) argumentOptions {
	_arguments := []string{}
	_options := map[string][]string{}
	_help := false

	for _, value := range values {
		switch {
		case utf8.RuneCountInString(value) < 2 || "--" != value[:2]:
			_arguments = append(_arguments, value)
		case "--help" == value:
			_help = true
		case utf8.RuneCountInString(value) > 2:
			key, newValue := splitKeyValue(value[2:])
			_options[key] = append(_options[key], newValue)
		}
	}

	return argumentOptions{
		arguments: _arguments,
		options:   newOption(_options),
		help:      _help,
	}
}
