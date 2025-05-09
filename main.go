package main

import (
	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()
	if err != nil {
		defer logger.Errorf("config initialization error: %v", err.Error())
		panic(err)
	}

	router.Initialize()
}
