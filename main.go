package main

import (
	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/handler"
	"github.com/Marcus-Nastasi/gopportunities/router"
	"github.com/Marcus-Nastasi/gopportunities/usecases"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	db, err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err.Error())
		panic(err)
	}

	// Initializing usecases
	getOpportunities := usecases.NewGetOpportunitieUsecase(&db)
	// Initializing Handler
	handler := handler.NewHandler(getOpportunities)

	router.Initialize(handler)
}
