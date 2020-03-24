package main

import (
	"sync"

	"github.com/mjakobczyk/daily-content/config"
	"github.com/mjakobczyk/daily-content/env"
	"github.com/mjakobczyk/daily-content/internal/article"
	"github.com/mjakobczyk/daily-content/internal/db"
	"github.com/mjakobczyk/daily-content/internal/server"

	"github.com/vrischmann/envconfig"
)

func main() {
	config := config.Config{}
	err := envconfig.Init(&config)
	panicOnError(err)

	db, err := db.NewDatabase(&config.DB)
	panicOnError(err)
	db.Info()

	env := env.NewEnvironment(config.Server.Logger.Type, db.DB)

	articleRepository := article.NewRepository()
	articleService := article.NewService(env, articleRepository)
	articleRouter := article.NewRouter(env, articleService)
	articleRouter.InitRoutes()

	srv := server.NewServer(&config.Server, env)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		if err := srv.Start(); err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
