package newsapi

import "encoding/json"

// TopHeadline stores information about top headlines data from News API.
type TopHeadline struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type topHeadlineRequestDTO struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// UnmarshalJSON provides custom unmarshaling for TopHealineDTO type.
func (t *TopHeadline) UnmarshalJSON(data []byte) error {
	var topHeadlineRequestDTO topHeadlineRequestDTO

	if err := json.Unmarshal(data, &topHeadlineRequestDTO); err != nil {
		return err
	}

	t.Status = topHeadlineRequestDTO.Status
	t.TotalResults = topHeadlineRequestDTO.TotalResults
	t.Articles = topHeadlineRequestDTO.Articles

	return nil
}
