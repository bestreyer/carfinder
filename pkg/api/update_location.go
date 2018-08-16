package api

import (
	"github.com/bestreyer/carfinder/pkg/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)


type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude" validate:"required"`
	Accuracy  float64 `json:"accuracy" validate:"required"`
}

func UpdateLocation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var location Location

	if err := r.Context().(context.ContextInterface).ShouldBindJSON(r.Body, &location); err != nil {
		r.Context().(context.ContextInterface).BadJSONResponse(w, err)
		return
	}


	r.Context().(context.ContextInterface).JSONResponse(w, nil, 204)
}
