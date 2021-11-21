package in_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ybalcin/cache-api/internal/application"
	"github.com/ybalcin/cache-api/internal/application/dtos"
	"github.com/ybalcin/cache-api/internal/common"
	"github.com/ybalcin/cache-api/internal/ports/in"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func mockRequest(reqBody interface{}) (w http.ResponseWriter, r *http.Request, err error) {
	r = new(http.Request)
	w = httptest.NewRecorder()

	if reqBody != nil {
		b, err := json.Marshal(reqBody)
		if err != nil {
			return nil, nil, err
		}
		r.Body = ioutil.NopCloser(bytes.NewReader(b))
	}

	return w, r, nil
}

func fatalF(t *testing.T, arg interface{}) {
	t.Fatalf("%v", arg)
}

func mustEqual(t *testing.T, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected: %#v, but got %#v", actual, expected)
	}
}

func mustNotNil(t *testing.T, val interface{}) {
	if val == nil {
		t.Fatalf("%#v must not nil", val)
	}
}

const (
	dummyKey   = "key1"
	dummyValue = "value1"
)

func TestHttpServer_SetKeyHandler(t *testing.T) {

	testCases := []struct {
		dto        interface{}
		expected   error
		code       int
		setInvoked bool
	}{
		{dtos.CacheDto{Key: dummyKey, Value: dummyValue}, nil, 0, true},
		{dtos.CacheDto{Key: "", Value: dummyValue}, errors.New("inmemorystore: key is empty"), http.StatusBadRequest, true},
		{dtos.CacheDto{Key: dummyKey, Value: ""}, errors.New("inmemorystore: key is empty"), http.StatusBadRequest, true},
	}

	mockCacheService := new(application.MockCacheService)
	mockApplication := application.Application{CacheService: mockCacheService}
	mockHttp := in.NewHttpServerWithApplication(&mockApplication)

	for _, c := range testCases {

		mockCacheService.SetFnInvoked = false
		mockCacheService.SetFn = func(dto *dtos.CacheDto) error {
			return c.expected
		}

		w, r, err := mockRequest(&c.dto)
		if err != nil {
			fatalF(t, err)
		}

		err = mockHttp.SetKeyHandler(w, r)
		if err != nil {
			e, ok := err.(common.StatusError)
			if ok {
				mustEqual(t, e, common.StatusError{Code: c.code, Err: c.expected})
			}
		}

		mustEqual(t, mockCacheService.SetFnInvoked, c.setInvoked)
	}
}

func TestHttpServer_GetValueHandler(t *testing.T) {
	dto := dtos.CacheDto{
		Key:   dummyKey,
		Value: dummyValue,
	}

	testCases := []struct {
		key        string
		err        error
		code       int
		expected   *dtos.CacheDto
		getInvoked bool
	}{
		{dummyKey, nil, 0, &dto, true},
		{"", errors.New(in.RequiredKey), http.StatusBadRequest, nil, false},
	}

	mockCacheService := new(application.MockCacheService)
	mockApplication := application.Application{CacheService: mockCacheService}
	mockHttp := in.NewHttpServerWithApplication(&mockApplication)

	for _, c := range testCases {

		mockCacheService.GetFnInvoked = false
		mockCacheService.GetFn = func(key string) (*dtos.CacheDto, error) {
			return c.expected, nil
		}

		w, r, err := mockRequest(nil)
		if err != nil {
			fatalF(t, err)
		}
		r.RequestURI = fmt.Sprintf("/%s", c.key)

		err = mockHttp.GetValueHandler(w, r)
		if err != nil {
			mustEqual(t, err, common.StatusError{Code: c.code, Err: c.err})
		}

		mustEqual(t, mockCacheService.GetFnInvoked, c.getInvoked)
	}
}

func TestHttpServer_FlushHandler(t *testing.T) {
	mockCacheService := new(application.MockCacheService)

	mockCacheService.ClearAllInvoked = false
	mockCacheService.ClearAllFn = func() {
	}

	w, r, err := mockRequest(nil)
	if err != nil {
		fatalF(t, err)
	}
	mockApplication := application.Application{CacheService: mockCacheService}
	mockHttp := in.NewHttpServerWithApplication(&mockApplication)

	err = mockHttp.FlushHandler(w, r)
	mustEqual(t, err, nil)

	mustEqual(t, mockCacheService.ClearAllInvoked, true)
}
