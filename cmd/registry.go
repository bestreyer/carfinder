package cmd

import (
	"github.com/mitchellh/cli"
	"github.com/bestreyer/carfinder/cmd/server_start"
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
	r.Register("server_start start", func(ui cli.Ui) (cli.Command, error) {
		hs, err := rf.ServerFactory.CreateHTTPServer()
		if nil != err {
			return nil, err
		}

		return server_start.New(ui, hs), nil
	})

	return r, nil
}
