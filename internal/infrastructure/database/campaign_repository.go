package database

import (
	"github.com/fabiokusaba/emailsender/internal/domain/campaign"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (cr *CampaignRepository) Save(campaign *campaign.Campaign) error {
	cr.campaigns = append(cr.campaigns, *campaign)
	return nil
}

func (cr *CampaignRepository) GetAll() ([]campaign.Campaign, error) {
	return cr.campaigns, nil
}
