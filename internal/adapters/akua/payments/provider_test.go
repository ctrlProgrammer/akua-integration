package adapters_akua_payments

import (
	adapters_akua "akua-project/internal/adapters/akua"
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Setup() (*adapters_akua.Client, *PaymentsProvider, error) {
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

	return akuaClient, NewPaymentsProvider(), nil
}

// Test to get the organizations from the Akua API
// This test will validate if the organizations are returned correctly
func Test_GetPayments_Success(t *testing.T) {
	akuaClient, provider, err := Setup()

	if err != nil {
		t.Fatal(err)
	}

	payments, err := provider.GetPayments(context.Background(), akuaClient)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("Payments: ", len(payments))

	paymentsJsonPath := filepath.Join("testdata", "payments.json")
	err = nil

	// Ensure testdata directory exists
	testdataDir := filepath.Dir(paymentsJsonPath)
	if err := os.MkdirAll(testdataDir, 0755); err != nil {
		t.Fatalf("failed to create testdata directory: %v", err)
	}

	file, err := os.Create(paymentsJsonPath)
	if err != nil {
		t.Fatalf("failed to create payments.json: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(payments); err != nil {
		t.Fatalf("failed to encode payments to JSON: %v", err)
	}

	assert.NotNil(t, payments)
	assert.NotEmpty(t, payments)
	assert.NotNil(t, payments[0].ID)
}
