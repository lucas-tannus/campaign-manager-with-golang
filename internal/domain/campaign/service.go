package campaign

import "campaign-manager/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := CreateCampaign(newCampaign.Name, newCampaign.Description, newCampaign.Emails)

	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", err
	}

	return campaign.ID, nil
}
