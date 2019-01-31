package main

import (
	"fmt"
	"os"

	"github.com/vrischmann/envconfig"
)

func main() {
	var config Config

	err := envconfig.Init(&config)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Config:\n", config)
}
