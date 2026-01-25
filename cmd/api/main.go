package main

import (
	"log"
	"net/http"

	"github.com/fabiokusaba/emailsender/internal/domain/campaign"
	"github.com/fabiokusaba/emailsender/internal/endpoints"
	"github.com/fabiokusaba/emailsender/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb()
	campaignService := campaign.ServiceImpl{Repository: &database.CampaignRepository{Db: db}}
	handler := endpoints.Handler{CampaignService: &campaignService}

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, 200)
		render.JSON(w, r, map[string]string{"health": "OK!"})
	})

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		r.Post("/", endpoints.HandlerError(handler.PostCampaign))
		r.Get("/{id}", endpoints.HandlerError(handler.GetCampaignById))
		r.Patch("/cancel/{id}", endpoints.HandlerError(handler.CampaignsCancelPatch))
		r.Delete("/delete/{id}", endpoints.HandlerError(handler.DeleteCampaign))
	})

	http.ListenAndServe(":3000", r)
}
