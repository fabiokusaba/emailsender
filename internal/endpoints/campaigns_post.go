package endpoints

import (
	"net/http"

	"github.com/fabiokusaba/emailsender/internal/contract"
	"github.com/go-chi/render"
)

func (h *Handler) PostCampaign(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign

	err := render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	return map[string]string{"id": id}, 201, err
}
