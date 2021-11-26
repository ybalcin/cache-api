// Package http provides http services for access to the application
package http

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/ybalcin/cache-api/internal/ports/in"
	"log"
)

const (
	port string = "8080"
)

// StartServer starts the http server
func StartServer() {
	httpPort := in.NewHttpServer()

	app := fiber.New()

	v1 := app.Group("/v1/cache")
	v1.Post("/", httpPort.SetKeyHandler)
	v1.Get("/:key", httpPort.GetValueHandler)
	v1.Delete("/flush", httpPort.FlushHandler)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
