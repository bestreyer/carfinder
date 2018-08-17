package server

import (
	"github.com/bestreyer/carfinder/pkg/context"
	"net/http"
)

type Server interface {
	Start(addr string) error
}

type server struct {
	http.Server
	h context.Factory
}

func (s server) Start(addr string) error {
	s.Addr = addr
	return s.ListenAndServe()
}
