package main

import (
	"fmt"
	"os"

	"github.com/mjakobczyk/daily-content/config"
	"github.com/mjakobczyk/daily-content/internal/newsapi"
	"github.com/mjakobczyk/daily-content/internal/server"

	"github.com/vrischmann/envconfig"
)

func main() {
	var config config.Config
	var err error

	err = envconfig.Init(&config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Config: ", config)

	newsapiService := newsapi.NewService(&config.NewsAPI)

	srv := server.NewServer(&config.Server, newsapiService)
	err = srv.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
