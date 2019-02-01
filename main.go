package main

import (
	"fmt"
	"os"

	"github.com/vrischmann/envconfig"
)

func main() {
	var config Config
	err := envconfig.Init(&config)
	fmt.Println("Config: ", config)

	// err = srv.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
