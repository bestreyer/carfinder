package integration

import (
	"testing"
	"github.com/bestreyer/carfinder/pkg/di"
	"context"
	"net/http"
	"fmt"
	"github.com/bestreyer/carfinder/pkg/env"
	"strings"
	"github.com/bestreyer/carfinder/pkg/location"
)

func TestUpdateLocationValidation(t *testing.T) {
	tables := []struct {
		id      int
		json    string
		expCode int
		expBody string
	}{
		{
			-1,
			"{}",
			404,
			"{}",
		},
		{
			50001,
			"{}",
			404,
			"{}",
		},
		{
			1,
			"{}",
			422,
			"{\"errors\":[\"Latitude is required\",\"Longitude is required\",\"Accuracy is required\"]}",
		},
		{
			1,
			"{\"latitude\": 45, \"longitude\": 181, \"accuracy\": 0.0001}",
			422,
			"{\"errors\":[\"Longitude should be between +/- 180\"]}",
		},

		{
			1,
			"{\"latitude\": -91, \"longitude\": 180, \"accuracy\": 0.0001}",
			422,
			"{\"errors\":[\"Latitude should be between +/- 90\"]}",
		},
		{
			1,
			"{\"latitude\": -81, \"longitude\": 180, \"accuracy\": 1.0001}",
			422,
			"{\"errors\":[\"Accuracy should be between [0..1]\"]}",
		},
		{
			1,
			"{\"latitude\": -81, \"longitude\": 180, \"accuracy\": 1}",
			200,
			"{}",
		},
	}

	for _, r := range tables {
		resp := sendUpdateLocationRequest(t, r.json, r.id)
		checkStatusAndBody(t, resp, r.expCode, r.expBody)
	}
}

func TestUpdateLocation(t *testing.T) {
	d := di.New()

	addDriverLocation(d, t, &location.Location{})
	addDriverLocation(d, t, &location.Location{})

	sendUpdateLocationRequest(t, "{\"latitude\": -81, \"longitude\": 45.1, \"accuracy\": 1}", 1)
	sendUpdateLocationRequest(t, "{\"latitude\": -81, \"longitude\": 45.1, \"accuracy\": 1}", 2)

	checkCountDriverLocationByLocation(d, t, 2, -81, 45.1)

	sendUpdateLocationRequest(t,  "{\"latitude\": -81, \"longitude\": 45.2, \"accuracy\": 1}", 2)

	checkCountDriverLocationByLocation(d, t, 1, -81, 45.1)
	checkCountDriverLocationByLocation(d, t, 1, -81, 45.2)

}

func checkCountDriverLocationByLocation(d di.DI, t *testing.T, count int, lat float64, lon float64) {
	var c int
	d.GetDbConn().QueryRowContext(
		context.Background(),
		"SELECT COUNT(*) FROM driver_location WHERE latitude=$1 and longitude=$2", lat, lon,
	).Scan(&c)

	if c != count {
		t.Errorf("Expected amount of drivers with latitude(%f) and longitude(%f): %d, actual: %d",
			lat, lon, count, c)
	}
}

func sendUpdateLocationRequest(t *testing.T, json string, id int) (*http.Response) {
	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf(
			"http://%s/api/v1/drivers/%d/location",
			env.GetEnv("CARFINDER_ADDR", ""),
			id,
		),
		strings.NewReader(json),
	)


	if nil != err {
		t.Error(err.Error())
	}

	return sendRequest(t, req)
}

