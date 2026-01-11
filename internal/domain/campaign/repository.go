package campaign

type Repository interface {
	Save(campaign *Campaign) error
	Update(campaign *Campaign) error
	GetAll() ([]Campaign, error)
	GetById(id string) (*Campaign, error)
	Delete(campaign *Campaign) error
}
