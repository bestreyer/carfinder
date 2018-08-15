package api

import (
	"github.com/bestreyer/carfinder/context"
	"github.com/bestreyer/carfinder/model"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func UpdateLocation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var location model.Location

	if err := r.Context().(context.ContextInterface).ShouldBindJSON(r.Body, &location); err != nil {
		r.Context().(context.ContextInterface).BadJSONResponse(w, err)
		return
	}

	r.Context().(context.ContextInterface).JSONResponse(w, nil, 204)
}
