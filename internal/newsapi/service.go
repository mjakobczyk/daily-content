package newsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mjakobczyk/daily-content/env"
)

// Service communicates with NewsAPI server.
type Service struct {
	Config *Config
	Doer   httpDoer
}

// NewService returns new instance of Service using passed configuration.
func NewService(config *Config) *Service {
	return &Service{
		Config: config,
		Doer:   NewHTTPClient(config.HTTP.Client.Timeout.Seconds),
	}
}

// GetTopHeadlines requests NewsAPI for news' top headlines.
func (s *Service) GetTopHeadlines() (TopHeadlineDTO, error) {
	var topHeadlineDTO TopHeadlineDTO

	url := fmt.Sprintf("%s/top-headlines?country=us&apiKey=%s",
		s.Config.Host.API,
		s.Config.Host.Key,
	)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return TopHeadlineDTO{}, err
	}

	response, err := s.Doer.Do(request)
	if err == nil {
		if response.StatusCode >= 400 && response.StatusCode < 500 {
			return TopHeadlineDTO{}, env.RequestFailedError
		}
	} else {
		return TopHeadlineDTO{}, err
	}

	defer func() { _ = response.Body.Close() }()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return TopHeadlineDTO{}, env.ReadDataFailedError
	}

	err = json.Unmarshal(data, &topHeadlineDTO)
	if err != nil {
		return TopHeadlineDTO{}, env.UnmarshalDataFailedError
	}

	return topHeadlineDTO, nil
}
