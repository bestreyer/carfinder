package integration

import (
	"testing"
	"github.com/bestreyer/carfinder/pkg/location"
	"github.com/bestreyer/carfinder/pkg/di"
	"context"
	"net/http"
	"fmt"
	"github.com/bestreyer/carfinder/pkg/env"
	"strings"
	"io/ioutil"
)

func TestUpdateLocation(t *testing.T) {
	d := di.New()
	l := location.Location{}

	err := d.GetLocationRepository().Create(context.Background(), &l)
	if err != nil {
		t.Error(err.Error())
	}

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

	c := &http.Client{}

	for _, r := range tables {
		resp := sendUpdateLocationRequest(t, c, r.json, r.id)
		contentB, err := ioutil.ReadAll(resp.Body)
		content := string(contentB)

		if nil != err {
			t.Error(err.Error())
		}

		if resp.StatusCode != r.expCode {
			t.Errorf("Invalid response http code, actual: %d, expected: %d", resp.StatusCode, r.expCode)
		}

		if content != r.expBody {
			t.Errorf("Invalid response http body, actual: %s, expected: %s", content, r.expBody)
		}
	}
}

func sendUpdateLocationRequest(t *testing.T, c *http.Client, json string, id int) (*http.Response) {
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

	resp, err := c.Do(req)

	if nil != err {
		t.Error(err.Error())
	}

	return resp
}
