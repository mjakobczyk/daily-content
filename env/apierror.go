package env

import (
	"encoding/json"
	"net/http"
)

// APIError returned from endpoints in json
type APIError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// NewAPIError is just a constructor. Nothing extraordinary
func NewAPIError(msg string, code int) APIError {
	return APIError{msg, code}
}

// Send sends a marshaled Error with given code
func (e APIError) Send(resp http.ResponseWriter) error {
	out, _ := json.Marshal(e)
	resp.Header().Set("X-Content-Type-Options", "nosniff")
	resp.WriteHeader(e.Code)
	_, err := resp.Write(out)

	return err
}
