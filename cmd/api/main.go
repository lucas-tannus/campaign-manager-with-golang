package main

import (
	"campaign-manager/internal/domain/campaign"
	"campaign-manager/internal/endpoints"
	"campaign-manager/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	campaignService := &campaign.ServiceImp{
		Repository: &database.CampaignRepository{},
	}
	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	router.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	router.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))

	http.ListenAndServe(":3000", router)
}
