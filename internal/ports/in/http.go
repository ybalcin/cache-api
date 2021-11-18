package in

import (
	"encoding/json"
	"errors"
	"github.com/ybalcin/cache-api/internal/application"
	"github.com/ybalcin/cache-api/internal/application/dtos"
	"github.com/ybalcin/cache-api/internal/common"
	"log"
	"net/http"
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

// NewHttpServer initializes a new http server input port
func NewHttpServer() *httpServer {
	app := application.New()

	return &httpServer{app}
}

// ServeHTTP middleware
func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != h.Method {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err := h.H(w, req)
	if err != nil {
		switch e := err.(type) {
		case common.Error:
			log.Printf("HTTP status: %d - Message: %s", e.Status(), e.Error())
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
		return err
	}

	if err := s.Application.CacheService.Set(&dto); err != nil {
		return common.StatusError{Code: http.StatusBadRequest, Err: err}
	}

	return nil
}

// GetValueHandler gets value by key from cache
func (s *httpServer) GetValueHandler(w http.ResponseWriter, req *http.Request) error {
	values := strings.Split(req.RequestURI, "/")
	key := values[len(values)-1]
	if key == "" {
		return common.StatusError{Code: http.StatusBadRequest, Err: errors.New("")}
	}

	dto, err := s.Application.CacheService.Get(key)
	if err != nil {
		return common.StatusError{Code: http.StatusNotFound, Err: err}
	}

	json.NewEncoder(w).Encode(dto)
	return nil
}

// FlushHandler clears cache
func (s *httpServer) FlushHandler(w http.ResponseWriter, req *http.Request) error {
	s.Application.CacheService.ClearAll()
	return nil
}
