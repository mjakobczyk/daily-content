package article

import "github.com/mjakobczyk/daily-content/internal/newsapi"

// Repository of articles.
type Repository struct {
}

// NewRepository creates a new instance of Repository.
func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) GetAllArticles() []newsapi.ArticleDTO {
	return []newsapi.ArticleDTO{
		newsapi.ArticleDTO{},
	}
}
