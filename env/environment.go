package env

import (
	"github.com/go-http-utils/logger"
	"github.com/jinzhu/gorm"
	"github.com/justinas/alice"
)

// Environment stores all components that should be
// shared between packages in application.
type Environment struct {
	Middleware alice.Chain
	Router     *Router
	DB         *gorm.DB
}

// NewEnvironment is a default constructor for an Environment.
// It takes logger type as an argument.
func NewEnvironment(logger logger.Type, db *gorm.DB) *Environment {
	return &Environment{
		Router: NewRouter(),
		Middleware: alice.New(
			Header("Access-Control-Allow-Origin", "*"),
			Header("Content-Type", "application/json"),
			Log(logger),
		),
		DB: db,
	}
}
