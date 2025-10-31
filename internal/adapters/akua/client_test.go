package internal_adapters_akua

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAkuaService struct {
	mock.Mock
}

func Setup() (*MockAkuaService, error) {
	envPath := filepath.Join("..", "..", "..", ".env")

	err := godotenv.Load(envPath)

	if err != nil {
		return nil, err
	}

	return &MockAkuaService{}, nil
}

// This is a real case test, this will call directly the akua services and try to get the JWT token to initialize the Akua adapter
// This is only for testing the connection purpose this will not be included in the isolated tests
// Will validate if with the loaded env we can load the JWT token
func Test_Connection_Real(t *testing.T) {
	_, err := Setup()

	if err != nil {
		t.Fatal(err.Error())
	}

	akuaClient, err := NewClient()

	if err != nil {
		t.Fatal(err.Error())
	}

	defer akuaClient.httpClient.CloseIdleConnections()

	err = akuaClient.LoadJwtToken()

	if err != nil {
		t.Fatal(err.Error())
	}

	assert.NotNil(t, akuaClient.jwtToken)
}

// The next ones are only test cases for the Akua adapter, I will use mocks to do it
// Test env variables loading
// Will validate if all the necessary env variables are set
func Test_EnvVariablesLoading_Real(t *testing.T) {
	_, err := Setup()

	if err != nil {
		t.Fatal(err.Error())
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

	assert.Nil(t, err)
}
