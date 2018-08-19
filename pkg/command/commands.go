package command

import (
	"github.com/mitchellh/cli"
	"github.com/bestreyer/carfinder/pkg/server"
	"github.com/bestreyer/carfinder/pkg/command/serverstart"
	"github.com/bestreyer/carfinder/pkg/command/drivergenerate"
	"github.com/bestreyer/carfinder/pkg/location"
)


func NewCommands(ui cli.Ui, sf server.Factory, lr location.Repository) (map[string]cli.CommandFactory) {
	return map[string]cli.CommandFactory {
		"server start": func() (cli.Command, error) {
			hs, err := sf.Create()
			if nil != err {
				return nil, err
			}

			return serverstart.New(ui, hs), nil
		},

		"driver generate": func() (cli.Command, error) {
			return drivergenerate.New(ui, lr), nil
		},
	}
}
