package main

import (
	"github.com/bestreyer/carfinder/cmd"
	"github.com/bestreyer/carfinder/context"
	"github.com/bestreyer/carfinder/env"
	"github.com/bestreyer/carfinder/server"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/mitchellh/cli"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"os"
)

func main() {
	err := env.LoadEnvVariables()

	if err != nil {
		log.Fatal(err)
	}

	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}

	en := en.New()

	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	rf := cmd.RegisterFactory{
		ServerFactory: &server.HTTPServerFactory{
			&context.ContextFactory{
				validator.New(),
				trans,
			},
		},
	}

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
