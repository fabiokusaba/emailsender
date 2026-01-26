package endpoints

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fabiokusaba/emailsender/internal/test/internalmock"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_SendCampaignsEmail_Returns_200(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServiceMock)
	service.On("SendEmail", mock.MatchedBy(func(id string) bool {
		return id == "1"
	})).Return(nil)

	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("PATCH", "/", nil)
	rr := httptest.NewRecorder()

	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

	_, status, err := handler.SendCampaignsEmail(rr, req)

	assert.Equal(200, status)
	assert.Nil(err)
}

func Test_SendCampaignsEmail_Returns_Err(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServiceMock)
	service.On("SendEmail", mock.Anything).Return(errors.New("something went wrong"))

	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("PATCH", "/", nil)
	rr := httptest.NewRecorder()

	_, _, err := handler.SendCampaignsEmail(rr, req)

	assert.NotNil(err)
	assert.Equal("something went wrong", err.Error())
}
