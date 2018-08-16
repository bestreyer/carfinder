package main

import (
	"github.com/bestreyer/carfinder/cmd"
	"github.com/bestreyer/carfinder/pkg/context"
	"github.com/bestreyer/carfinder/pkg/env"
	"github.com/bestreyer/carfinder/pkg/http"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/mitchellh/cli"
	"gopkg.in/go-playground/validator.v9"
	"log"
	"os"
	"fmt"
	"database/sql"
)

func main() {
	loadEnvVariables()
	r := createRegister()
	//dbConn := createDbConnFromEnv()
	ui := &cli.BasicUi{Writer: os.Stdout, ErrorWriter: os.Stderr}
	cli := &cli.CLI{
		Args:         os.Args[1:],
		Commands:     r.Map(ui),
		Autocomplete: true,
		Name:         "carfinder",
	}

	cli.Run()
}

func loadEnvVariables() {
	err := env.LoadEnvVariables()

	if nil != err {
		log.Fatal(err)
	}
}

func createTranslator() (ut.Translator) {
	en := en.New()
	uni := ut.New(en, en)
	trans, found := uni.GetTranslator("en")

	if true != found {
		log.Fatal("Translator has not been found")
	}

	return trans
}

func createRegister() (cmd.RegisterInterface) {
	rf := cmd.RegisterFactory{
		ServerFactory: &http.HTTPServerFactory{
			&context.ContextFactory{
				validator.New(),
				createTranslator(),
			},
		},
	}

	r, err := rf.Create()

	if err != nil {
		log.Fatal(err.Error())
	}

	return r
}

func createDbConnFromEnv() (*sql.DB) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=verify-full",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DBNAME"),
	)

	db, err := sql.Open("postgres", connStr)

	if nil != err {
		log.Fatal(err)
	}

	return db
}
