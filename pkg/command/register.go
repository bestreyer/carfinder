package command

import (
	"github.com/bestreyer/carfinder/pkg/command/server_start"
	"github.com/mitchellh/cli"
	"github.com/bestreyer/carfinder/pkg/server"
)

type Factory func(cli.Ui) (cli.Command, error)

type Register interface {
	Register(name string, fn Factory)
	Map(ui cli.Ui) map[string]cli.CommandFactory
}

type register struct {
	registry map[string]Factory
}

func (r register) Register(name string, fn Factory) {
	r.registry[name] = fn
}

func (r register) Map(ui cli.Ui) map[string]cli.CommandFactory {
	m := make(map[string]cli.CommandFactory)

	for name, fn := range r.registry {
		m[name] = func() (cli.Command, error) {
			return fn(ui)
		}
	}

	return m
}

func NewRegister(sf server.Factory) (Register, error) {
	r := &register{registry: map[string]Factory{}}

	r.Register("server start", func(ui cli.Ui) (cli.Command, error) {
		hs, err := sf.Create()
		if nil != err {
			return nil, err
		}

		return server_start.New(ui, hs), nil
	})

	return r, nil
}
