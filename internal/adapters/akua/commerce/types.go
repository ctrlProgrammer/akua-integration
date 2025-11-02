package adapters_akua_commerce

import commerce "akua-project/internal/commerce"

type GetOrganizationCommercesResponse struct {
	Data []commerce.Commerce `json:"data"`
}

type CreateCommerceRequest struct {
	Type                string           `json:"type" validate:"required"`
	Name                string           `json:"name" validate:"required"`
	OrganizationID      string           `json:"organization_id" validate:"required"`
	SupportedCurrencies []string         `json:"supported_currencies" validate:"required"`
	DefaultCurrency     string           `json:"default_currency" validate:"required"`
	BillingAddress      commerce.Address `json:"billing_address" validate:"required"`
	LocationAddress     commerce.Address `json:"location_address" validate:"required"`
	Website             string           `json:"website" validate:"required"`
}
