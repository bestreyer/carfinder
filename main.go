package main

import (
	"github.com/mitchellh/cli"
	"github.com/bestreyer/carfinder/command"
	"os"
	"github.com/bestreyer/carfinder/server"
	"log"
	"github.com/bestreyer/carfinder/env"
)

func main() {
	err := env.LoadEnvVariables()

	if err != nil {
		log.Fatal(err)
	}

	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	hsf:= &server.HTTPServerFactory{}
	rf := command.RegisterFactory{ServerFactory: hsf}

	r, err := rf.Create()
	if err != nil {
		log.Fatal(err.Error())
	}

	cli := &cli.CLI{
		Args:         os.Args[1:],
		Commands:     r.Map(ui),
		Autocomplete: true,
		Name:         "carfinder",
	}

	cli.Run()
}

