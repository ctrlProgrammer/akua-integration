package adapters_akua_payments

import (
	adapters_akua "akua-project/internal/adapters/akua"
	"context"
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

	assert.NotNil(t, payments)
	assert.NotEmpty(t, payments)
	assert.NotNil(t, payments[0].ID)
}
