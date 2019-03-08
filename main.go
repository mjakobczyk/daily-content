package main

import (
	"fmt"
	"os"

	"github.com/mjakobczyk/daily-content/web"
	"github.com/vrischmann/envconfig"
)

func main() {
	var config Config
	var err error

	err = envconfig.Init(&config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Config: ", config)

	srv := web.NewServer(&config.Server)
	err = srv.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
