package webserver

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Headers map[string][]string

type JSONResponse struct {
	ServiceID string                 `json:"service_id"`
	Message   string                 `json:"message"`
	Headers   Headers                `json:"headers"`
	Params    url.Values             `json:"query_params"`
	Body      map[string]interface{} `json:"body"`
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
		values = append(values, value)
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
	s.Log.Printf("INFO request from %v %v %v", r.RemoteAddr, r.Method, r.RequestURI)

	// Prepare response
	response := JSONResponse{}
	response.ServiceID = s.ID
	response.Message = "Hello World!!!!"

	// Parse request headers
	response.Headers = s.GetHeaders(r.Header)

	// Parse query params
	queryParams := r.URL.Query()
	response.Params = s.GetParams(queryParams)

	// Parse request body
	// var parsedBody map[string]interface{}
	scanner := bufio.NewScanner(r.Body)

	data := scanner.Bytes()

	parseErr := json.Unmarshal(data, &response.Body)

	if parseErr != nil {
		s.Log.Printf("ERROR - Failed to decode body: %v\n", parseErr)
		response.Body = make(map[string]interface{})
	}

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
