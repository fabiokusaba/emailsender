package campaign

import (
	"errors"

	"github.com/fabiokusaba/emailsender/internal/contract"
	"github.com/fabiokusaba/emailsender/internal/internalerrors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetById(id string) (*contract.CompaignResponse, error)
	Cancel(id string) error
	Delete(id string) error
	SendEmail(id string) error
}

type ServiceImpl struct {
	Repository Repository
	SendMail   func(campaign *Campaign) error
}

func (s *ServiceImpl) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.CreatedBy, newCampaign.Emails)
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
		return nil, internalerrors.ProcessErrorToReturn(err)
	}

	return &contract.CompaignResponse{
		ID:                   campaign.ID,
		Name:                 campaign.Name,
		Content:              campaign.Content,
		Status:               campaign.Status,
		AmountOfEmailsToSend: len(campaign.Contacts),
		CreatedBy:            campaign.CreatedBy,
	}, nil
}

func (s *ServiceImpl) Cancel(id string) error {
	campaign, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return errors.New("campaign is not available to cancel")
	}

	campaign.Cancel()

	err = s.Repository.Update(campaign)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}

func (s *ServiceImpl) Delete(id string) error {
	campaign, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return errors.New("campaign is not available to delete")
	}

	campaign.Delete()

	err = s.Repository.Delete(campaign)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}

func (s *ServiceImpl) Start(campaignSaved *Campaign) {
	err := s.SendMail(campaignSaved)
	if err != nil {
		campaignSaved.Fail()
	} else {
		campaignSaved.Done()
	}

	s.Repository.Update(campaignSaved)
}

func (s *ServiceImpl) SendEmail(id string) error {
	campaignSaved, err := s.Repository.GetById(id)
	if err != nil {
		return internalerrors.ProcessErrorToReturn(err)
	}

	if campaignSaved.Status != Pending {
		return errors.New("campaign is not available to send email")
	}

	go s.Start(campaignSaved)

	campaignSaved.Started()

	err = s.Repository.Update(campaignSaved)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
