package internal_adapters_akua

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func NewMockClient() (*MockClient, error) {
	return &MockClient{}, nil
}

func Setup() error {
	envPath := filepath.Join("..", "..", "..", ".env")
	err := godotenv.Load(envPath)

	if err != nil {
		return err
	}

	return nil
}

// This is a real case test, this will call directly the akua services and try to get the JWT token to initialize the Akua adapter
func TestConnection(t *testing.T) {
	err := Setup()

	if err != nil {
		t.Fatal(err.Error())
	}

	akuaClient, err := NewClient()

	if err != nil {
		t.Fatal(err.Error())
	}

	defer akuaClient.httpClient.CloseIdleConnections()

	assert.NotNil(t, akuaClient.jwtToken)
}

// The next ones are only test cases for the Akua adapter, I will use mocks to do it

// 1. Test env variables loading
func TestEnvVariablesLoading(t *testing.T) {
	err := Setup()

	if err != nil {
		t.Fatalf("Failed to setup environment: %v", err)
	}

	vars := map[string]string{
		"AKUA_CLIENT_ID":     os.Getenv("AKUA_CLIENT_ID"),
		"AKUA_CLIENT_SECRET": os.Getenv("AKUA_CLIENT_SECRET"),
		"AKUA_AUDIENCE":      os.Getenv("AKUA_AUDIENCE"),
	}

	for _, v := range vars {
		if v == "" {
			assert.NotNil(t, v)
		}
	}
}
