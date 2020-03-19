package newsapi

import "encoding/json"

// ArticleDTO type.
type ArticleDTO struct {
	Source      SourceDTO `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt string    `json:"publishedAt"`
	Content     string    `json:"content"`
}

type articleRequestDTO struct {
	Source      SourceDTO `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt string    `json:"publishedAt"`
	Content     string    `json:"content"`
}

// UnmarshalJSON provides custom unmarshaling for ArticleDTO type.
func (a *ArticleDTO) UnmarshalJSON(data []byte) error {
	var articleRequestDTO articleRequestDTO

	if err := json.Unmarshal(data, &articleRequestDTO); err != nil {
		return err
	}

	a.Source = articleRequestDTO.Source
	a.Author = articleRequestDTO.Author
	a.Title = articleRequestDTO.Title
	a.Description = articleRequestDTO.Description
	a.URL = articleRequestDTO.URL
	a.URLToImage = articleRequestDTO.URLToImage
	a.PublishedAt = articleRequestDTO.PublishedAt
	a.Content = articleRequestDTO.Content

	return nil
}
