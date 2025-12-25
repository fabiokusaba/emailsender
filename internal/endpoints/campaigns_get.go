package endpoints

import (
	"net/http"
)

func (h *Handler) GetCampaigns(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.Repository.GetAll()

	return campaigns, 200, err
}
