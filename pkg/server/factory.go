package server

import (
	"github.com/bestreyer/carfinder/pkg/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"github.com/bestreyer/carfinder/pkg/route"
)

type Factory interface {
	Create() (Server, error)
}

type factory struct {
	routeCollection []route.Route
	cf              context.Factory
}

func (f *factory) Create() (Server, error) {
	r := httprouter.New()

	for _, rt := range f.routeCollection {
		r.Handle(rt.Method(), rt.Path(), func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			rt.Controller().Handle(w, r, ps)
		})
	}

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		log.Println(i)
	}

	s := &server{
		Server: http.Server{
			Handler: f.replaceContextMiddleware(r, f.cf),
		},
	}

	return s, nil
}

func (f *factory) replaceContextMiddleware(next http.Handler, cf context.Factory) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := cf.CreateContext(r.Context())
		if nil != err {
			log.Fatal(err)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func NewFactory(cf context.Factory, routeCollection []route.Route) (Factory) {
	return &factory{cf: cf, routeCollection: routeCollection}
}
