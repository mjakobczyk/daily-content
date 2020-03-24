package article

import "github.com/mjakobczyk/daily-content/internal/newsapi"

// Repository of articles.
type Repository struct {
}

// NewRepository creates a new instance of Repository.
func NewRepository() *Repository {
	return &Repository{}
}

// GetAllArticles returns all available Articles.
func (r *Repository) GetAllArticles() []newsapi.ArticleDTO {
	// TODO: change this mocked implementation
	return []newsapi.ArticleDTO{
		newsapi.ArticleDTO{},
	}
}
