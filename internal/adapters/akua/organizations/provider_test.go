package adapters_akua_organizations

import (
	adapters_akua "akua-project/internal/adapters/akua"
	"context"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Setup() (*adapters_akua.Client, *OrganizationProvider, error) {
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

	return akuaClient, NewOrganizationProvider(), nil
}

// Test to get the organizations from the Akua API
// This test will validate if the organizations are returned correctly
func Test_GetOrganizations_Success(t *testing.T) {
	akuaClient, provider, err := Setup()

	if err != nil {
		t.Fatal(err)
	}

	organizations, err := provider.GetOrganizations(context.Background(), akuaClient)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, organizations)
	assert.NotEmpty(t, organizations)
	assert.Equal(t, len(organizations), 1)
	assert.NotNil(t, organizations[0].ID)
}
