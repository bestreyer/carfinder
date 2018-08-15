package command

import (
	"github.com/mitchellh/cli"
	"github.com/bestreyer/carfinder/command/server"
	httpserver "github.com/bestreyer/carfinder/server"
	"errors"
)

type Factory func(cli.Ui) (cli.Command, error)

type RegisterInterface interface {
	Register(name string, fn Factory)
	Map(ui cli.Ui) *cli.CommandFactory
}

type RegisterFactoryInterface interface {
	Create() (*Register, error)
}

type Register struct {
	registry map[string]Factory;
}

func (r Register) Register(name string, fn Factory) {
	r.registry[name] = fn;
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
	ServerFactory *httpserver.HTTPServerFactory
}

func (rf RegisterFactory) Create() (*Register, error) {
	if nil == rf.ServerFactory {
		return nil, errors.New("serverFactory should not be nil.")
	}

	r := &Register{registry: make(map[string]Factory)};
	r.Register("server start", func(ui cli.Ui) (cli.Command, error) {
		hs, err := rf.ServerFactory.Create()
		if nil != err {
			return nil, err
		}

		return server.New(ui, hs), nil
	})

	return r, nil
}
