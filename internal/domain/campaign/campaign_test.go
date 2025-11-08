package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name        = "New Campaign"
	description = "This is a test campaign"
	emails      = []string{
		"test@test.com",
		"test2@test.com",
	}
)

func Test_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := CreateCampaign(name, description, emails)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Description, description)
	assert.Equal(len(campaign.Contacts), len(emails))
}

func Test_CreateCampaign_IDNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := CreateCampaign(name, description, emails)

	assert.NotNil(campaign.ID)
}

func Test_CreateCampaign_CreatedAtNotNil(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := CreateCampaign(name, description, emails)

	assert.NotNil(campaign.CreatedAt)
}

func Test_CreateCampaign_NameIsRequired(t *testing.T) {
	assert := assert.New(t)

	campaign, err := CreateCampaign("", description, emails)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "campaign name cannot be empty")
}

func Test_CreateCampaign_DescriptionIsRequired(t *testing.T) {
	assert := assert.New(t)

	campaign, err := CreateCampaign(name, "", emails)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "campaign description cannot be empty")
}

func Test_CreateCampaign_EmailNotEmpty(t *testing.T) {
	assert := assert.New(t)

	campaign, err := CreateCampaign(name, description, []string{})

	assert.Nil(campaign)
	assert.Equal(err.Error(), "at least one contact email is required")
}
