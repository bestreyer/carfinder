package main

import (
	"github.com/mitchellh/cli"
	"os"
	"github.com/bestreyer/carfinder/pkg/di"
	"github.com/bestreyer/carfinder/pkg/command"
)

func main() {
	di := di.New()

	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}
	c := command.NewCommands(ui, di.GetServerFactory(), di.GetLocationRepository())

	cli := &cli.CLI{
		Args:         os.Args[1:],
		Commands:     c,
		Autocomplete: true,
		Name:         "carfinder",
	}

	cli.Run()
}