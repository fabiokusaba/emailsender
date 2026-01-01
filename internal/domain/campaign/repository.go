package campaign

type Repository interface {
	Save(campaign *Campaign) error
	GetAll() ([]Campaign, error)
	GetById(id string) (*Campaign, error)
}
