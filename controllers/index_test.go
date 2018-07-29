package controllers_test

import (
	"net/http"
	"testing"

	"github.com/blong14/goping-web/helpers"
)

func TestGetIndex(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	response := helpers.ExecuteRequest(req)

	expected := http.StatusOK

	if expected != response.Code {
		t.Errorf("Expected response code %d. Got %d\n", expected, response.Code)
	}
}
