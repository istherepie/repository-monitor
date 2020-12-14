package webserver

import (
	"encoding/json"
	"log"
	"net/http"
)

type JSONResponse struct {
	ServiceID string `json:"service_id"`
	Message   string `json:"message"`
}

type ServiceHandler struct {
	ID string
}

func (s ServiceHandler) encode(source interface{}, target *[]byte) error {
	var err error
	*target, err = json.Marshal(source)
	return err
}

func (s ServiceHandler) send(w http.ResponseWriter, response JSONResponse) {

	var jsonResponse []byte

	encodeErr := s.encode(response, &jsonResponse)

	if encodeErr != nil {
		http.Error(w, "UNKNOWN_ERROR", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (s ServiceHandler) Process(w http.ResponseWriter, r *http.Request) {

	// Access log
	log.Printf("[INFO] %v request from %v %v %v", s.ID, r.RemoteAddr, r.Method, r.RequestURI)

	response := JSONResponse{}
	response.ServiceID = s.ID
	response.Message = "Hello World!!!!"

	s.send(w, response)
}

func (s ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")

	// CORS header
	w.Header().Set("Access-Control-Allow-Origin", "*")

	s.Process(w, r)
}

func Mux(serviceID string) *http.ServeMux {

	// Multiplexer
	mux := http.NewServeMux()

	// Init handler
	handler := ServiceHandler{ID: serviceID}

	// Register route(s)
	mux.Handle("/", handler)

	return mux
}
