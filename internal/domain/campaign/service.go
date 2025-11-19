package campaign

import (
	"campaign-manager/internal/contract"
	internalerrors "campaign-manager/internal/internalErrors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	FindAll() ([]Campaign, error)
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
	return s.Repository.Get()
}
