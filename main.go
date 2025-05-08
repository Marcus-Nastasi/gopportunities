package main

import (
	"log"

	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/router"
)

func main() {

	err := config.Init()
	if err != nil {
		defer log.Fatal(err)
		panic(err)
	}

	router.Initialize()
}
