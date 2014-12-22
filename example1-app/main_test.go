package main

import (
	"testing"
	"net/http"
	"strings"
	"io/ioutil"
)

func TestEventsGet(t *testing.T) {
	res, err := http.Get("http://localhost:8080/events")

	if res.StatusCode != 200 {
		t.Error(err, res)
	}
}

func TestEventsPost(t *testing.T) {
	json := `{"key":"value"}`
	b := strings.NewReader(json)
	res, err := http.Post("http://localhost:8080/events", "application/json", b)

	if res.StatusCode != 200 {
		t.Error(err, res)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	if string(body) != "{\"Name\":\"abc\"}" {
		t.Error("dif", res)
	}

}
