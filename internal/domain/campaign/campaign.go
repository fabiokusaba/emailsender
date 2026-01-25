package campaign

import (
	"time"

	"github.com/fabiokusaba/emailsender/internal/infrastructure/validator"
	"github.com/rs/xid"
)

const (
	Pending  string = "Pending"
	Started  string = "Started"
	Canceled string = "Canceled"
	Deleted  string = "Deleted"
	Done     string = "Done"
)

type Contact struct {
	ID         string `gorm:"size:50;primaryKey"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"primaryKey;size:50"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required" gorm:"type:timestamp with time zone"`
	Content   string    `validate:"min=5,max=1024" gorm:"type:text"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20"`
	CreatedBy string    `validate:"required,email" gorm:"size:100"`
}

func NewCampaign(name, content, createdBy string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].ID = xid.New().String()
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
		CreatedBy: createdBy,
	}

	err := validator.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}

func (c *Campaign) Done() {
	c.Status = Done
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Delete() {
	c.Status = Deleted
}
