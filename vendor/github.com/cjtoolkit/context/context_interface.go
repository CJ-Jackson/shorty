package context

type ContextInterface interface {
	Get(name string) interface{}
	Set(name string, v interface{})
	Dealer(name string, gotIt func(v interface{}), dontHaveIt func() interface{})
	AddFunc(fn func())
	ExecFuncs()
}
