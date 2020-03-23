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

func NewRouter(e *env.Environment, s service) *Router {
	router := &Router{
		env:     e,
		service: s,
	}

	router.initRoutes()
	return router
}

type service interface {
	GetAllArticles() []newsapi.ArticleDTO
	GetLatestArticles() []newsapi.ArticleDTO
}

func (r *Router) initRoutes() {
	r.env.Router.Get("/articles", r.getAllArticlesHandler)
}

func (r *Router) getAllArticlesHandler(resp http.ResponseWriter, req *http.Request) {
	articles := r.service.GetAllArticles()

	body, err := json.Marshal(articles)
	if err != nil {
		log.Println(err)
		_ = env.NewAPIError(env.ReadDataFailedError.String(), http.StatusInternalServerError).Send(resp)
	}

	_, _ = resp.Write(body)
}
