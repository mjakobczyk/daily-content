package article

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mjakobczyk/daily-content/env"
	"github.com/mjakobczyk/daily-content/internal/newsapi"
)

// Router defines routes for Articles.
type Router struct {
	env     *env.Environment
	service service
}

// NewRouter is a default constructor of Router.
func NewRouter(e *env.Environment, s service) *Router {
	return &Router{
		env:     e,
		service: s,
	}
}

type service interface {
	GetAllArticles() []newsapi.Article
	GetLatestArticles() []newsapi.Article
}

// InitRoutes of articles.
func (r *Router) InitRoutes() {
	r.env.Router.Get("/articles", r.getAllArticlesHandler)
}

func (r *Router) getAllArticlesHandler(resp http.ResponseWriter, req *http.Request) {
	articles := r.service.GetAllArticles()

	body, err := json.Marshal(articles)
	if err != nil {
		log.Println(err)
		_ = env.NewAPIResponse(env.ReadDataFailedError.String(), http.StatusInternalServerError).Send(resp)
	}

	_, _ = resp.Write(body)
}
