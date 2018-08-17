package server

import (
	"github.com/bestreyer/carfinder/pkg/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"fmt"
	"github.com/bestreyer/carfinder/pkg/api"
)

type Factory interface {
	Create() (Server, error)
}

type factory struct {
	cf context.Factory
}

func (f *factory) Create() (Server, error) {
	r := httprouter.New()

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		log.Println(i)
	}

	r.GET("/api/v1/drivers", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "/drivers")
	})

	r.PUT("/api/v1/drivers/:id/location", api.UpdateLocation)

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

func NewFactory(cf context.Factory) (Factory) {
	return &factory{cf: cf}
}
