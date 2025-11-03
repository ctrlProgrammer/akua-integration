package flows

import (
	adapters_akua "akua-project/internal/adapters/akua"
	"context"
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

func InitializeEnvVariables() error {
	envPath := filepath.Join("..", "..", "..", ".env")

	err := godotenv.Load(envPath)

	if err != nil {
		return err
	}

	return nil
}

func InitializePaymentFlow(ctx context.Context) error {
	InitializeEnvVariables()

	akuaClient, err := adapters_akua.NewClient()

	if err != nil {
		return err
	}

	err = akuaClient.LoadJwtToken()

	if err != nil {
		return err
	}

	log.Println("Starting payment flow initialization...")
	log.Println("Creating new Akua client...")

	if akuaClient == nil {
		log.Println("Failed to create Akua client.")
	} else {
		log.Println("Akua client created successfully.")
	}

	log.Println("Loading JWT token for Akua client...")

	if err != nil {
		log.Printf("Failed to load JWT token: %v\n", err)
	} else {
		log.Println("JWT token loaded successfully.")
	}

	return nil
}

func Test_Authorize_Auto_Capture_Success(t *testing.T) {
	log.Println("=================================================")
	log.Println("Testing Authorize with automatic capture...")
	log.Println("=================================================")

	err := InitializePaymentFlow(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	log.Println("=================================================")
	log.Println("Finish flow initialization...")
	log.Println("=================================================")

	fmt.Println("Payment Flow")
}
