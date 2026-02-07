package database

import (
	"github.com/fabiokusaba/emailsender/internal/domain/campaign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (cr *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := cr.Db.Create(campaign)
	return tx.Error
}

func (cr *CampaignRepository) Update(campaign *campaign.Campaign) error {
	tx := cr.Db.Save(campaign)
	return tx.Error
}

func (cr *CampaignRepository) GetAll() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := cr.Db.Find(&campaigns)
	return campaigns, tx.Error
}

func (cr *CampaignRepository) GetById(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	tx := cr.Db.Preload("Contacts").First(&campaign, "id = ?", id)
	return &campaign, tx.Error
}

func (cr *CampaignRepository) Delete(campaign *campaign.Campaign) error {
	tx := cr.Db.Select("Contacts").Delete(campaign)
	return tx.Error
}

func (cr *CampaignRepository) GetCampaignsToBeSent() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := cr.Db.Preload("Contacts").
		Where("status = ? and date_part('minute', now()::timestamp - updated_on::timestamp) > ?", campaign.Started, 1).
		Find(&campaigns)
	return campaigns, tx.Error
}
