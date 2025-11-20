package database

import "campaign-manager/internal/domain/campaign"

type CampaignRepository struct {
	Campaigns []campaign.Campaign
}

func (c *CampaignRepository) Save(campaign *campaign.Campaign) error {
	c.Campaigns = append(c.Campaigns, *campaign)
	return nil
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	return c.Campaigns, nil
}

func (c *CampaignRepository) GetByUuid(uuid string) (*campaign.Campaign, error) {
	var response *campaign.Campaign

	for i := 0; i < len(c.Campaigns); i++ {
		if c.Campaigns[i].ID == uuid {
			response = &c.Campaigns[i]
		}
	}

	return response, nil
}
