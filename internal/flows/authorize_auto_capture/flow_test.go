package flows

import (
	adapters_akua "akua-project/internal/adapters/akua"
	adapters_akua_authorization "akua-project/internal/adapters/akua/authorization"
	"akua-project/internal/instruments"
	"context"
	"log"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func InitializeEnvVariables() error {
	envPath := filepath.Join("..", "..", "..", ".env")

	err := godotenv.Load(envPath)

	if err != nil {
		return err
	}

	return nil
}

func InitializePaymentFlow(ctx context.Context) (*adapters_akua.Client, *adapters_akua_authorization.AuthorizationProvider, error) {
	InitializeEnvVariables()

	akuaClient, err := adapters_akua.NewClient()

	if err != nil {
		return nil, nil, err
	}

	err = akuaClient.LoadJwtToken()

	if err != nil {
		return nil, nil, err
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

	return akuaClient, adapters_akua_authorization.NewAuthorizationProvider(), nil
}

func Test_Authorize_AutoCapture_Success(t *testing.T) {
	log.Println("=================================================")
	log.Println("Testing Authorize with automatic capture...")
	log.Println("=================================================")

	akuaClient, authorizationProvider, err := InitializePaymentFlow(context.Background())

	if err != nil {
		t.Fatal(err)
	}

	log.Println("=================================================")
	log.Println("Finish flow initialization...")
	log.Println("=================================================")

	log.Println("=================================================")
	log.Println("Start authorization...")
	log.Println("=================================================")

	request := adapters_akua_authorization.AuthorizeRequest{
		Amount: instruments.AmountObject{
			Currency: "USD",
			Value:    100,
		},
		Intent:     adapters_akua_authorization.INTENT_AUTHORIZE,
		MerchantId: akuaClient.GetMerchantId(),
		Capture: instruments.CaptureObject{
			Mode: instruments.CAPTURE_MODE_AUTOMATIC,
		},
		Instrument: instruments.InstrumentObject{
			Type: "CARD",
			Card: instruments.Instrument{
				Number:          instruments.CARD_APPROVED_CREDIT,
				CVV:             "123",
				ExpirationMonth: "12",
				ExpirationYear:  "26",
				HolderName:      "John Doe",
			},
		},
	}

	log.Println("=================================================")
	log.Println("Authorization Request:")
	log.Println("Amount: ", request.Amount)
	log.Println("Intent: ", request.Intent)
	log.Println("MerchantId: ", request.MerchantId)
	log.Println("Capture Mode: ", request.Capture.Mode)
	log.Println("Instrument: ", request.Instrument)
	log.Println("=================================================")

	authorization, err := authorizationProvider.Authorize(context.Background(), akuaClient, request)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("=================================================")
	log.Println("Authorization Response:")
	log.Println("Authorization ID: ", authorization.PaymentID)
	log.Println("Authorization Status: ", authorization.Transaction.Status)
	log.Println("Authorization Amount: ", authorization.Transaction.Amount)
	log.Println("Authorization Transaction ID: ", authorization.Transaction.ID)
	log.Println("Authorization Transaction Status: ", authorization.Transaction.Status)
	log.Println("=================================================")

	assert.NotNil(t, authorization)
	assert.NotNil(t, authorization.PaymentID)
	assert.Equal(t, "APPROVED", authorization.Transaction.Status, "expected transaction status to be APPROVED")
}
