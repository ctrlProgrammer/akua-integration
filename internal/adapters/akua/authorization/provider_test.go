package adapters_akua_authorization

import (
	adapters_akua "akua-project/internal/adapters/akua"
	"akua-project/internal/instruments"
	"context"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Setup() (*adapters_akua.Client, *AuthorizationProvider, error) {
	envPath := filepath.Join("..", "..", "..", "..", ".env")

	err := godotenv.Load(envPath)

	if err != nil {
		return nil, nil, err
	}

	akuaClient, err := adapters_akua.NewClient()

	if err != nil {
		return nil, nil, err
	}

	err = akuaClient.LoadJwtToken()

	if err != nil {
		return nil, nil, err
	}

	return akuaClient, NewAuthorizationProvider(), nil
}

// With this we will create an authorization with the automatic capture method
// Will be approved and captured automatically
func Test_Authorize_AutomaticCapture_Success(t *testing.T) {
	akuaClient, provider, err := Setup()

	if err != nil {
		t.Fatal(err)
	}

	request := AuthorizeRequest{
		Amount: instruments.AmountObject{
			Currency: "USD",
			Value:    100,
		},
		Intent:     INTENT_AUTHORIZE,
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

	authorization, err := provider.Authorize(context.Background(), akuaClient, request)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, authorization)
	assert.NotNil(t, authorization.PaymentID)
}

// With this we will create an authorization with the automatic capture method
// Will be approved and captured automatically
func Test_PreAuthorize_Success(t *testing.T) {
	akuaClient, provider, err := Setup()

	if err != nil {
		t.Fatal(err)
	}

	request := AuthorizeRequest{
		Amount: instruments.AmountObject{
			Currency: "USD",
			Value:    20.00,
		},
		Intent: INTENT_PREAUTHORIZE,
		Capture: instruments.CaptureObject{
			Mode: instruments.CAPTURE_MODE_AUTOMATIC,
		},
		MerchantId: akuaClient.GetMerchantId(),
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

	authorization, err := provider.Authorize(context.Background(), akuaClient, request)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, authorization)
	assert.NotNil(t, authorization.PaymentID)
}

// With this we will create an authorization with the manual cature method
// Will be approved but not captured automatically
func Test_Authorize_ManualCapture_Success(t *testing.T) {
	akuaClient, provider, err := Setup()

	if err != nil {
		t.Fatal(err)
	}

	request := AuthorizeRequest{
		Amount: instruments.AmountObject{
			Currency: "USD",
			Value:    20.00,
		},
		Intent: INTENT_AUTHORIZE,
		Capture: instruments.CaptureObject{
			Mode: instruments.CAPTURE_MODE_MANUAL,
		},
		MerchantId: akuaClient.GetMerchantId(),
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

	authorization, err := provider.Authorize(context.Background(), akuaClient, request)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, authorization)
	assert.NotNil(t, authorization.PaymentID)
}
