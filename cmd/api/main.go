package main

import (
	"net/http"

	"github.com/fabiokusaba/emailsender/internal/domain/campaign"
	"github.com/fabiokusaba/emailsender/internal/endpoints"
	"github.com/fabiokusaba/emailsender/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.ServiceImpl{Repository: &database.CampaignRepository{}}
	handler := endpoints.Handler{CampaignService: &campaignService}

	r.Post("/campaigns", endpoints.HandlerError(handler.PostCampaign))
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.GetCampaignById))

	http.ListenAndServe(":3000", r)
}
