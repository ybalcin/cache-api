package in

import (
	"encoding/json"
	"errors"
	"github.com/ybalcin/cache-api/internal/application"
	"github.com/ybalcin/cache-api/internal/application/dtos"
	"github.com/ybalcin/cache-api/internal/common"
	"log"
	"net/http"
	"os"
	"strings"
)

type (
	httpServer struct {
		Application *application.Application
	}

	Handler struct {
		H      func(rw http.ResponseWriter, req *http.Request) error
		Method string
	}
)

const (
	contentType     = "Content-Type"
	applicationJson = "application/json"
	requestId       = "X-Request-ID"

	RequiredKey = "key is required"
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
func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// log incoming request as stdout
	defer logRequest(req)

	if req.Method != h.Method {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set(contentType, applicationJson)

	err := h.H(w, req)
	if err == nil {
		if e, ok := recover().(error); ok {
			err = e
		}
	}

	if err != nil {
		switch e := err.(type) {
		case common.Error:
			logger.Printf("HTTP status: %d - Message: %s", e.Status(), e.Error())
			http.Error(w, "", e.Status())
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}

// SetKeyHandler sets a key-value pair in cache
func (s *httpServer) SetKeyHandler(w http.ResponseWriter, req *http.Request) error {
	decoder := json.NewDecoder(req.Body)
	var dto dtos.CacheDto
	if err := decoder.Decode(&dto); err != nil {
		return common.NewStatusError(http.StatusBadRequest, err)
	}

	if err := s.Application.CacheService.Set(&dto); err != nil {
		return common.NewStatusError(http.StatusBadRequest, err)
	}

	return nil
}

// GetValueHandler gets value by key from cache
func (s *httpServer) GetValueHandler(w http.ResponseWriter, req *http.Request) error {
	values := strings.Split(req.RequestURI, "/")
	key := values[len(values)-1]
	if key == "" {
		return common.NewStatusError(http.StatusBadRequest, errors.New(RequiredKey))
	}

	dto, err := s.Application.CacheService.Get(key)
	if err != nil {
		return common.NewStatusError(http.StatusNotFound, err)
	}

	json.NewEncoder(w).Encode(dto)
	return nil
}

// FlushHandler clears cache
func (s *httpServer) FlushHandler(w http.ResponseWriter, req *http.Request) error {
	s.Application.CacheService.ClearAll()
	return nil
}

func logRequest(req *http.Request) {
	requestID := req.Header.Get(requestId)
	if requestID == "" {
		requestID = "unknown"
	}
	logger.Println(requestID, req.Method, req.URL.Path, req.RemoteAddr, req.UserAgent())
}
