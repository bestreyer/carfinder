package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Controller interface {
	Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
