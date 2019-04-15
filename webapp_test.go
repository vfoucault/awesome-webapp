package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vfoucault/awesome-webapp"
)

var app main.App

func TestHomePage(t *testing.T) {
	app.Create()
	app.AddRoutes()

	req, _ := http.NewRequest("GET", "/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	expected := `{
    "data": {
        "page": "home page"
    }
}`

	if body := response.Body.String(); body != expected {
		t.Errorf("Expected '%v'. Got %s", expected, body)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

