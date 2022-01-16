package apiserver

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiServer_HandleHello(t *testing.T) {
	s := NewApiServer(NewConfig())
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/hello", nil)
	s.handleHello().ServeHTTP(w, r)
	assert.Equal(t, w.Body.String(), "Hello")
}
