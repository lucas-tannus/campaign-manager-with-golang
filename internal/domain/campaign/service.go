package campaign

import (
	"campaign-manager/internal/contract"
	internalerrors "campaign-manager/internal/internalErrors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	FindAll() ([]Campaign, error)
	FindByUuid(campaignUuid string) (*contract.CampaignResponse, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := CreateCampaign(newCampaign.Name, newCampaign.Description, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerrors.ErrInternalError
	}

	return campaign.ID, nil
}

func (s *ServiceImp) FindAll() ([]Campaign, error) {
	res, err := s.Repository.Get()

	if err != nil {
		return nil, internalerrors.ErrInternalError
	}

	return res, nil
}

func (s *ServiceImp) FindByUuid(campaignUuid string) (*contract.CampaignResponse, error) {
	res, err := s.Repository.GetByUuid(campaignUuid)

	if err != nil {
		return nil, internalerrors.ErrInternalError
	}

	if res == nil {
		return (*contract.CampaignResponse)(nil), internalerrors.ErrResourceNotFound
	}

	return &contract.CampaignResponse{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Status:      res.Status,
	}, nil
}
