// Package http provides http services for access to the application
package http

import (
	"fmt"
	"github.com/ybalcin/cache-api/internal/ports/in"
)

// StartServer starts the http server
func StartServer() {
	in.NewHttpServer()

	fmt.Println("httpserver started")
}
