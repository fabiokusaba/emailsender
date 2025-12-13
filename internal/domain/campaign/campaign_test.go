package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Content of Campaign X"
	contacts = []string{"johndoe@gmail.com", "fizzbuzz@gmail.com"}
	now      = time.Now().Add(-time.Minute)
	fake     = faker.New()
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.NotNil(campaign.CreatedOn)
	assert.Greater(campaign.CreatedOn, now)
}

func TestNewCampaignMustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name must be at least 5", err.Error())
}

func TestNewCampaignMustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name must be at most 24", err.Error())
}

func TestNewCampaignMustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content must be at least 5", err.Error())
}

func TestNewCampaignMustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1030), contacts)

	assert.Equal("content must be at most 1024", err.Error())
}

func TestNewCampaignMustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil)

	assert.Equal("contacts must be at least 1", err.Error())
}

func TestNewCampaignMustValidateContactsEmail(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("email is invalid", err.Error())
}
