package api

import (
	"github.com/bestreyer/carfinder/pkg/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/bestreyer/carfinder/pkg/location"
)

func UpdateLocation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var l location.Location

	if err := r.Context().(context.Context).ShouldBindJSON(r.Body, &l); err != nil {
		r.Context().(context.Context).BadJSONResponse(w, err)
		return
	}


	r.Context().(context.Context).JSONResponse(w, nil, 204)
}
