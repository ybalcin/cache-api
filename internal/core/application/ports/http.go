package ports

import (
	"github.com/ybalcin/cache-api/internal/core/application"
)

type (
	httpServer struct {
		Application *application.Application
	}
)

// NewHttpServer initializes a new http server input port
func NewHttpServer() *httpServer {
	app := application.New()

	return &httpServer{app}
}
