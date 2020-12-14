package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/istherepie/request-monitor/webserver"
)

func run() int {
	listener, err := net.Listen("tcp", "localhost:8000")

	if err != nil {
		log.Println(err)
		return 1
	}

	defer listener.Close()

	mux := webserver.Mux("service-1")

	serveError := http.Serve(listener, mux)

	if serveError != nil {
		log.Println(serveError)
		return 1
	}

	return 0
}

func main() {
	code := run()
	os.Exit(code)
}
