package endpoints

import "github.com/fabiokusaba/emailsender/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
