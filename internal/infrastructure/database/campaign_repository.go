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
