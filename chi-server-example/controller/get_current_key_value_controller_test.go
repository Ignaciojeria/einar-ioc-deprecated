package controller

import (
	"example/chi_router"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCurrentKeyController(t *testing.T) {
	// load chi router
	chi_router.Mux.Load()

	// Load getCurrentKeyController
	ctrl, err := getCurrentKeyControllerDependency.Load()
	if err != nil {
		t.Errorf("Failed to load getCurrentKeyController: %s", err)
		return
	}
	controller := ctrl.(getCurrentKeyController)
	//Mock getCurrencKey Behaviour
	controller.getCurrentKey = func(key string) string {
		if key == "a" {
			return "valueA"
		}
		if key == "b" {
			return "valueB"
		}
		return ""
	}

	// Create an HTTP request and response recorder
	req, _ := http.NewRequest("GET", "/getKeys", nil)
	rr := httptest.NewRecorder()

	// Call the handle method
	controller.handle(rr, req)

	// Check the response body
	expected := "current key value for a : valueA\n" + "current key value for b : valueB"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
