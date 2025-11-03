package flows

import (
	adapters_akua "akua-project/internal/adapters/akua"
	"context"
	"fmt"
	"log"
)

func InitializePaymentFlow(ctx context.Context) error {
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

func PaymentFlow(ctx context.Context) error {
	fmt.Println("Payment Flow")
	return nil
}
