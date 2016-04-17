package cli

type commandArgument struct {
	Name        string
	Constraint  string
	Description string
	Transformer ArgumentTransformerInterface
}

func (cA *commandArgument) postCheck() {
	switch {
	case "" == cA.Name:
		panic("'name' cannot be left blank")
	case !otherNamePattern.MatchString(cA.Name):
		panic("'name' does not match `" + otherNamePattern.String() + "`")
	case nil == cA.Transformer:
		panic(cA.Name + ": 'transformer' cannot be nil")
	default:
		defer handleErrorAndPanicAgain(cA.Name + ": Transformer: ")
		cA.Transformer.PreCheck()
		cA.Constraint = cA.Transformer.Constraint()
	}
}

func (cA *commandArgument) populate(argument string) {
	defer handleErrorAndPanicAgain("Argument: " + cA.Name + ": ")
	cA.Transformer.ArgumentTransform(argument)
}
