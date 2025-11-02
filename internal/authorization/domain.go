package authorization

type Authorization struct {
	InstrumentID            string                   `json:"instrument_id"`
	PaymentID               string                   `json:"payment_id"`
	ResponseCode            string                   `json:"response_code"`
	ResponseCodeDescription string                   `json:"response_code_description"`
	Transaction             AuthorizationTransaction `json:"transaction"`
}

type AuthorizationTransaction struct {
	Amount       string                              `json:"amount"`
	ID           string                              `json:"id"`
	NetworkData  AuthorizationTransactionNetworkData `json:"network_data"`
	RiskID       string                              `json:"risk_id"`
	Status       string                              `json:"status"`
	StatusDetail string                              `json:"status_detail"`
	Type         string                              `json:"type"`
}

type AuthorizationTransactionNetworkData struct {
	ApprovalCode              string `json:"approval_code"`
	BanknetReferenceNumber    string `json:"banknet_reference_number"`
	FinancialNetworkCode      string `json:"financial_network_code"`
	ResponseCode              string `json:"response_code"`
	ResponseCodeDescription   string `json:"response_code_description"`
	SettlementDate            string `json:"settlement_date"`
	SystemTraceAuditNumber    string `json:"system_trace_audit_number"`
	TransmissionDateTime      string `json:"transmission_date_time"`
	MerchantAdviceCode        string `json:"merchant_advice_code"`
	MerchantAdviceDescription string `json:"merchant_advice_description"`
	MerchantAdviceAction      string `json:"merchant_advice_action"`
}
