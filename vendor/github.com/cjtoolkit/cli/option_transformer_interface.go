package cli

type OptionTransformerInterface interface {
	PreCheck()
	Constraint() string
	OptionTransform(option OptionsInterface)
}
