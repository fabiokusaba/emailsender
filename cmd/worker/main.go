package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fabiokusaba/emailsender/internal/domain/campaign"
	"github.com/fabiokusaba/emailsender/internal/infrastructure/database"
	"github.com/fabiokusaba/emailsender/internal/infrastructure/mail"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Starting worker...")

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDb()
	if db == nil {
		log.Fatal("Error connecting to database")
	}

	repository := database.CampaignRepository{Db: db}
	service := campaign.ServiceImpl{
		Repository: &repository,
		SendMail:   mail.Send,
	}

	for {
		campaigns, err := repository.GetCampaignsToBeSent()
		if err != nil {
			log.Fatal("Error getting campaigns to be sent")
		}

		fmt.Printf("Found %d campaigns to be sent\n", len(campaigns))

		for _, campaign := range campaigns {
			service.Start(&campaign)
			fmt.Println("Sending campaign", campaign.ID)
		}

		time.Sleep(10 * time.Second)
	}
}
