package internal_adapters_akua

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type JwtTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Client struct {
	httpClient *http.Client
	apiClient  string
	apiSecret  string
	jwtToken   string
	baseUrl    string
	audience   string
}

func NewClient() (*Client, error) {
	vars := map[string]string{
		"AKUA_API_CLIENT": os.Getenv("AKUA_API_CLIENT"),
		"AKUA_API_SECRET": os.Getenv("AKUA_API_SECRET"),
		"AKUA_BASE_URL":   os.Getenv("AKUA_BASE_URL"),
		"AKUA_AUDIENCE":   os.Getenv("AKUA_AUDIENCE"),
	}

	for k, v := range vars {
		if v == "" {
			return nil, fmt.Errorf("%s is not set", k)
		}
	}

	client := &Client{
		httpClient: &http.Client{},
		apiClient:  vars["AKUA_API_CLIENT"],
		apiSecret:  vars["AKUA_API_SECRET"],
		baseUrl:    vars["AKUA_BASE_URL"],
		audience:   vars["AKUA_AUDIENCE"],
	}

	err := client.LoadJwtToken()

	if err != nil {
		return nil, err
	}

	return client, nil
}

// LoadJwtToken retrieves a new JWT access token using the client credentials (client ID, client secret, and audience)
// and updates the Client with the new token. It does so by making a POST request to the /oauth/token endpoint of the
// authentication provider specified by baseUrl. If the request is successful and returns a 201 status, the token is
// parsed from the response and set in the Client. Returns an error if the request fails or the response cannot be parsed.
func (c *Client) LoadJwtToken() error {
	requestBody := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     c.apiClient,
		"client_secret": c.apiSecret,
		"audience":      c.audience,
	}

	bodyBytes, err := json.Marshal(requestBody)

	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", c.baseUrl+"/oauth/token", bytes.NewBuffer(bodyBytes))

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := c.httpClient.Do(request)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	bodyBytes, err = io.ReadAll(response.Body)

	if err != nil {
		return err
	}

	switch response.StatusCode {
	case http.StatusCreated: // 201
		// Success, continue parsing token
		var jwtTokenResponse JwtTokenResponse

		err = json.Unmarshal(bodyBytes, &jwtTokenResponse)

		if err != nil {
			return err
		}

		c.jwtToken = jwtTokenResponse.AccessToken

		return nil
	case http.StatusBadRequest: // 400
		return fmt.Errorf("received 400 Bad Request: %s", string(bodyBytes))
	case http.StatusInternalServerError: // 500
		return fmt.Errorf("received 500 Internal Server Error: %s", string(bodyBytes))
	default:
		return fmt.Errorf("unexpected status code %d: %s", response.StatusCode, string(bodyBytes))
	}
}
