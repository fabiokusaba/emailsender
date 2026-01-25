package database

import (
	"os"

	"github.com/fabiokusaba/emailsender/internal/domain/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := os.Getenv("DATABASE")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})
	if err != nil {
		panic("failed to migrate models")
	}

	return db
}
