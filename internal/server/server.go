package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mjakobczyk/daily-content/env"
	"github.com/mjakobczyk/daily-content/internal/newsapi"
)

type newsapiService interface {
	GetTopHeadlines() (newsapi.TopHeadlineDTO, error)
}

// Server struct holds connectors and settings.
type Server struct {
	config  Config
	env     env.Environment
	newsapi newsapiService
}

// NewServer function creates new instance of Server.
func NewServer(c Config, e env.Environment, newsapi newsapiService) *Server {
	srv := Server{
		config:  c,
		env:     e,
		newsapi: newsapi,
	}

	srv.env.Router.HandleFunc("/randomstuff", srv.randomStuffGETHandler).Methods("GET")
	srv.env.Router.HandleFunc("/headlines", srv.headlinesGETHandler).Methods("GET")

	return &srv
}

// Start function kicks off the server.
func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%v", s.config.IP, s.config.Port)
	log.Println("Server starting at: ", address)

	return http.ListenAndServe(address, s.env.Middleware.Then(&s.env.Router))
}
