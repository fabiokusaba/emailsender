package internalmock

import (
	"github.com/fabiokusaba/emailsender/internal/contract"
	"github.com/stretchr/testify/mock"
)

type CampaignServiceMock struct {
	mock.Mock
}

func (s *CampaignServiceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := s.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (s *CampaignServiceMock) GetById(id string) (*contract.CompaignResponse, error) {
	args := s.Called(id)

	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*contract.CompaignResponse), nil
}

func (s *CampaignServiceMock) Cancel(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *CampaignServiceMock) Delete(id string) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *CampaignServiceMock) SendEmail(id string) error {
	args := s.Called(id)
	return args.Error(0)
}
