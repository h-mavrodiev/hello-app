package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloRouteDefault(t *testing.T) {

	expected_message := `{"message":"Hello World!"}`

	router := helloAppRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected_message, w.Body.String())
}

func TestHelloRouteWithName(t *testing.T) {

	expected_message := `{"message":"Hello Hristo!"}`

	router := helloAppRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello?name=Hristo", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected_message, w.Body.String())
}

func TestBreakRoute(t *testing.T) {

	expected_message := `{"message":"Break request was successful!"}`

	router := helloAppRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/break", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected_message, w.Body.String())
}

func TestHealthzRouteOk(t *testing.T) {

	expected_message := `{"message":"HTTP status 200"}`

	router := helloAppRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/healthz", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, expected_message, w.Body.String())
}

func TestHealthzRouteError(t *testing.T) {

	expected_message := `{"message":"HTTP status 500"}`

	router := helloAppRouter()
	w := httptest.NewRecorder()
	req_break, _ := http.NewRequest("POST", "/break", nil)
	router.ServeHTTP(w, req_break)

	req, _ := http.NewRequest("GET", "/healthz", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, expected_message, w.Body.String())
}
