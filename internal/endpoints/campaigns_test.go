package endpoints

import (
	"bytes"
	"campaign-manager/internal/contract"
	"campaign-manager/internal/domain/campaign"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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

func (r *serviceMock) FindAll() ([]campaign.Campaign, error) {
	return nil, nil
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
