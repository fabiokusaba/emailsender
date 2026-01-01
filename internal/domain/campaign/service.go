package campaign

import (
	"github.com/fabiokusaba/emailsender/internal/contract"
	"github.com/fabiokusaba/emailsender/internal/internalerrors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetById(id string) (*contract.CompaignResponse, error)
}

type ServiceImpl struct {
	Repository Repository
}

func (s *ServiceImpl) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImpl) GetById(id string) (*contract.CompaignResponse, error) {
	campaign, err := s.Repository.GetById(id)
	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	return &contract.CompaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}
