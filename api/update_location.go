package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"log"
	"github.com/bestreyer/carfinder/model"
)

func UpdateLocation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.Context()
	var location model.Location
	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
	}

	log.Println(location);
}
