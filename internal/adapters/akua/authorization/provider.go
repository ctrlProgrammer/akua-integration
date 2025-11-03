package adapters_akua_authorization

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	adaters_akua "akua-project/internal/adapters/akua"
	"akua-project/internal/authorization"
)

type AuthorizationProvider struct {
}

func NewAuthorizationProvider() *AuthorizationProvider {
	return &AuthorizationProvider{}
}

func (p *AuthorizationProvider) Authorize(ctx context.Context, client *adaters_akua.Client, requestData AuthorizeRequest) (*authorization.Authorization, error) {
	if !client.JwtIsValid() {
		return nil, adaters_akua.ErrJWTTokenNotSet
	}

	requestBody, err := json.Marshal(requestData)

	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", client.GetAudience()+"/v1/authorizations", bytes.NewBuffer(requestBody))

	if err != nil {
		return nil, err
	}

	hour := time.Now()
	idempotencyKey := fmt.Sprintf("%v%v", requestData.Amount.Value, hour.UnixNano())

	// TODO Change the idempotency key to a more secure way
	// Do it based on tx details to create a more realistic scenario
	request.Header.Set("Idempotency-Key", idempotencyKey)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+client.GetJwtToken())

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
	case http.StatusCreated: // 201
		var authorization authorization.Authorization

		err = json.Unmarshal(bodyBytes, &authorization)

		if err != nil {
			return nil, err
		}

		return &authorization, nil
	default: // 400, 500, etc.
		return nil, fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}

func (p *AuthorizationProvider) Capture(ctx context.Context, client *adaters_akua.Client, requestData CaptureRequest) (*CaptureResponse, error) {
	if !client.JwtIsValid() {
		return nil, adaters_akua.ErrJWTTokenNotSet
	}

	request, err := http.NewRequest("POST", client.GetAudience()+"/v1/payments/"+requestData.ID+"/captures", nil)

	if err != nil {
		return nil, err
	}

	hour := time.Now()
	idempotencyKey := fmt.Sprintf("%v%v", requestData.ID, hour.UnixNano())
	request.Header.Set("Idempotency-Key", idempotencyKey)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+client.GetJwtToken())

	response, err := client.GetHttpClient().Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	log.Println("Capture Response: ", string(bodyBytes))
	log.Println("Capture Response Status Code: ", response.StatusCode)

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated: // 200, 201
		var capture CaptureResponse

		err = json.Unmarshal(bodyBytes, &capture)

		if err != nil {
			return nil, err
		}

		return &capture, nil
	default: // 400, 500, etc.
		return nil, fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}

func (p *AuthorizationProvider) Reversal(ctx context.Context, client *adaters_akua.Client, paymentId string) (*ReversalResponse, error) {
	if !client.JwtIsValid() {
		return nil, adaters_akua.ErrJWTTokenNotSet
	}

	request, err := http.NewRequest("POST", client.GetAudience()+"/v1/payments/"+paymentId+"/reversals", nil)

	if err != nil {
		return nil, err
	}

	hour := time.Now()
	idempotencyKey := fmt.Sprintf("%v%v", paymentId, hour.UnixNano())
	request.Header.Set("Idempotency-Key", idempotencyKey)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+client.GetJwtToken())

	response, err := client.GetHttpClient().Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	log.Println("Capture Response: ", string(bodyBytes))
	log.Println("Capture Response Status Code: ", response.StatusCode)

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated: // 200, 201
		var reversal ReversalResponse

		err = json.Unmarshal(bodyBytes, &reversal)

		if err != nil {
			return nil, err
		}

		return &reversal, nil
	default: // 400, 500, etc.
		return nil, fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}

func (p *AuthorizationProvider) Refund(ctx context.Context, client *adaters_akua.Client, paymentId string) (*RefundResponse, error) {
	if !client.JwtIsValid() {
		return nil, adaters_akua.ErrJWTTokenNotSet
	}

	request, err := http.NewRequest("POST", client.GetAudience()+"/v1/payments/"+paymentId+"/refund", nil)

	if err != nil {
		return nil, err
	}

	hour := time.Now()
	idempotencyKey := fmt.Sprintf("%v%v", paymentId, hour.UnixNano())
	request.Header.Set("Idempotency-Key", idempotencyKey)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+client.GetJwtToken())

	response, err := client.GetHttpClient().Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	log.Println("Capture Response: ", string(bodyBytes))
	log.Println("Capture Response Status Code: ", response.StatusCode)

	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated: // 200, 201
		var refund RefundResponse

		err = json.Unmarshal(bodyBytes, &refund)

		if err != nil {
			return nil, err
		}

		return &refund, nil
	default: // 400, 500, etc.
		return nil, fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}
