package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	assert := assert.New(t)

	name := "Campaign X"
	content := "Content of Campaign X"
	contacts := []string{"johndoe@gmail.com", "fizzbuzz@gmail.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.NotNil(campaign.CreatedOn)
	assert.Greater(campaign.CreatedOn, now)
}
