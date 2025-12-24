package endpoints

import (
	"errors"
	"net/http"

	"github.com/fabiokusaba/emailsender/internal/contract"
	"github.com/fabiokusaba/emailsender/internal/internalerrors"
	"github.com/go-chi/render"
)

func (h *Handler) PostCampaign(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaign

	err := render.DecodeJSON(r.Body, &request)
	if err != nil {
		render.Status(r, 500)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	id, err := h.CampaignService.Create(request)
	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, 500)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}

		render.Status(r, 400)
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}
