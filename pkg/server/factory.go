package server

import (
	"github.com/bestreyer/carfinder/pkg/context"
	"github.com/bestreyer/carfinder/pkg/route"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
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
		f.createHandleByRoute(r, rt)
	}

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		log.Println(i)
	}

	r.NotFound = func(w http.ResponseWriter, r *http.Request) {
		r.Context().(context.Context).JSONResponse(w, nil, 404)
	}

	s := &server{
		Server: http.Server{
			Handler: f.replaceContextMiddleware(r, f.cf),
		},
	}

	return s, nil
}

func (f *factory) createHandleByRoute(router *httprouter.Router, rt route.Route) {
	router.Handle(rt.Method(), rt.Path(), func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rt.Controller().Handle(w, r, ps)
	})
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

func NewFactory(cf context.Factory, routeCollection []route.Route) Factory {
	return &factory{cf: cf, routeCollection: routeCollection}
}
