package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Content of Campaign X"
	contacts = []string{"johndoe@gmail.com", "fizzbuzz@gmail.com"}
	now      = time.Now().Add(-time.Minute)
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

func TestNewCampaignMustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required", err.Error())
}

func TestNewCampaignMustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("content is required", err.Error())
}

func TestNewCampaignMustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("list of contacts is required", err.Error())
}
