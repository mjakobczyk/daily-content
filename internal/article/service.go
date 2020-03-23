package article

import (
	"github.com/mjakobczyk/daily-content/env"
	"github.com/mjakobczyk/daily-content/internal/newsapi"
)

// Service that provides articles data.
type Service struct {
	Env        *env.Environment
	repository repository
}

type repository interface {
	GetAllArticles() []newsapi.ArticleDTO
}

// NewService creates a new instance of Service.
func NewService(e *env.Environment, r repository) *Service {
	return &Service{
		Env:        e,
		repository: r,
	}
}

// GetLatestArticles returns all Articles from the system.
func (s *Service) GetAllArticles() []newsapi.ArticleDTO {
	return s.repository.GetAllArticles()
}

// GetLatestArticles returns recently added articles.
func (s *Service) GetLatestArticles() []newsapi.ArticleDTO {
	return s.repository.GetAllArticles()
}
