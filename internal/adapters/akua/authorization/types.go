package adapters_akua_authorization

import (
	"akua-project/internal/instruments"
)

const (
	INTENT_AUTHORIZE    = "authorization"
	INTENT_PREAUTHORIZE = "pre-authorization"
)

// Authorization

type AuthorizeRequest struct {
	Amount     instruments.AmountObject     `json:"amount"`
	Intent     string                       `json:"intent"`
	MerchantId string                       `json:"merchant_id"`
	Instrument instruments.InstrumentObject `json:"instrument"`
	Capture    instruments.CaptureObject    `json:"capture"`
}

// Capture

type CaptureRequest struct {
	ID string `json:"id"`
}

type CapturePaymentTransaction struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	Amount string `json:"amount"`
}

type CaptureResponse struct {
	PaymentId   string                    `json:"payment_id"`
	Transaction CapturePaymentTransaction `json:"transaction"`
}

// Reversal

type ReversalPaymentTransaction struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	Amount string `json:"amount"`
}

type ReversalResponse struct {
	PaymentId   string                     `json:"payment_id"`
	Transaction ReversalPaymentTransaction `json:"transaction"`
}

// Refund

type RefundRequest struct {
	Amount instruments.AmountObject `json:"amount"`
}

type RefundPaymentTransaction struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
	Amount string `json:"amount"`
}

type RefundResponse struct {
	PaymentId   string                   `json:"payment_id"`
	Transaction RefundPaymentTransaction `json:"transaction"`
}
