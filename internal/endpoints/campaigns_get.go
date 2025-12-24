package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) GetCampaigns(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 200)
	render.JSON(w, r, h.CampaignService.Repository.GetAll())
}
