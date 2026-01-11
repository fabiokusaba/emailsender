package campaign_test

import (
	"errors"
	"testing"

	"github.com/fabiokusaba/emailsender/internal/contract"
	"github.com/fabiokusaba/emailsender/internal/domain/campaign"
	"github.com/fabiokusaba/emailsender/internal/internalerrors"
	"github.com/fabiokusaba/emailsender/internal/test/internalmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi!",
		Emails:  []string{"teste1@test.com"},
	}
	service = campaign.ServiceImpl{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)

	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{})

	assert.False(errors.Is(err, internalerrors.ErrInternal))
}

func Test_Create_SaveCampaign(t *testing.T) {
	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		} else if campaign.Content != newCampaign.Content {
			return false
		} else if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}
		return true
	})).Return(nil)

	service.Repository = repositoryMock

	service.Create(newCampaign)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)

	repositoryMock.On("Save", mock.Anything).Return(errors.New("error to save on database"))

	service.Repository = repositoryMock

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)

	service.Repository = repositoryMock

	campaignReturned, _ := service.GetById(campaign.ID)

	assert.Equal(campaign.ID, campaignReturned.ID)
	assert.Equal(campaign.Name, campaignReturned.Name)
	assert.Equal(campaign.Content, campaignReturned.Content)
	assert.Equal(campaign.Status, campaignReturned.Status)
}

func Test_GetById_ReturnErrorWhenSomethingWentWrong(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetById", mock.Anything).Return(nil, errors.New("Something went wrong"))

	service.Repository = repositoryMock

	_, err := service.GetById("123")

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnRecordNotFound_When_Campaign_Does_Not_Exists(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	service.Repository = repositoryMock

	err := service.Delete("123")

	assert.Equal(gorm.ErrRecordNotFound, err)
}

func Test_Delete_ReturnStatusInvalid_When_Campaign_Is_Not_Pending(t *testing.T) {
	assert := assert.New(t)
	campaign := campaign.Campaign{
		ID:     "1",
		Status: campaign.Started,
	}

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetById", mock.Anything).Return(&campaign, nil)

	service.Repository = repositoryMock

	err := service.Delete(campaign.ID)

	assert.Equal(errors.New("campaign is not available to delete"), err)
}

func Test_Delete_ReturnInternalError_When_Something_Went_Wrong(t *testing.T) {
	assert := assert.New(t)
	campaignMock := campaign.Campaign{
		ID:     "1",
		Status: campaign.Pending,
	}

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetById", mock.Anything).Return(&campaignMock, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaign.ID == campaignMock.ID
	})).Return(errors.New("something went wrong"))

	service.Repository = repositoryMock

	err := service.Delete(campaignMock.ID)

	assert.Equal(internalerrors.ErrInternal, err)
}

func Test_Delete_ReturnNil_When_Campaign_Is_Deleted(t *testing.T) {
	assert := assert.New(t)
	campaignMock := campaign.Campaign{
		ID:     "1",
		Status: campaign.Pending,
	}

	repositoryMock := new(internalmock.CampaignRepositoryMock)
	repositoryMock.On("GetById", mock.Anything).Return(&campaignMock, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaign.ID == campaignMock.ID
	})).Return(nil)

	service.Repository = repositoryMock

	err := service.Delete(campaignMock.ID)

	assert.Nil(err)
}
