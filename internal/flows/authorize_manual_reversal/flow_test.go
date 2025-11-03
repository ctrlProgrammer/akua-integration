package flows

import (
	adapters_akua "akua-project/internal/adapters/akua"
	adapters_akua_authorization "akua-project/internal/adapters/akua/authorization"
	adapters_akua_payments "akua-project/internal/adapters/akua/payments"
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

func InitializePaymentFlow(ctx context.Context) (*adapters_akua.Client, *adapters_akua_authorization.AuthorizationProvider, *adapters_akua_payments.PaymentsProvider, error) {
	InitializeEnvVariables()

	akuaClient, err := adapters_akua.NewClient()

	if err != nil {
		return nil, nil, nil, err
	}

	err = akuaClient.LoadJwtToken()

	if err != nil {
		return nil, nil, nil, err
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

	return akuaClient, adapters_akua_authorization.NewAuthorizationProvider(), adapters_akua_payments.NewPaymentsProvider(), nil
}

func Test_Authorize_ManualCapture_Success(t *testing.T) {
	log.Println("=================================================")
	log.Println("Testing Authorize with manual capture, the payment will be captured manually...")
	log.Println("=================================================")

	akuaClient, authorizationProvider, paymentsProvider, err := InitializePaymentFlow(context.Background())

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
			Mode: instruments.CAPTURE_MODE_MANUAL,
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

	log.Println("=================================================")
	log.Println("Getting payment state...")
	log.Println("=================================================")

	payment, err := paymentsProvider.GetPaymentById(context.Background(), akuaClient, authorization.PaymentID)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("=================================================")
	log.Println("Payment State:")
	log.Println("The payment must be in manual capture mode")
	log.Println("Payment ID: ", payment.ID)
	log.Println("Payment Status: ", payment.Status)
	log.Println("Payment Amount: ", payment.CurrentAmount)
	log.Println("Capture Mode: ", payment.Capture.Mode)
	log.Println("Payment Transactions: ", len(payment.Transactions))
	for _, transaction := range payment.Transactions {
		log.Println("=================================================")
		log.Println("Transaction ID: ", transaction.ID)
		log.Println("Transaction Type: ", transaction.Type)
		log.Println("Transaction Status: ", transaction.Status)
		log.Println("Transaction Amount: ", transaction.Amount)
		log.Println("=================================================")
	}
	log.Println("=================================================")

	assert.Equal(t, "MANUAL", payment.Capture.Mode, "expected capture mode to be MANUAL")
	assert.Equal(t, "AUTHORIZED", payment.Status, "expected payment status to be AUTHORIZED")

	log.Println("=================================================")
	log.Println("Initializing reversal...")
	log.Println("=================================================")

	reversalResponse, err := authorizationProvider.Reversal(context.Background(), akuaClient, authorization.PaymentID)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("=================================================")
	log.Println("Reversal Response:")
	log.Println("Reversal ID: ", reversalResponse.PaymentId)
	log.Println("Reversal Status: ", reversalResponse.Transaction.Status)
	log.Println("Reversal Amount: ", reversalResponse.Transaction.Amount)
	log.Println("=================================================")

	lastPaymentState, err := paymentsProvider.GetPaymentById(context.Background(), akuaClient, reversalResponse.PaymentId)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("=================================================")
	log.Println("Last Payment State:")
	log.Println("In this case we need to have the reversal transaction in the payment transactions including authorization and reversal")
	log.Println("Payment ID: ", lastPaymentState.ID)
	log.Println("Payment Status: ", lastPaymentState.Status)
	log.Println("Payment Amount: ", lastPaymentState.CurrentAmount)
	log.Println("Capture Mode: ", lastPaymentState.Capture.Mode)
	log.Println("Payment Transactions: ", len(lastPaymentState.Transactions))

	for _, transaction := range lastPaymentState.Transactions {
		log.Println("=================================================")
		log.Println("Transaction ID: ", transaction.ID)
		log.Println("Transaction Type: ", transaction.Type)
		log.Println("Transaction Status: ", transaction.Status)
		log.Println("Transaction Amount: ", transaction.Amount)
		log.Println("=================================================")
	}

	log.Println("=================================================")

	// Validate if the payment was canceled successfully

	assert.Equal(t, "CANCELLED", lastPaymentState.Status, "expected payment status to be CANCELLED")
}
