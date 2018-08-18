package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type Controller interface {
	Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
