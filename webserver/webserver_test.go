package webserver

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerStatusCode(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	// setup handler
	logger := log.New(ioutil.Discard, "", log.Ldate|log.Ltime)
	handler := Mux("test-service-1", logger)

	// Run handler
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected: %v - GOT: %v", http.StatusOK, rec.Code)
	}

}

func TestHandlerResponse(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	// setup handler
	logger := log.New(ioutil.Discard, "", log.Ldate|log.Ltime)
	handler := Mux("test-service-1", logger)

	// Run handler
	handler.ServeHTTP(rec, req)

	expected := `{"service_id":"test-service-1","message":"Hello World!!!!"}`

	if rec.Body.String() != expected {
		t.Errorf("Expected: %v - GOT: %s", expected, rec.Body)
	}

}
