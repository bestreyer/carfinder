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
	CreateHTTPServer() (*HTTPServerInterface, error)
}

type HTTPServer struct {
	http.Server
}

func (h HTTPServer) Start(addr string) error {
	h.Addr = addr

	return h.ListenAndServe()
}

type HTTPServerFactory struct{}

func (hf HTTPServerFactory) CreateHTTPServer() (*HTTPServerInterface, error) {
	r := httprouter.New()

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		log.Println(i)
	}

	r.GET("/drivers", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "/drivers")
	})

	r.PUT("/drivers/:id/location", api.UpdateLocation)

	//jm := middleware.JsonRequestMiddleware(r)

	h := &HTTPServer{
		Server: http.Server{
			Handler: r,
		},
	}

	return h, nil
}
