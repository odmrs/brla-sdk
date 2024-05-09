package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/odmrs/brla-sdk/pkg/models"
)

type ApiError struct {
	Code    int    `json:"code"`
	Message string `json:"error"`
}

const (
	createEndpoint string = "/v1/business/create"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) CreateAccount(account *models.Account) error {
	url := c.BaseURL + createEndpoint
	jsonBody, err := json.Marshal(account)
	if err != nil {
		return fmt.Errorf("failed to marshal account data: %v", err)
	}

	bodyRender := bytes.NewReader([]byte(jsonBody))

	resp, err := http.Post(url, "application/json", bodyRender)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if len(body) != 0 {
		var apiErr ApiError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("failed to unmarshal error response: %v", err)
		}

		return fmt.Errorf("API error: %s\n", apiErr.Message)
	}

	return nil
}
