package endpoints

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiokusaba/emailsender/internal/contract"
	"github.com/fabiokusaba/emailsender/internal/test/internalmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignGetById_Should_Return_Campaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.CompaignResponse{
		ID:      "123",
		Name:    "test campaign",
		Content: "test campaign content",
		Status:  "Pending",
	}

	service := new(internalmock.CampaignServiceMock)
	service.On("GetById", mock.Anything).Return(&campaign, nil)

	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.GetCampaignById(rr, req)

	assert.Equal(http.StatusOK, status)
	assert.Equal(campaign.ID, response.(*contract.CompaignResponse).ID)
	assert.Equal(campaign.Name, response.(*contract.CompaignResponse).Name)
	assert.Equal(campaign.Content, response.(*contract.CompaignResponse).Content)
	assert.Equal(campaign.Status, response.(*contract.CompaignResponse).Status)
}

func Test_CampaignGetById_Should_Return_Error_When_Something_Went_Wrong(t *testing.T) {
	assert := assert.New(t)
	errExpected := errors.New("Something went wrong")

	service := new(internalmock.CampaignServiceMock)
	service.On("GetById", mock.Anything).Return(nil, errExpected)

	handler := Handler{CampaignService: service}

	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, err := handler.GetCampaignById(rr, req)

	assert.Equal(errExpected.Error(), err.Error())
}
