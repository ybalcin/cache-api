// Package http provides http services for access to the application
package http

import (
	"fmt"
	"github.com/ybalcin/cache-api/internal/ports/in"
	"log"
	"net/http"
)

const (
	pathPrefix string = "/v1/cache"

	port string = "8080"
)

func path(suffix string) string {
	return fmt.Sprintf("%s%s", pathPrefix, suffix)
}

// StartServer starts the http server
func StartServer() {
	httpPort := in.NewHttpServer()

	mux := http.NewServeMux()

	mux.Handle(path("/set"), in.Handler{H: httpPort.SetKeyHandler, Method: http.MethodPost})
	mux.Handle(path("/get/"), in.Handler{H: httpPort.GetValueHandler, Method: http.MethodGet})
	mux.Handle(path("/flush"), in.Handler{H: httpPort.FlushHandler, Method: http.MethodDelete})

	log.Printf("http server listening on port: %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
