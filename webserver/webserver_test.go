package webserver

import (
	"encoding/json"
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

	var body map[string]interface{}
	err = json.NewDecoder(rec.Body).Decode(&body)

	if err != nil {
		t.Errorf("Failed decoding response body: %v\n", err)
	}

	result := body["message"]
	expected := "Hello World!!!!"

	if result != expected {
		t.Errorf("Expected: %v - GOT: %s", expected, result)
	}

}
