package users_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/blong14/goping-web/config"
	"github.com/gin-gonic/gin"
)

func TestGetPingUnAuthorized(t *testing.T) {
	req, _ := http.NewRequest("GET", "/ping", nil)
	response := config.ExecuteRequest(req)

	expected := http.StatusUnauthorized

	if expected != response.Code {
		t.Errorf("Expected response code %d. Got %d\n", expected, response.Code)
	}
}

// TestMain used to do setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
