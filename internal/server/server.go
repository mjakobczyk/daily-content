package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/mjakobczyk/daily-content/internal/newsapi"
)

type newsapiService interface {
	GetTopHeadlines() (newsapi.TopHeadlineDTO, error)
}

// Server struct holds connectors and settings.
type Server struct {
	config     *Config
	middleware alice.Chain
	router     mux.Router
	newsapi    newsapiService
}

// NewServer function creates new instance of Server.
func NewServer(c *Config, newsapi newsapiService) *Server {
	router := *mux.NewRouter()
	middleware := alice.New(
		Header("Access-Control-Allow-Origin", "*"),
		Header("Content-Type", "application/json"),
		Log(c.Logger.Type),
	)

	srv := Server{
		config:     c,
		middleware: middleware,
		router:     router,
		newsapi:    newsapi,
	}

	// TODO: add handling endpoints
	srv.router.HandleFunc("/randomstuff", srv.randomStuffGETHandler).Methods("GET")
	srv.router.HandleFunc("/headlines", srv.headlinesGETHandler).Methods("GET")

	return &srv
}

// Start function kicks off the server.
func (s *Server) Start() error {
	address := fmt.Sprintf("%s:%v", s.config.IP, s.config.Port)

	log.Println("Server starting at: ", address)

	return http.ListenAndServe(address, s.middleware.Then(&s.router))
}
