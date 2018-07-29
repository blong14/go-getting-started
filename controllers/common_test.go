package controllers_test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestMain used to do setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
