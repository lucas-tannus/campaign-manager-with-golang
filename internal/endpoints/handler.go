package endpoints

import "campaign-manager/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
