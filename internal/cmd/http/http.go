// Package http provides http services for access to the application
package http

import (
	"fmt"
	"github.com/ybalcin/cache-api/internal/core/application/ports"
)

// StartServer starts the http server
func StartServer() {
	ports.NewHttpServer()

	fmt.Println("httpserver started")
}
