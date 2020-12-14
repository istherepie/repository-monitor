package webserver

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Headers map[string][]string

type JSONResponse struct {
	ServiceID string     `json:"service_id"`
	Message   string     `json:"message"`
	Headers   Headers    `json:"headers"`
	Params    url.Values `json:"params"`
}

type ServiceHandler struct {
	ID  string
	Log *log.Logger
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

func (s ServiceHandler) Reformat(data []string) []string {
	var values []string
	for _, value := range data {
		value = strings.ToLower(value)
		value = strings.Replace(value, ",", "", -1)
		values = strings.Fields(value)
	}
	return values
}

func (s ServiceHandler) GetHeaders(header http.Header) Headers {

	h := make(Headers)

	for headerKey, headerValues := range header {
		key := strings.ToLower(headerKey)
		h[key] = s.Reformat(headerValues)
	}

	return h
}

func (s ServiceHandler) GetParams(params url.Values) map[string][]string {

	h := make(map[string][]string)

	for pKey, pValues := range params {
		key := strings.ToLower(pKey)
		h[key] = s.Reformat(pValues)
	}

	return h
}

func (s ServiceHandler) Process(w http.ResponseWriter, r *http.Request) {

	// Access log
	s.Log.Printf("[INFO] %v request from %v %v %v", s.ID, r.RemoteAddr, r.Method, r.RequestURI)

	queryParams := r.URL.Query()

	response := JSONResponse{}
	response.ServiceID = s.ID
	response.Message = "Hello World!!!!"
	response.Headers = s.GetHeaders(r.Header)
	response.Params = s.GetParams(queryParams)

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

func Mux(serviceID string, logger *log.Logger) *http.ServeMux {

	// Multiplexer
	mux := http.NewServeMux()

	// Init handler
	handler := ServiceHandler{ID: serviceID, Log: logger}

	// Register route(s)
	mux.Handle("/", handler)

	return mux
}
