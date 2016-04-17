package context

import "net/http"

type Context struct {
	m   map[string]interface{}
	fns []func()
}

func newContext() *Context {
	return &Context{
		m: map[string]interface{}{},
	}
}

func GetContext(w http.ResponseWriter) *Context {
	return w.(*responseWriter).c
}

func (c *Context) Get(name string) interface{} {
	return c.m[name]
}

func (c *Context) Set(name string, v interface{}) {
	c.m[name] = v
}

func (c *Context) Dealer(name string, gotIt func(v interface{}), dontHaveIt func() interface{}) {
	v := c.m[name]
	if nil != v && nil != gotIt {
		gotIt(v)
	} else if nil != dontHaveIt {
		c.Set(name, dontHaveIt())
	}

}

func (c *Context) AddFunc(fn func()) {
	if nil == fn {
		return
	}
	c.fns = append(c.fns, fn)
}

func (c *Context) ExecFuncs() {
	for _, fn := range c.fns {
		fn()
	}
	c.fns = nil
}
