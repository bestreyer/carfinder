package di

import (
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/locales/en"
	"github.com/bestreyer/carfinder/pkg/command"
	"github.com/bestreyer/carfinder/pkg/context"
	"gopkg.in/go-playground/validator.v9"
	"github.com/bestreyer/carfinder/pkg/server"
	"database/sql"
	"fmt"
	"github.com/bestreyer/carfinder/pkg/env"
	"github.com/labstack/gommon/log"
)

type DI interface {
	GetTranslator() (ut.Translator)
	GetCommandRegister() (command.Register)
	GetDbConn() (*sql.DB)
	GetValidator() (*validator.Validate)
}

type di struct {
	translator ut.Translator
	register   command.Register
	dbConn     *sql.DB
	validator  *validator.Validate
}

func (d *di) GetValidator() (*validator.Validate) {
	if nil != d.validator {
		return d.validator
	}

	d.validator = validator.New()
	return d.validator
}

func (d *di) GetTranslator() (ut.Translator) {
	if nil != d.translator {
		return d.translator
	}

	en := en.New()
	uni := ut.New(en, en)
	trans, found := uni.GetTranslator("en")

	if true != found {
		log.Fatal("Translator has not been found")
	}

	d.translator = trans
	return d.translator
}

func (d *di) GetCommandRegister() (command.Register) {
	if nil != d.register {
		return d.register
	}

	cf := context.NewFactory(d.GetValidator(), d.GetTranslator())
	sf := server.NewFactory(cf)
	r, err := command.NewRegister(sf)

	if err != nil {
		log.Fatal(err.Error())
	}

	d.register = r

	return d.register
}

func (d *di) GetDbConn() (*sql.DB) {
	if nil != d.dbConn {
		return d.dbConn
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=verify-full",
		env.GetEnv("POSTGRES_USER", "root"),
		env.GetEnv("POSTGRES_PASS", "root"),
		env.GetEnv("POSTGRES_HOST", "127.0.0.1"),
		env.GetEnv("POSTGRES_DBNAME", "carfinder"),
	)

	db, err := sql.Open("postgres", connStr)

	if nil != err {
		log.Fatal(err)
	}

	d.dbConn = db
	return d.dbConn
}

func New() (DI) {
	return &di{}
}