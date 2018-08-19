package integration

import (
	"testing"
	"net/http"
	"fmt"
	"github.com/bestreyer/carfinder/pkg/env"
	"github.com/bestreyer/carfinder/pkg/di"
	"github.com/bestreyer/carfinder/pkg/location"
	"encoding/json"
)

func TestGetDriversValidation(t *testing.T) {
	tables := []struct {
		queryString string
		expCode     int
		expBody     string
	}{
		{
			"",
			422,
			"{\"errors\":[\"Latitude is required\",\"Longitude is required\"]}",
		},
		{
			"latitude=181",
			422,
			"{\"errors\":[\"Latitude should be between +/- 90\",\"Longitude is required\"]}",
		},
		{
			"latitude=-91&longitude=181",
			422,
			"{\"errors\":[\"Latitude should be between +/- 90\",\"Longitude should be between +/- 180\"]}",
		},
		{
			"latitude=25&longitude=91&radius=1&limit=34",
			200,
			"[]",
		},
	}

	for _, r := range tables {
		resp := sendGetDriverLocationRequest(t, r.queryString)
		checkStatusAndBody(t, resp, r.expCode, r.expBody)
	}
}

func TestGetDriverInRadius(t *testing.T) {
	d := di.New()

	drPer := &location.Location{Latitude: 59.959232, Longitude: 30.320070}
	drPoch := &location.Location{Latitude: 59.960437, Longitude: 30.317402}

	addDriverLocation(d, t, drPer)
	addDriverLocation(d, t, drPoch)

	drivers := *getNearDrivers(t, "latitude=59.956206&longitude=30.318714&radius=360")
	checkAmountOfDrivers(t, &drivers, 1)

	if drivers[0].Id != drPer.DriverId {
		t.Errorf("Expected driver: %d, actual: %d", drPer.DriverId, drivers[0].Id)
	}

	if 0 == drivers[0].Distance {
		t.Errorf("Invalid distance value")
	}


	drivers = *getNearDrivers(t, "latitude=59.956206&longitude=30.318714&radius=600")
	checkAmountOfDrivers(t, &drivers, 2)

	drivers = *getNearDrivers(t, "latitude=59.956206&longitude=30.318714&radius=600&limit=1")
	checkAmountOfDrivers(t, &drivers, 1)
}

func sendGetDriverLocationRequest(t *testing.T, q string) (*http.Response) {
	req, err := http.NewRequest("GET", fmt.Sprintf(
		"http://%s/api/v1/drivers?%s",
		env.GetEnv("CARFINDER_ADDR", ""),
		q,
	), nil)


	if nil != err {
		t.Error(err.Error())
	}

	return sendRequest(t, req)
}

func checkAmountOfDrivers(t *testing.T, drivers *[]location.LocationWithDistance, count int) {
	if len(*drivers) != count {
		t.Errorf("Expected amount of near drivers: %d, actual: %d", 1, len(*drivers))
	}
}

func getNearDrivers(t* testing.T, q string) (*[]location.LocationWithDistance) {
	rlwd := make([]location.LocationWithDistance, 0)
	resp := sendGetDriverLocationRequest(t, q)

	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&rlwd)

	return &rlwd
}
