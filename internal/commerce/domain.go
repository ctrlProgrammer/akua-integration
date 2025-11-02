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
	FeeType        string  `json:"fee_type"`
	TransactionFee float64 `json:"transaction_fee"`
	WithdrawalFee  float64 `json:"withdrawal_fee"`
	Currency       string  `json:"currency"`
}

type Settlement struct {
	Frequency string `json:"frequency"`
}

type Product struct {
	Enabled           bool       `json:"enabled"`
	MerchantNetworkID string     `json:"merchant_network_id"`
	Settlement        Settlement `json:"settlement"`
	Fee               Fee        `json:"fee"`
}

type Products map[string]Product

type AnnualVolume struct {
	Currency string `json:"currency"`
	Value    int64  `json:"value"`
}

type Rail struct {
	MCC                     string         `json:"mcc"`
	RailMerchantExternalID  string         `json:"rail_merchant_external_id"`
	PayfacID                string         `json:"payfac_id"`
	MerchantVerificationVal string         `json:"merchant_verification_value"`
	AnnualVolume            []AnnualVolume `json:"annual_volume"`
	Products                Products       `json:"products"`
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
	ID                  string              `json:"id"`
	OrganizationID      string              `json:"organization_id"`
	Name                string              `json:"name"`
	Type                string              `json:"type"`
	Alias               string              `json:"alias"`
	BillingAddress      Address             `json:"billing_address"`
	LocationAddress     Address             `json:"location_address"`
	Email               string              `json:"email"`
	Phone               string              `json:"phone"`
	Status              string              `json:"status"`
	Activity            string              `json:"activity"`
	DefaultCurrency     string              `json:"default_currency"`
	SupportedCurrencies []string            `json:"supported_currencies"`
	Rails               Rails               `json:"rails"`
	Website             string              `json:"website"`
	TaxInformation      TaxInformation      `json:"tax_information"`
	PayoutInformation   PayoutInformation   `json:"payout_information"`
	FeeConfiguration    FeeConfiguration    `json:"fee_configuration"`
	ContractInformation ContractInformation `json:"contract_information"`
	Notes               string              `json:"notes"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
}
