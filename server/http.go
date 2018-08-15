package server

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"github.com/bestreyer/carfinder/server/middleware"
)

type HTTPServerFactoryInterface interface {
	Create() (HTTPServerInterface, error)
}

type HTTPServerInterface interface {
	Start(addr string) error
}

type HTTPServer struct {
	http.Server
}

func (h HTTPServer) Start(addr string) error {
	h.Addr = addr

	return h.ListenAndServe()
}

type HTTPServerFactory struct{}

func (hf HTTPServerFactory) Create() (HTTPServerInterface, error) {
	r := httprouter.New()
	r.GET("/drivers", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		request.Context().
			fmt.Fprint(writer, "/drivers")
	})

	r.PUT("/drivers/:id/location", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprintf(writer, "Update driver location with id = %d", params.ByName("id"))
	})

	jm := middleware.JsonRequestMiddleware(r)
	
	h := &HTTPServer{
		Server: http.Server{Handler: jm},
	}

	return h, nil
}
