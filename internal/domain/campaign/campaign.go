package campaign

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	Contacts    []Contact
}

func CreateCampaign(name string, description string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("campaign name cannot be empty")
	} else if description == "" {
		return nil, errors.New("campaign description cannot be empty")
	} else if len(emails) == 0 {
		return nil, errors.New("at least one contact email is required")
	}

	contacts := make([]Contact, len(emails))
	for i, email := range emails {
		contacts[i] = Contact{Email: email}
	}

	return &Campaign{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		Contacts:    contacts,
	}, nil
}
