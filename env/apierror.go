package env

import (
	"encoding/json"
	"net/http"
)

// APIResponse returned from endpoints in JSON form.
type APIResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// APIResponse is a default constructor of APIResponse.
func NewAPIResponse(msg string, code int) APIResponse {
	return APIResponse{msg, code}
}

// Send marshaled APIError with given code.
func (a APIResponse) Send(resp http.ResponseWriter) error {
	out, _ := json.Marshal(a)
	resp.Header().Set("X-Content-Type-Options", "nosniff")
	resp.WriteHeader(a.Code)
	_, err := resp.Write(out)

	return err
}
