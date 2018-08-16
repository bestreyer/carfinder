package command

import (
	"errors"
	"github.com/bestreyer/carfinder/pkg/command/server_start"
	httpserver "github.com/bestreyer/carfinder/pkg/http"
	"github.com/mitchellh/cli"
)

type Factory func(cli.Ui) (cli.Command, error)

type RegisterInterface interface {
	Register(name string, fn Factory)
	Map(ui cli.Ui) map[string]cli.CommandFactory
}

type RegisterFactoryInterface interface {
	Create() (RegisterInterface, error)
}

type Register struct {
	registry map[string]Factory
}

func (r Register) Register(name string, fn Factory) {
	r.registry[name] = fn
}

func (r Register) Map(ui cli.Ui) map[string]cli.CommandFactory {
	m := make(map[string]cli.CommandFactory)

	for name, fn := range r.registry {
		m[name] = func() (cli.Command, error) {
			return fn(ui)
		}
	}

	return m
}

type RegisterFactory struct {
	ServerFactory httpserver.HTTPServerFactoryInterface
}

func (rf RegisterFactory) Create() (RegisterInterface, error) {
	if nil == rf.ServerFactory {
		return nil, errors.New("ServerFactory should not be nil.")
	}

	r := &Register{registry: make(map[string]Factory)}
	r.Register("server start", func(ui cli.Ui) (cli.Command, error) {
		hs, err := rf.ServerFactory.CreateHTTPServer()
		if nil != err {
			return nil, err
		}

		return server_start.New(ui, hs), nil
	})

	return r, nil
}
