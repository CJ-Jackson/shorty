package cli

type option struct {
	Name        string
	Constraint  string
	Description string
	Transformer OptionTransformerInterface
}

func (o *option) postCheck() {
	switch {
	case "" == o.Name:
		panic("'name' cannot be left blank")
	case !otherNamePattern.MatchString(o.Name):
		panic("'name' does not match `" + otherNamePattern.String() + "`")
	case nil == o.Transformer:
		panic(o.Name + ": 'transformer' cannot be nil")
	default:
		defer handleErrorAndPanicAgain(o.Name + ": Transformer: ")
		o.Transformer.PreCheck()
		o.Constraint = o.Transformer.Constraint()
	}
}

func (o *option) populate(op *options) {
	defer handleErrorAndPanicAgain("Option: " + o.Name + ": ")
	op.setName(o.Name)

	execTrueFalse(op.HasOne(), func() {
		o.Transformer.OptionTransform(op)
	}, blankFn)
}
