package payment

type Payment struct {
	ID                    string               `json:"id"`
	ExternalClientTraceID string               `json:"external_client_trace_id"`
	ClientID              string               `json:"client_id"`
	Capture               PaymentCapture       `json:"capture"`
	Status                string               `json:"status"`
	StatusDetail          string               `json:"status_detail"`
	InitialAmount         float64              `json:"initial_amount"`
	CurrentAmount         float64              `json:"current_amount"`
	Currency              string               `json:"currency"`
	Country               string               `json:"country"`
	OrganizationID        string               `json:"organization_id"`
	PaymentInstrument     PaymentInstrument    `json:"payment_instrument"`
	MerchantID            string               `json:"merchant_id"`
	CreatedAt             interface{}          `json:"created_at"`
	UpdatedAt             interface{}          `json:"updated_at"`
	Transactions          []PaymentTransaction `json:"transactions"`
}

type PaymentCapture struct {
	Mode                string      `json:"mode"`
	EstimatedDate       interface{} `json:"estimated_date"`
	EstimatedDateReason string      `json:"estimated_date_reason"`
}

type PaymentInstrument struct {
	Type string      `json:"type"`
	ID   string      `json:"id"`
	Rail PaymentRail `json:"rail"`
}

type PaymentRail struct {
	ID      string `json:"id"`
	Product string `json:"product"`
}

type PaymentTransaction struct {
	ID                string               `json:"id"`
	Type              string               `json:"type"`
	Status            string               `json:"status"`
	StatusDetail      string               `json:"status_detail,omitempty"`
	Amount            float64              `json:"amount"`
	AuthorizationCode string               `json:"authorization_code,omitempty"`
	Timestamp         string               `json:"timestamp"`
	CreatedAt         interface{}          `json:"created_at"`
	UpdatedAt         interface{}          `json:"updated_at"`
	NetworkAudit      *PaymentNetworkAudit `json:"network_audit,omitempty"`
}

type PaymentNetworkAudit struct {
	ApprovalCode           string `json:"approval_code"`
	SystemTraceAuditNumber string `json:"system_trace_audit_number"`
	TransmissionDateTime   string `json:"transmission_date_time"`
}
