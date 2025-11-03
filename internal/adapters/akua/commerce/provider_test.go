package adapters_akua_commerce

import (
	adapters_akua "akua-project/internal/adapters/akua"
	commerce "akua-project/internal/commerce"
	"context"
	"log"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Setup() (*adapters_akua.Client, *CommerceProvider, error) {
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

	return akuaClient, NewCommerceProvider(), nil
}

// Test to get the organizations from the Akua API
// This test will validate if the organizations are returned correctly
func Test_GetOrganizationCommerces_Success(t *testing.T) {
	akuaClient, provider, err := Setup()

	if err != nil {
		t.Fatal(err)
	}

	commerces, err := provider.GetOrganizationCommerces(context.Background(), akuaClient)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, commerces)
	assert.NotEmpty(t, commerces)
	assert.Equal(t, len(commerces), 1)
	assert.NotNil(t, commerces[0].ID)

	log.Println(commerces)
}

func Test_CreateCommerce_Success(t *testing.T) {
	akuaClient, provider, err := Setup()

	if err != nil {
		t.Fatal(err)
	}

	request := CreateCommerceRequest{
		Type:                "ECOMMERCE",
		Name:                "Test Commerce",
		OrganizationID:      akuaClient.GetOrganizationId(),
		SupportedCurrencies: []string{"USD"},
		DefaultCurrency:     "USD",
		Website:             "https://test.com",
		BillingAddress: commerce.Address{
			Street:  "Test Street",
			Number:  "123",
			City:    "Bogotá",
			State:   "Cundinamarca",
			ZipCode: "12345",
			Country: "COL", // Take care with the country code because is a static param based on the Akua country codes
		},
		LocationAddress: commerce.Address{
			Street:  "Test Street",
			Number:  "123",
			City:    "Bogotá",
			State:   "Cundinamarca",
			ZipCode: "12345",
			Country: "COL", // Take care with the country code because is a static param based on the Akua country codes
		},
	}

	generatedCommerce, err := provider.CreateCommerce(context.Background(), akuaClient, request)

	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, generatedCommerce)
	assert.NotNil(t, generatedCommerce.ID)
}

func Test_UpdateCommerceRails_Success(t *testing.T) {
	akuaClient, provider, err := Setup()

	if err != nil {
		t.Fatal(err)
	}

	request := UpdateCommerceRailsRequest{
		ID: akuaClient.GetMerchantId(),
		Rails: commerce.Rails{
			"MASTERCARD": commerce.Rail{
				MCC: "5678",
				AnnualVolume: []commerce.AnnualVolume{
					{
						Currency: "USD",
						Value:    1000000,
					},
				},
				Products: commerce.Products{
					"CREDIT": commerce.Product{
						Enabled: true,
					},
					"DEBIT": commerce.Product{
						Enabled: true,
					},
				},
			},
		},
	}

	generatedRails, err := provider.UpdateCommerceRails(context.Background(), akuaClient, request)

	if err != nil {
		t.Fatal(err)
	}

	log.Println("Generated Rails: ", generatedRails)

	assert.NotNil(t, generatedRails)
}
