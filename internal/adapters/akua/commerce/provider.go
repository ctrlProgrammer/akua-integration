package adapters_akua_commerce

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	adaters_akua "akua-project/internal/adapters/akua"
	commerce "akua-project/internal/commerce"
)

type CommerceProvider struct {
}

func NewCommerceProvider() *CommerceProvider {
	return &CommerceProvider{}
}

func (p *CommerceProvider) GetOrganizationCommerces(ctx context.Context, client *adaters_akua.Client) ([]commerce.Commerce, error) {
	if !client.JwtIsValid() {
		return nil, adaters_akua.ErrJWTTokenNotSet
	}

	request, err := http.NewRequest("GET", client.GetAudience()+"/v1/merchants", nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+client.GetJwtToken())
	request.Header.Set("Content-Type", "application/json")

	response, err := client.GetHttpClient().Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK: // 200
		var commerces GetOrganizationCommercesResponse

		err = json.Unmarshal(bodyBytes, &commerces)

		if err != nil {
			return nil, err
		}

		return commerces.Data, nil
	default: // 400, 500, etc.
		return nil, fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}

func (p *CommerceProvider) CreateCommerce(ctx context.Context, client *adaters_akua.Client, requestData CreateCommerceRequest) (*commerce.Commerce, error) {
	if !client.JwtIsValid() {
		return nil, adaters_akua.ErrJWTTokenNotSet
	}

	requestBody, err := json.Marshal(requestData)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", client.GetAudience()+"/v1/merchants", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+client.GetJwtToken())
	request.Header.Set("Content-Type", "application/json")

	response, err := client.GetHttpClient().Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusCreated: // 201
		generatedCommerce := &commerce.Commerce{}

		err = json.Unmarshal(bodyBytes, &generatedCommerce)

		if err != nil {
			return nil, err
		}

		return generatedCommerce, nil
	default: // 400, 500, etc.
		return nil, fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}

func (p *CommerceProvider) UpdateCommerceRails(ctx context.Context, client *adaters_akua.Client, requestData UpdateCommerceRailsRequest) (*commerce.Rails, error) {
	if !client.JwtIsValid() {
		return nil, adaters_akua.ErrJWTTokenNotSet
	}

	requestBody, err := json.Marshal(requestData.Rails)

	if err != nil {
		return nil, err
	}

	log.Println("Request Body: ", string(requestBody))
	log.Println("Request ID: ", requestData.ID)

	request, err := http.NewRequest("PUT", client.GetAudience()+"/v1/merchants/"+requestData.ID+"/rails", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+client.GetJwtToken())
	request.Header.Set("Content-Type", "application/json")

	response, err := client.GetHttpClient().Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK: // 200
		generatedRails := &commerce.Rails{}

		err = json.Unmarshal(bodyBytes, &generatedRails)

		if err != nil {
			return nil, err
		}

		return generatedRails, nil
	default: // 400, 500, etc.
		return nil, fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}
