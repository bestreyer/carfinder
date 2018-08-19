package di

import (
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/locales/en"
	"github.com/bestreyer/carfinder/pkg/context"
	"gopkg.in/go-playground/validator.v9"
	"github.com/bestreyer/carfinder/pkg/server"
	"database/sql"
	"fmt"
	"github.com/bestreyer/carfinder/pkg/env"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/bestreyer/carfinder/pkg/location"
	"github.com/bestreyer/carfinder/pkg/api"
	"github.com/bestreyer/carfinder/pkg/route"
	"github.com/bestreyer/carfinder/pkg/validation"
)

type DI interface {
	GetTranslator() (ut.Translator)
	GetDbConn() (*sql.DB)
	GetValidator() (*validator.Validate)
	GetContextFactory() (context.Factory)
	GetServerFactory() (server.Factory)
	GetLocationRepository() (location.Repository)
	GetRouteCollection() ([]route.Route)
}

type di struct {
	translator         ut.Translator
	dbConn             *sql.DB
	validator          *validator.Validate
	serverFactory      server.Factory
	contextFactory     context.Factory
	locationRepository location.Repository
	routeCollection    []route.Route
}

func (d *di) GetRouteCollection() ([]route.Route) {
	if nil != d.routeCollection {
		return d.routeCollection
	}

	d.routeCollection = make([]route.Route, 2, 2)

	d.routeCollection[0] = route.New(
		"GET",
		"/api/v1/drivers",
		api.NewGetDriverController(d.GetLocationRepository()),
	)

	d.routeCollection[1] = route.New(
		"PUT",
		"/api/v1/drivers/:id/location",
		api.NewUpdateLocationController(d.GetLocationRepository()),
	)

	return d.routeCollection
}

func (d *di) GetLocationRepository() (location.Repository) {
	if nil != d.locationRepository {
		return d.locationRepository
	}

	d.locationRepository = location.NewPostgreRepository(d.GetDbConn())

	return d.locationRepository
}

func (d *di) GetValidator() (*validator.Validate) {
	if nil != d.validator {
		return d.validator
	}

	d.validator = validation.New(d.GetTranslator())

	return d.validator
}

func (d *di) GetContextFactory() (context.Factory) {
	if nil != d.contextFactory {
		return d.contextFactory
	}

	d.contextFactory = context.NewFactory(d.GetValidator(), d.GetTranslator());
	return d.contextFactory
}

func (d *di) GetServerFactory() (server.Factory) {
	if nil != d.serverFactory {
		return d.serverFactory
	}

	d.serverFactory = server.NewFactory(d.GetContextFactory(), d.GetRouteCollection())

	return d.serverFactory
}

func (d *di) GetTranslator() (ut.Translator) {
	if nil != d.translator {
		return d.translator
	}

	en := en.New()
	uni := ut.New(en, en)
	trans, found := uni.GetTranslator("en")

	if false == found {
		log.Fatal("Translator has not been found")
	}

	d.translator = trans
	return d.translator
}

func (d *di) GetDbConn() (*sql.DB) {
	if nil != d.dbConn {
		return d.dbConn
	}

	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		env.GetEnv("POSTGRES_USER", "root"),
		env.GetEnv("POSTGRES_PASS", "root"),
		env.GetEnv("POSTGRES_HOST", "127.0.0.1"),
		env.GetEnv("POSTGRES_DBNAME", "carfinder"),
	)

	db, err := sql.Open("postgres", connStr)
	db.SetMaxOpenConns(20)

	if nil != err {
		log.Fatal(err)
	}

	d.dbConn = db
	return d.dbConn
}

func New() (DI) {
	return &di{}
}
