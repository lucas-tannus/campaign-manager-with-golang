package endpoints

import (
	"campaign-manager/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var newCampaign contract.NewCampaign
	render.DecodeJSON(r.Body, &newCampaign)

	id, err := h.CampaignService.Create(newCampaign)

	return map[string]string{
		"id": id,
	}, 201, err
}

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaigns, err := h.CampaignService.FindAll()
	return campaigns, 200, err
}
