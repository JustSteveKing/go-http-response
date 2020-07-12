package response

import (
	"encoding/json"
	"net/http"
)

// Response is a generic HTTP Response Struct
type Response struct {
	Data string `json:"data"`
}

// Send a HTTP Response
func Send(responseWriter http.ResponseWriter, code int, payload interface{}, contentType string) {
	response, _ := json.Marshal(payload)

	responseWriter.Header().Set("Content-Type", contentType)
	responseWriter.WriteHeader(code)
	responseWriter.Write(response)
}
