package endpoints

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiokusaba/emailsender/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) Create(newCampaign contract.NewCampaign) (string, error) {
	args := s.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	handler := Handler{CampaignService: service}

	body := contract.NewCampaign{
		Name:    "teste campaign",
		Content: "teste content campaign",
		Emails:  []string{"teste@email.com"},
	}

	service.On("Create", mock.MatchedBy(func(request contract.NewCampaign) bool {
		if request.Name == body.Name && request.Content == body.Content && len(request.Emails) == len(body.Emails) {
			return true
		}
		return false
	})).Return("IdQualquer", nil)

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	_, status, err := handler.PostCampaign(res, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)
}

func Test_CampaignsPost_should_inform_error_when_exists(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	handler := Handler{CampaignService: service}

	body := contract.NewCampaign{
		Name:    "teste campaign",
		Content: "teste content campaign",
		Emails:  []string{"teste@email.com"},
	}

	service.On("Create", mock.Anything).Return("", errors.New("testando error"))

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	_, _, err := handler.PostCampaign(res, req)

	assert.NotNil(err)
}
