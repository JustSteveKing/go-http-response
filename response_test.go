package response

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResponse(t *testing.T) {
	t.Run("Test something", func(t *testing.T) {
		handler := func(w http.ResponseWriter, r *http.Request) {
			body := &Response{
				Data: "foo",
			}
			Send(w, 200, body, "application/json")
		}

		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		w := httptest.NewRecorder()
		handler(w, req)
	})
}
