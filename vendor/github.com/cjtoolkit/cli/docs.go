/*
Google Go port of Symfony Console Component (http://symfony.com/doc/current/components/console/introduction.html).

Usage

As a starting point

	cli.Run()

In the main function.  To setup a command place something like this somewhere else in the source code of your project.

	type greetCommand struct {
		text string
		yell bool
	}

	func (gC *greetCommand) CommandConfigure(c *cli.Command) {
		c.SetName("demo:greet").
			SetDescription("Greet someone").
			AddOption("name", "Who do you want to greet?",
				options.String{
					Ptr: &gC.text,
					MaxRune: 50,
				}).
			AddOption("yell", "If set, the task will yell in uppercase letters",
				options.Bool{Ptr: &gC.yell})
	}

	func (gC *greetCommand) CommandExecute() {
		if "" == gC.text {
			gC.text = "Hello!"
		} else {
			gC.text = "Hello " + gC.text + "!"
		}

		if gC.yell {
			gC.text = strings.ToUpper(gC.text)
		}

		fmt.Println(gC.text)
	}

And finally register the command. (Make sure you import the code.)

	func init() {
		cli.RegisterCommand(&greetCommand{})
	}

If you need to setup global options, add something like this somewhere else in the source code of your project.

	type Global struct {
		Prod bool
	}

	var (
		global     = &Global{}
		globalSync sync.RWMutex
	)

	func GetGlobal() Global {
		globalSync.RLock()
		defer globalSync.RUnlock()

		return *global
	}

	type globalGlobal struct{}

	func (_ globalGlobal) GlobalConfigure(g *cli.Global) {
		g.AddOption("prod", "Set to Production Mode", options.Bool{Ptr: &global.Prod})
	}

	func (_ globalGlobal) Lock() {
		globalSync.Lock()
	}

	func (_ globalGlobal) Unlock() {
		globalSync.Unlock()
	}

And finally register global option (can only be called once.)

	func init() {
		cli.RegisterGlobal(globalGlobal{})
	}

Note

* Only support long option (with '=' sign, if using non-bool value, no need to use '=' with bool), there is no supports
for short option.  (No spaces, just only '=')

* All arguments are mandatory, while all options are optional.  There is no support for optional arguments or mandatory
options.  There is no mandatory flag like there is with Symfony Console Component.

* Commands and Globals both work like Object Oriented compartment.

* For extra flexibility Argument and Option use those interfaces retrospectively, in favour of those constant variables
you find in Symfony Console Component
	ArgumentTransformerInterface
	OptionTransformerInterface
*/
package cli
