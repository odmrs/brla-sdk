package sdk

import (
	"fmt"
	"net/http"

	"github.com/odmrs/brla-sdk/models"
	"github.com/odmrs/brla-sdk/pkg/requests"
)

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

	err := requests.SendRequest(url, account, http.MethodPost)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %v", err)
	}

	return nil
}

func (c *Client) ValidateAccount(email, token string) error {
	url := c.BaseURL + concludesCreationEndpoint
	reqBody := map[string]string{
		"email": email,
		"token": token,
	}

	err := requests.SendRequest(url, reqBody, http.MethodPatch)

	if err != nil {
		return fmt.Errorf("HTTP request failed: %v", err)
	}
	return nil
}
