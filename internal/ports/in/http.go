package in

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ybalcin/cache-api/internal/application"
	"github.com/ybalcin/cache-api/internal/application/dtos"
	"github.com/ybalcin/cache-api/internal/common"
	"log"
	"net/http"
	"os"
)

type (
	httpServer struct {
		Application *application.Application
	}
)

const (
	requestId = "X-Request-ID"
)

var (
	logger = log.New(os.Stdout, "http: ", log.LstdFlags)
)

// NewHttpServer initializes a new http server input port
func NewHttpServer() *httpServer {
	app := application.New()

	return &httpServer{app}
}

// NewHttpServerWithApplication initializes new httpserver with application argument
func NewHttpServerWithApplication(application *application.Application) *httpServer {
	return &httpServer{Application: application}
}

// ServeHTTP middleware
//func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	// log incoming request as stdout
//	defer logRequest(req)
//
//	if req.Method != h.Method {
//		http.Error(w, "", http.StatusMethodNotAllowed)
//		return
//	}
//
//	w.Header().Set(contentType, applicationJson)
//
//	err := h.H(w, req)
//	if err == nil {
//		if e, ok := recover().(error); ok {
//			err = e
//		}
//	}
//
//	if err != nil {
//		switch e := err.(type) {
//		case common.Error:
//			logger.Printf("HTTP status: %d - Message: %s", e.Status(), e.Error())
//			http.Error(w, "", e.Status())
//		default:
//			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//		}
//	}
//}

func (s *httpServer) SetKeyHandler(c *fiber.Ctx) error {
	dto := new(dtos.CacheDto)
	var err error

	if err = c.BodyParser(dto); err != nil {
		// return err
	}

	if err = s.Application.CacheService.Set(dto); err != nil {
		// return err
	}

	return nil
}

func (s *httpServer) FlushHandler(c *fiber.Ctx) error {
	s.Application.CacheService.ClearAll()
	return nil
}

func (s *httpServer) GetValueHandler(c *fiber.Ctx) error {
	key := c.Params("key")

	dto, err := s.Application.CacheService.Get(key)
	if err != nil {
		return common.NewStatusError(http.StatusNotFound, err)
	}

	return c.JSON(dto)
}

func logRequest(req *http.Request) {
	requestID := req.Header.Get(requestId)
	if requestID == "" {
		requestID = "unknown"
	}
	logger.Println(requestID, req.Method, req.URL.Path, req.RemoteAddr, req.UserAgent())
}
