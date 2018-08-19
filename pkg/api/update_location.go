package api

import (
	"github.com/bestreyer/carfinder/pkg/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/bestreyer/carfinder/pkg/location"
	"time"
	"strconv"
)

type updateLocationController struct {
	lr location.Repository
}

func (uc *updateLocationController) Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if nil != err || id > 50000 || id < 1 {
		r.Context().(context.Context).JSONResponse(w, nil, 404)
		return;
	}

	var l location.UpdateLocation
	if err := r.Context().(context.Context).ShouldBindJSON(r, &l); err != nil {
		r.Context().(context.Context).BadJSONResponse(w, err)
		return
	}

	l.DriverId = id
	l.UpdatedAt = time.Now()

	uc.lr.Update(r.Context(), &l)

	r.Context().(context.Context).JSONResponse(w, nil, 200)
}

func NewUpdateLocationController(lr location.Repository) (Controller) {
	return &updateLocationController{lr: lr}
}
