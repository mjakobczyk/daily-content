package env

import (
	"github.com/go-http-utils/logger"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// Environment stores all components that should be
// shared between packages in application.
type Environment struct {
	Middleware alice.Chain
	Router     mux.Router
}

func NewEnvironment(logger logger.Type) *Environment {
	return &Environment{
		Router: *mux.NewRouter(),
		Middleware: alice.New(
			Header("Access-Control-Allow-Origin", "*"),
			Header("Content-Type", "application/json"),
			Log(logger),
		),
	}
}
