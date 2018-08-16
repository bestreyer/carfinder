package http

import (
	"fmt"
	"github.com/bestreyer/carfinder/pkg/api"
	"github.com/bestreyer/carfinder/pkg/context"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type HTTPServerInterface interface {
	Start(addr string) error
}

type HTTPServerFactoryInterface interface {
	CreateHTTPServer() (HTTPServerInterface, error)
}

type HTTPServer struct {
	http.Server
}

func (h HTTPServer) Start(addr string) error {
	h.Addr = addr

	return h.ListenAndServe()
}

type HTTPServerFactory struct {
	ContextFactory context.ContextFactoryInterface
}

func (hf *HTTPServerFactory) CreateHTTPServer() (HTTPServerInterface, error) {
	r := httprouter.New()

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		log.Println(i)
	}

	r.GET("/drivers", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "/drivers")
	})

	r.PUT("/drivers/:id/location", api.UpdateLocation)

	h := &HTTPServer{
		Server: http.Server{
			Handler: hf.replaceContextMiddleware(r),
		},
	}

	return h, nil
}

func (hf *HTTPServerFactory) replaceContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := hf.ContextFactory.CreateContext(r.Context())
		if nil != err {
			log.Fatal(err)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
