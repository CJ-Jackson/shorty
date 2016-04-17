package cli

type ArgumentTransformerInterface interface {
	PreCheck()
	Constraint() string
	ArgumentTransform(argument string)
}
