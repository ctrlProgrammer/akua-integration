package instruments

// Test cards, this will be used to test the payment flows
// The cards are not real, they are just for testing purposes
// https://docs.akua.la/reference/tarjetas-de-prueba
const (
	CARD_APPROVED_CREDIT           = "5191872272166422"
	CARD_APPROVED_DEBIT            = "5200000000000007"
	CARD_DECLINED_NO_FUNDS         = "5555444433331111"
	CARD_DECLINED_FRAUDULENT       = "5404000000000001"
	CARD_DECLINED_EXIRED           = "5100000000000008"
	CARD_DECLINED_PROCESSING_ERROR = "5555555555554444"
	CARD_DECLINED_REJECTED         = "5200828282828210"
	CARD_REVERSE_CANCEL_SUCCESS    = "5100000000000018"
	CARD_REVERSE_CANCEL_ERROR      = "5100000000000019"
)
