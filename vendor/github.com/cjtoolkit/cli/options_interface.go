package cli

type OptionsInterface interface {
	GetName() string
	HasOne() bool
	GetOne() string
	GetAll() []string
	Len() int
}
