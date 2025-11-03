package adapters_akua_authorization

import (
	"akua-project/internal/instruments"
)

const (
	INTENT_AUTHORIZE    = "authorization"
	INTENT_PREAUTHORIZE = "pre-authorization"
)

type AuthorizeRequest struct {
	Amount     instruments.AmountObject     `json:"amount"`
	Intent     string                       `json:"intent"`
	MerchantId string                       `json:"merchant_id"`
	Instrument instruments.InstrumentObject `json:"instrument"`
	Capture    instruments.CaptureObject    `json:"capture"`
}

type CaptureRequest struct {
	ID string `json:"id"`
}

type CapturePaymentTransaction struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

type CapturePayment struct {
	ID          string                    `json:"id"`
	Transaction CapturePaymentTransaction `json:"transaction"`
}

type CaptureResponse struct {
	Payment CapturePayment `json:"payment"`
}
