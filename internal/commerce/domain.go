package commerce

import "time"

type Address struct {
	Street  string `json:"street"`
	Number  string `json:"number"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

type Fee struct {
	FeeType        string  `json:"fee_type,omitempty"`
	TransactionFee float64 `json:"transaction_fee,omitempty"`
	WithdrawalFee  float64 `json:"withdrawal_fee,omitempty"`
	Currency       string  `json:"currency,omitempty"`
}

type Settlement struct {
	Frequency string `json:"frequency,omitempty"`
}

type Product struct {
	Enabled           bool        `json:"enabled,omitempty"`
	MerchantNetworkID string      `json:"merchant_network_id,omitempty"`
	Settlement        *Settlement `json:"settlement,omitempty"`
	Fee               *Fee        `json:"fee,omitempty"`
}

type Products map[string]Product

type AnnualVolume struct {
	Currency string `json:"currency"`
	Value    int64  `json:"value"`
}

type Rail struct {
	MCC                     string         `json:"mcc,omitempty"`
	RailMerchantExternalID  string         `json:"rail_merchant_external_id,omitempty"`
	PayfacID                string         `json:"payfac_id,omitempty"`
	MerchantVerificationVal string         `json:"merchant_verification_value,omitempty"`
	AnnualVolume            []AnnualVolume `json:"annual_volume,omitempty"`
	Products                Products       `json:"products,omitempty"`
}

type Rails map[string]Rail

type FiscalResponsibility struct {
	Code string `json:"code"`
}

type RetentionRule struct {
	TaxType             string `json:"tax_type"`
	RetentionApplicable bool   `json:"retention_applicable"`
}

type LegalRepresentative struct {
	FullName           string `json:"full_name"`
	Email              string `json:"email"`
	Role               string `json:"role"`
	Phone              string `json:"phone"`
	IdentificationType string `json:"identification_type"`
	IdentificationNum  string `json:"identification_number"`
}

type TaxInformation struct {
	TaxID                  string                 `json:"tax_id"`
	TaxIDType              string                 `json:"tax_id_type"`
	LegalName              string                 `json:"legal_name"`
	FiscalResponsibilities []FiscalResponsibility `json:"fiscal_responsibilities"`
	RetentionRules         []RetentionRule        `json:"retention_rules"`
	ApplicableLaw          string                 `json:"applicable_law"`
	LegalRepresentative    LegalRepresentative    `json:"legal_representative"`
}

type PayoutInformation struct {
	BankName              string `json:"bank_name"`
	BankAccountNumber     string `json:"bank_account_number"`
	BankAccountHolderName string `json:"bank_account_holder_name"`
	BankAccountType       string `json:"bank_account_type"`
	BankSwiftCode         string `json:"bank_swift_code"`
	BankCountry           string `json:"bank_country"`
	Currency              string `json:"currency"`
}

type FeeConfiguration struct {
	FeeType        string  `json:"fee_type"`
	TransactionFee float64 `json:"transaction_fee"`
	WithdrawalFee  float64 `json:"withdrawal_fee"`
	Currency       string  `json:"currency"`
}

type ContractInformation struct {
	Type      string    `json:"type"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type Commerce struct {
	ID                  string              `json:"id,omitempty"`
	OrganizationID      string              `json:"organization_id" validate:"required"`
	Name                string              `json:"name" validate:"required"`
	Type                string              `json:"type" validate:"required"`
	Alias               string              `json:"alias,omitempty"`
	BillingAddress      Address             `json:"billing_address" validate:"required"`
	LocationAddress     Address             `json:"location_address" validate:"required"`
	Email               string              `json:"email,omitempty"`
	Phone               string              `json:"phone,omitempty"`
	Status              string              `json:"status,omitempty"`
	Activity            string              `json:"activity,omitempty"`
	DefaultCurrency     string              `json:"default_currency" validate:"required"`
	SupportedCurrencies []string            `json:"supported_currencies" validate:"required"`
	Rails               Rails               `json:"rails,omitempty"`
	Website             string              `json:"website,omitempty"`
	TaxInformation      TaxInformation      `json:"tax_information,omitempty"`
	PayoutInformation   PayoutInformation   `json:"payout_information,omitempty"`
	FeeConfiguration    FeeConfiguration    `json:"fee_configuration,omitempty"`
	ContractInformation ContractInformation `json:"contract_information,omitempty"`
	Notes               string              `json:"notes,omitempty"`
	CreatedAt           time.Time           `json:"created_at,omitempty"`
	UpdatedAt           time.Time           `json:"updated_at,omitempty"`
}
