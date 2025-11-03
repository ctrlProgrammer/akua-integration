package adapters_akua_payments

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	internal_adapters_akua "akua-project/internal/adapters/akua"
	"akua-project/internal/payment"
)

type PaymentsResponse struct {
	Data []payment.Payment `json:"data"`
}

type PaymentsProvider struct {
}

func NewPaymentsProvider() *PaymentsProvider {
	return &PaymentsProvider{}
}

func (p *PaymentsProvider) GetPayments(ctx context.Context, client *internal_adapters_akua.Client) ([]payment.Payment, error) {
	if !client.JwtIsValid() {
		return nil, internal_adapters_akua.ErrJWTTokenNotSet
	}

	request, err := http.NewRequest("GET", client.GetAudience()+"/v1/payments", nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", "Bearer "+client.GetJwtToken())
	request.Header.Set("Content-Type", "application/json")

	response, err := client.GetHttpClient().Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case http.StatusOK: // 200
		var payments PaymentsResponse

		err = json.Unmarshal(bodyBytes, &payments)

		if err != nil {
			return nil, err
		}

		return payments.Data, nil
	default: // 400, 500, etc.
		return nil, fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}
