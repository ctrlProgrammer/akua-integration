package adapters_akua_commerce

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	internal_adapters_akua "akua-project/internal/adapters/akua"
	internal_organization "akua-project/internal/organization"
)

type OrganizationResponse struct {
	Data []internal_organization.Organization `json:"data"`
}

type OrganizationProvider struct {
}

func NewOrganizationProvider() *OrganizationProvider {
	return &OrganizationProvider{}
}

func (p *OrganizationProvider) GetOrganizations(ctx context.Context, client *internal_adapters_akua.Client) ([]internal_organization.Organization, error) {
	if !client.JwtIsValid() {
		return nil, internal_adapters_akua.ErrJWTTokenNotSet
	}

	request, err := http.NewRequest("GET", client.GetAudience()+"/v1/organizations", nil)

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
		var organizations OrganizationResponse

		err = json.Unmarshal(bodyBytes, &organizations)

		if err != nil {
			return nil, err
		}

		return organizations.Data, nil
	default: // 400, 500, etc.
		return nil, fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}
