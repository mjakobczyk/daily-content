package newsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func (s *Service) GetTopHeadlines() (TopHeadlineDTO, error) {
	var topHeadlineDTO TopHeadlineDTO

	url := fmt.Sprintf("%s/top-headlines?country=us&apiKey=%s",
		s.Config.Host.API,
		s.Config.Host.Key,
	)

	fmt.Println(url)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return TopHeadlineDTO{}, err
	}

	response, err := s.Doer.Do(request)
	if err == nil {
		if response.StatusCode >= 400 && response.StatusCode < 500 {
			return TopHeadlineDTO{}, RequestFailedError
		}
	} else {
		return TopHeadlineDTO{}, err
	}

	fmt.Println("Ree2")

	defer func() { _ = response.Body.Close() }()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return TopHeadlineDTO{}, ReadDataFailedError
	}

	err = json.Unmarshal(data, &topHeadlineDTO)

	if err != nil {
		return TopHeadlineDTO{}, UnmarshalDataFailedError
	}

	return topHeadlineDTO, nil
}
