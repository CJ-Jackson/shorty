package http

import (
	"fmt"
	"github.com/CJ-Jackson/shorty/src/csrf"
	"github.com/CJ-Jackson/shorty/src/globals"
	"github.com/CJ-Jackson/shorty/src/router"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/options"
	"log"
	"net/http"
)

type httpCommand struct {
	address string
}

func (hC *httpCommand) CommandConfigure(c *cli.Command) {
	c.
		SetName("http:start").
		SetDescription("Start HTTP Server").
		AddOption("address", "address:port", options.String{
			Ptr: &hC.address,
		})
}

func (hC *httpCommand) CommandExecute() {
	router.SetUpShortyFileServers()
	csrf.InitShortyCsrf()

	fmt.Printf("Starting up HTTP Server at '%s'", hC.address)
	fmt.Println()
	fmt.Println()

	log.Fatal(http.ListenAndServe(hC.address, httpBoot{
		debug: !(globals.GetShortyGlobals().Production),
	}))
}

func init() {
	cli.RegisterCommand(&httpCommand{address: ":8080"})
}
