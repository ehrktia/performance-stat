package http_adapter

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handler(t *testing.T) {
	got := New()
	assert.NotNil(t, got)
	handler := got.GetRouter()
	t.Run("get handler", func(t *testing.T) {
		assert.NotNil(t, handler)
		if !assert.IsType(t, &http.ServeMux{}, handler) {
			t.Fatalf("expected:%v\n,got:%v\n", "http.handler", handler)
		}
	})
}
