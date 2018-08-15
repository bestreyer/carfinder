package server

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/bestreyer/carfinder/api"
	"log"
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
	ctxf ContextFactoryInterface
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
		ctx, err := hf.ctxf.CreateContext(r.Context())
		if nil != err {
			log.Fatal(err)
		}

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
