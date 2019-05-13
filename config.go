package main

import (
	"github.com/mjakobczyk/daily-content/newsapi"
	"github.com/mjakobczyk/daily-content/server"
)

// Config stores all the configuration used in project.
type Config struct {
	Server  server.Config
	NewsAPI newsapi.Config
}
