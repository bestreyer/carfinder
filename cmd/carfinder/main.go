package main

import (
	"github.com/mitchellh/cli"
	"os"
	"github.com/bestreyer/carfinder/pkg/di"
)

func main() {
	di := di.New()
	r := di.GetCommandRegister()

	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	cli := &cli.CLI{
		Args:         os.Args[1:],
		Commands:     r.Map(ui),
		Autocomplete: true,
		Name:         "carfinder",
	}

	cli.Run()
}