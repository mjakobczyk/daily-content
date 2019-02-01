package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Server struct holds connectors and settings.
type Server struct {
	config     *Config
	middleware alice.Chain
	router     mux.Router
}

// NewServer function creates new instance of Server.
func NewServer(c *Config) *Server {
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
	}

	// TODO: add handling endpoints
	// 	srv.Router.HandleFunc("/endpoint", srv.GEThandler).Methods("GET")

	return &srv
}

// Start function kicks off the server.
func Start(s *Server) error {
	address := fmt.Sprintf("%s:%v", s.config.IP, s.config.Port)

	log.Println("Server starting at: ", address)

	return http.ListenAndServe(address, s.middleware.Then(&s.router))
}
