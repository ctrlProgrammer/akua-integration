package adapters_akua_authorization

import (
	"akua-project/internal/instruments"
)

const (
	INTENT_AUTHORIZE    = "authorization"
	INTENT_PREAUTHORIZE = "preauthorization"
)

type AuthorizeRequest struct {
	Amount     instruments.AmountObject     `json:"amount"`
	Intent     string                       `json:"intent"`
	MerchantId string                       `json:"merchant_id"`
	Instrument instruments.InstrumentObject `json:"instrument"`
	Capture    instruments.CaptureObject    `json:"capture"`
}
