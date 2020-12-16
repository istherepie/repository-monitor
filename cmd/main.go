package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/istherepie/request-echo/webserver"
)

func run(ID string, hostname string, port int, logger *log.Logger) int {
	// Setup listener
	address := fmt.Sprintf("%v:%d", hostname, port)
	listener, err := net.Listen("tcp", address)

	if err != nil {
		logger.Println(err)
		return 1
	}

	defer listener.Close()

	// Init multiplexer
	mux := webserver.Mux(ID, logger)

	// Serve
	logger.Printf("Starting webserver on address %v\n", address)
	serveError := http.Serve(listener, mux)

	if serveError != nil {
		logger.Println(serveError)
		return 1
	}

	return 0
}

func main() {

	// Service port
	var port int

	// Get Args
	serviceID := flag.String("id", "service-1", "Meta ID/name of the service")
	servicePort := flag.Int("port", 8080, "Service Port (tcp)")
	serviceHost := flag.String("host", "localhost", "Service hostname/IP address")
	flag.Parse()

	if !flag.Parsed() {
		flag.Usage()
		return
	}

	// GET ENV VARS
	envServicePort := os.Getenv("RM_SERVICE_PORT")

	// Setup logger
	loggerPrefix := fmt.Sprintf("[%v] ", *serviceID)
	logger := log.New(os.Stdout, loggerPrefix, log.Ldate|log.Ltime)

	altPort, err := strconv.ParseInt(envServicePort, 10, 64)

	// Port set by env variable overrides flag
	if err == nil {
		port = int(altPort)
	} else {
		port = *servicePort
	}

	// Start
	code := run(*serviceID, *serviceHost, port, logger)
	os.Exit(code)
}
