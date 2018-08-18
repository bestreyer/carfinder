package route

import (
	"github.com/bestreyer/carfinder/pkg/api"
)

type Route interface {
	Method() (string)
	Path() (string)
	Controller() (api.Controller)
}

type route struct {
	path       string
	method     string
	controller api.Controller
}

func (r *route) Method() (string) {
	return r.method
}

func (r *route) Path() (string) {
	return r.path
}

func (r *route) Controller() (api.Controller) {
	return r.controller
}

func New(method string, path string, controller api.Controller) (Route) {
	return &route{path: path, method: method, controller: controller}
}
