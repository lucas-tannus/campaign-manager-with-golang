package campaign

type Repository interface {
	Save(campaign *Campaign) error
	Get() ([]Campaign, error)
	GetByUuid(uuid string) (*Campaign, error)
}
