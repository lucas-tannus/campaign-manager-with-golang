package endpoints

import (
	"bytes"
	"campaign-manager/internal/contract"
	"campaign-manager/internal/domain/campaign"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
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

func (s *serviceMock) FindAll() ([]campaign.Campaign, error) {
	args := s.Called()
	return args.Get(0).([]campaign.Campaign), args.Error(1)
}

func (s *serviceMock) FindByUuid(campaignUuid string) (*contract.CampaignResponse, error) {
	args := s.Called(campaignUuid)
	return args.Get(0).(*contract.CampaignResponse), args.Error(1)
}

var (
	campaignId = uuid.New().String()
	body       = contract.NewCampaign{
		Name:        "Campaign Test",
		Description: "This is a campaign test",
		Emails: []string{
			"test@test.com",
		},
	}
	campaigns = []campaign.Campaign{
		{
			ID:          uuid.New().String(),
			Name:        "Campaign Test",
			Description: "This is a campaign test",
			CreatedAt:   time.Now(),
			Contacts: []campaign.Contact{
				{
					Email: "test@test.com",
				},
			},
		},
	}
)

func Test_CampaignPost_With_Success(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	handler := Handler{
		CampaignService: service,
	}
	service.On("Create", mock.MatchedBy(func(newCampaign contract.NewCampaign) bool {
		if newCampaign.Name == body.Name && newCampaign.Description == body.Description {
			return true
		} else {
			return false
		}
	})).Return(campaignId, nil)
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/", &buf)
	res := httptest.NewRecorder()

	obj, status, err := handler.CampaignPost(res, req)

	assert.Equal(status, 201)
	assert.Nil(err)
	assert.NotNil(obj)
}

func Test_CampaignGet_With_Success(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	handler := Handler{
		CampaignService: service,
	}
	service.On("FindAll", mock.Anything).Return(campaigns, nil)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	obj, status, err := handler.CampaignGet(res, req)

	assert.Equal(status, 200)
	assert.Nil(err)
	assert.NotNil(obj)
}

func Test_CampaignGetByUuid_With_Success(t *testing.T) {
	assert := assert.New(t)
	service := new(serviceMock)
	handler := Handler{
		CampaignService: service,
	}
	response := &contract.CampaignResponse{
		ID:          campaigns[0].ID,
		Name:        campaigns[0].Name,
		Description: campaigns[0].Description,
		Status:      campaigns[0].Status,
	}
	service.On("FindByUuid", mock.Anything).Return(response, nil)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	obj, status, err := handler.CampaignGetByUuid(res, req)

	assert.Equal(status, 200)
	assert.Nil(err)
	assert.NotNil(obj)
}
