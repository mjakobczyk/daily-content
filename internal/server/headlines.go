package server

import (
	"encoding/json"
	"net/http"
)

func (s *Server) headlinesGETHandler(w http.ResponseWriter, r *http.Request) {
	defer func() { _ = r.Body.Close() }()

	response, err := s.newsapi.GetTopHeadlines()
	if err != nil {
		_ = NewAPIError(err.Error(), http.StatusInternalServerError).Send(w)
		return
	}

	body, err := json.Marshal(response)
	if err != nil {
		_ = NewAPIError(err.Error(), http.StatusInternalServerError).Send(w)
		return
	}

	_, _ = w.Write(body)
}
