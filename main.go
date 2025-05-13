package main

import (
	"github.com/Marcus-Nastasi/gopportunities/config"
	"github.com/Marcus-Nastasi/gopportunities/handler"
	"github.com/Marcus-Nastasi/gopportunities/repository"
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

	// Initializing repository
	repo := repository.NewRepo(&db)

	// Initializing usecases
	getOpportunities := usecases.NewGetOpportunitiesUsecase(&repo)
	getOpportunitie := usecases.NewGetOpportunitieUsecase(&repo)
	createOpportunitie := usecases.NewCreateOpportunitieUsecase(&repo)
	updateOpportunitie := usecases.NewUpdateOpportunitieUsecase(&repo)
	deleteOpportunitie := usecases.NewDeleteOpportunitieUsecase(&repo)

	// Initializing Handler
	handler := handler.NewHandler(getOpportunities, getOpportunitie, createOpportunitie, updateOpportunitie, deleteOpportunitie)

	router.Initialize(handler)
}
