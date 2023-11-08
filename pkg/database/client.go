package database

import (
	"fmt"
	"log"
	"os"
	"thirsty-jotter/pkg/api/models"

	"github.com/joho/godotenv"
	xg "github.com/kerdokurs/xata-go"
)

// MyAPIClient represents
type MyAPIClient struct {
	*xg.Client
	Plants xg.Table[models.Plant]
}

// BuildXataClient creates a xata client
func BuildXataClient() *MyAPIClient {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file ")
	}

	xataAPIKey := os.Getenv("XATA_API_KEY")
	xataDatabaseURL := os.Getenv("XATA_DB_URL")
	xataClient := xg.NewClient(xataAPIKey, xataDatabaseURL)

	fmt.Println("Building the xata client...")

	client := &MyAPIClient{
		Client: xataClient,
		Plants: xg.NewTableImpl[models.Plant](xataClient, "Plants"),
	}

	return client
}
