package main

import (
	"fmt"
	_ "github.com/CJ-Jackson/shorty/src"
	"github.com/CJ-Jackson/shorty/src/mongo"
	"github.com/CJ-Jackson/shorty/src/parameters"
	_ "github.com/CJ-Jackson/shorty/web"
	"github.com/cjtoolkit/cli"
)

func main() {
	parameters.InitShortyParameters()
	mongo.InitShortyMongoDb()

	fmt.Println("Welcome to Shorty, version 1.0")
	fmt.Println()

	cli.Run()
}
