package campaign

import (
	"campaign-manager/internal/contract"
	internalerrors "campaign-manager/internal/internalErrors"
	"errors"
	"testing"

	"github.com/google/uuid"
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
	args := r.Called()
	return args.Get(0).([]Campaign), args.Error(1)
}

func (r *repositoryMock) GetByUuid(uuid string) (*Campaign, error) {
	args := r.Called(uuid)
	return args.Get(0).(*Campaign), args.Error(1)
}

var campaignUuid string = uuid.New().String()
var (
	newCampaign = contract.NewCampaign{
		Name:        "New Campaign",
		Description: "This is a test campaign",
		Emails: []string{
			"test@test.com",
		},
	}
	service   = ServiceImp{}
	campaigns = []Campaign{
		{
			ID:          campaignUuid,
			Name:        "New Campaign",
			Description: "This is a mocked campaign",
			Contacts: []Contact{
				{
					Email: "test@test.com",
				},
			},
		},
	}
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

func Test_FindAll_GetCampaigns(t *testing.T) {
	assert := assert.New(t)
	mockedRepository := new(repositoryMock)
	mockedRepository.On("Get", mock.Anything).Return(campaigns, nil)
	service.Repository = mockedRepository

	res, err := service.FindAll()

	assert.NotNil(res)
	assert.Nil(err)
}

func Test_FindAll_GetCampaigns_With_Error(t *testing.T) {
	assert := assert.New(t)
	mockedRepository := new(repositoryMock)
	mockedRepository.On("Get", mock.Anything).Return([]Campaign{}, errors.New("critical error"))
	service.Repository = mockedRepository

	res, err := service.FindAll()

	errors.Is(internalerrors.ErrInternalError, err)
	assert.Nil(res)
}

func Test_FindByUuid_GetByUuid(t *testing.T) {
	assert := assert.New(t)
	mockedRepository := new(repositoryMock)
	mockedRepository.On("GetByUuid", mock.MatchedBy(func(uuid string) bool {
		return uuid == campaignUuid
	})).Return(&campaigns[0], nil)
	service.Repository = mockedRepository

	res, err := service.FindByUuid(campaignUuid)

	assert.Equal(campaignUuid, res.ID)
	assert.Nil(err)
}

func Test_FindByUuid_GetByUuid_With_Error(t *testing.T) {
	assert := assert.New(t)
	mockedRepository := new(repositoryMock)
	mockedRepository.On("GetByUuid", mock.Anything).Return((*Campaign)(nil), errors.New("critical error"))
	service.Repository = mockedRepository

	res, err := service.FindByUuid(campaignUuid)

	assert.Nil(res)
	errors.Is(internalerrors.ErrInternalError, err)
}

func Test_FindByUuid_GetByUuid_Not_Found(t *testing.T) {
	assert := assert.New(t)
	mockedRepository := new(repositoryMock)
	mockedRepository.On("GetByUuid", mock.Anything).Return((*Campaign)(nil), nil)
	service.Repository = mockedRepository

	res, err := service.FindByUuid(campaignUuid)

	assert.Nil(res)
	errors.Is(internalerrors.ErrResourceNotFound, err)
}
