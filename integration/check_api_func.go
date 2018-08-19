package integration

import (
	"github.com/bestreyer/carfinder/pkg/di"
	"testing"
	"github.com/bestreyer/carfinder/pkg/location"
	"net/http"
	"context"
	"io/ioutil"
)

func addDriverLocation(d di.DI, t *testing.T, l *location.Location) {
	err := d.GetLocationRepository().Create(context.Background(), l)
	if err != nil {
		t.Error(err.Error())
	}
}

func sendRequest(t *testing.T, r *http.Request) (*http.Response) {
	resp, err := http.DefaultClient.Do(r)

	if nil != err {
		t.Error(err.Error())
	}

	return resp
}

func checkStatusAndBody(t *testing.T, resp *http.Response, expCode int, expBody string) {
	contentB, err := ioutil.ReadAll(resp.Body)
	content := string(contentB)

	if nil != err {
		t.Error(err.Error())
	}

	if resp.StatusCode != expCode {
		t.Errorf("Invalid response http code, actual: %d, expected: %d", resp.StatusCode, expCode)
	}

	if content != expBody {
		t.Errorf("Invalid response http body, actual: %s, expected: %s", content, expBody)
	}
}
