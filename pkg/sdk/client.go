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
	Message string `json:"error"`
}

const (
	createEndpoint            string = "/v1/business/create"
	concludesCreationEndpoint string = "/v1/business/validate"
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

		return fmt.Errorf("API error: %s", string(body))
	}

	return nil
}

func (c *Client) ValidateAccount(email, token string) error {
	url := c.BaseURL + concludesCreationEndpoint
	reqBody := map[string]string{
		"email": email,
		"token": token,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("HTTP request creation failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json") // Definindo cabeçalho de conteúdo

	resp, err := client.Do(req)
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

		return fmt.Errorf("API error: %s", apiErr.Message)
	}
	return nil
}
