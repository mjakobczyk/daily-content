package web

import (
	"encoding/json"
	"net/http"
)

func (s *Server) randomStuffGETHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()

	var err error

	if err != nil {
		_ = NewAPIError(err.Error(), 500).Send(w)
		return
	}

	response := "Random stuff"

	body, err := json.Marshal(response)
	if err != nil {
		_ = NewAPIError(err.Error(), http.StatusInternalServerError).Send(w)
		return
	}

	_, _ = w.Write(body)
}
