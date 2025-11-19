package campaign

import (
	internalerrors "campaign-manager/internal/internalErrors"
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID          string    `validate:"required"`
	Name        string    `validate:"min=5,max=24"`
	Description string    `validate:"min=5,max=1024"`
	CreatedAt   time.Time `validate:"required"`
	Contacts    []Contact `validate:"min=1,dive"`
}

func CreateCampaign(name string, description string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i] = Contact{Email: email}
	}

	campaign := &Campaign{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		Contacts:    contacts,
	}

	err := internalerrors.ValidateDomain(campaign)

	if err != nil {
		return nil, err
	}
	return campaign, nil
}
