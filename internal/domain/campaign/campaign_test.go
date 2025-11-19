package campaign

import (
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name        = "New Campaign"
	description = "This is a test campaign"
	emails      = []string{
		"test@test.com",
		"test2@test.com",
	}
	fake = faker.New()
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

func Test_CreateCampaign_NameValidateMinSize(t *testing.T) {
	assert := assert.New(t)

	campaign, err := CreateCampaign("", description, emails)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "name should have at least 5 characters")
}

func Test_CreateCampaign_NameValidateMaxSize(t *testing.T) {
	assert := assert.New(t)

	campaign, err := CreateCampaign(fake.Lorem().Text(30), description, emails)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "name should have less than 24 characters")
}

func Test_CreateCampaign_DescriptionValidateMinSize(t *testing.T) {
	assert := assert.New(t)

	campaign, err := CreateCampaign(name, "", emails)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "description should have at least 5 characters")
}

func Test_CreateCampaign_DescriptionValidateMaxSize(t *testing.T) {
	assert := assert.New(t)

	campaign, err := CreateCampaign(name, fake.Lorem().Text(1045), emails)

	assert.Nil(campaign)
	assert.Equal(err.Error(), "description should have less than 1024 characters")
}

func Test_CreateCampaign_EmailNotEmpty(t *testing.T) {
	assert := assert.New(t)

	campaign, err := CreateCampaign(name, description, []string{
		"test",
	})

	assert.Nil(campaign)
	assert.Equal(err.Error(), "email is invalid")
}
