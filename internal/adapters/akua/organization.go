package internal_adapters_akua

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	internal_organization "akua-project/internal/organization"
)

type OrganizationProvider struct {
}

func NewOrganizationProvider() *OrganizationProvider {
	return &OrganizationProvider{}
}

func (p *OrganizationProvider) GetOrganizations(ctx context.Context, client *Client) ([]internal_organization.Organization, error) {
	if client.jwtToken == "" {
		return nil, ErrJWTTokenNotSet
	}

	request, err := http.NewRequest("GET", client.baseUrl+"/organizations", nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+client.jwtToken)
	request.Header.Set("Content-Type", "application/json")

	response, err := client.httpClient.Do(request)

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
		var organizations []internal_organization.Organization

		err = json.Unmarshal(bodyBytes, &organizations)

		if err != nil {
			return nil, err
		}

		return organizations, nil
	default: // 400, 500, etc.
		return nil, fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}
