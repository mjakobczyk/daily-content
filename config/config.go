package config

import (
	"github.com/mjakobczyk/daily-content/internal/db"
	"github.com/mjakobczyk/daily-content/internal/newsapi"
	"github.com/mjakobczyk/daily-content/internal/server"
)

// Config stores all the configuration used in project.
type Config struct {
	Server  server.Config
	NewsAPI newsapi.Config
	DB      db.Config
}
