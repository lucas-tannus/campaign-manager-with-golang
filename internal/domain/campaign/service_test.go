package campaign

import (
	"campaign-manager/internal/contract"
	internalerrors "campaign-manager/internal/internalErrors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func (r *repositoryMock) Get() ([]Campaign, error) {
	return nil, nil
}

var (
	newCampaign = contract.NewCampaign{
		Name:        "New Campaign",
		Description: "This is a test campaign",
		Emails: []string{
			"test@test.com",
		},
	}
	service = ServiceImp{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	mockedRepository := new(repositoryMock)
	mockedRepository.On("Save", mock.Anything).Return(nil)
	service.Repository = mockedRepository

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_Campaign_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaign{
		Name:        "",
		Description: "This is a test campaign",
		Emails: []string{
			"test@test.com",
		},
	})

	assert.NotNil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	mockedRepository := new(repositoryMock)
	mockedRepository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name ||
			campaign.Description != newCampaign.Description ||
			len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)
	service.Repository = mockedRepository

	service.Create(newCampaign)

	mockedRepository.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	mockedRepository := new(repositoryMock)
	mockedRepository.On("Save", mock.Anything).Return(errors.New("error"))
	service.Repository = mockedRepository

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(internalerrors.ErrInternalError, err))
}
