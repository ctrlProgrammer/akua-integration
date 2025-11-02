package instruments

type CaptureMode string

const (
	CAPTURE_MODE_AUTOMATIC CaptureMode = "AUTOMATIC"
	CAPTURE_MODE_MANUAL    CaptureMode = "MANUAL"
)

type Instrument struct {
	Number          string `json:"number"`
	CVV             string `json:"cvv"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
	HolderName      string `json:"holder_name"`
}

type AmountObject struct {
	Currency string  `json:"currency"`
	Value    float64 `json:"value"`
}

type InstrumentObject struct {
	Type string     `json:"type"`
	Card Instrument `json:"card"`
}

type CaptureObject struct {
	Mode CaptureMode `json:"mode"`
}
