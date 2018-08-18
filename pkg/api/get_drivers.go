package api

import (
	"github.com/bestreyer/carfinder/pkg/location"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/bestreyer/carfinder/pkg/context"
	"github.com/labstack/gommon/log"
)

type getDriversController struct {
	lr location.Repository
}

func (uc *getDriversController) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var dal location.DriverAroundLocation

	if err := r.Context().(context.Context).ShouldBindQuery(r, &dal); err != nil {
		r.Context().(context.Context).BadJSONResponse(w, err)
		return
	}

	if 0 == dal.Radius {
		dal.Radius = 500
	}

	if 0 == dal.Limit {
		dal.Limit = 10
	}

	rows, err := uc.lr.GetDrivers(r.Context(), &dal)
	if nil != err {
		log.Error(err.Error())
		r.Context().(context.Context).InternalErrorResponse(w)
		return
	}

	r.Context().(context.Context).JSONResponse(w, rows, 200)
}

func NewGetDriverController(lr location.Repository) (Controller) {
	return &getDriversController{lr: lr}
}
