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
	authLoginPasswordEndpoint string = "/v1/business/login"
	resetPassword             string = "/v1/business/forgot-password"
	concludesResetPassword    string = "/v1/business/reset-password/"
	changePassword            string = "/v1/business/change-password"
	loggoutAccount            string = "/v1/business/logout"

	// Get endpoints
	getAccount     string = "/v1/business/info"
	accountFees    string = "/v1/business/fees"
	accountBalance string = "/v1/business/balance"
	accountLimit   string = "/v1/business/used-limit"

	// Pay in endpoints
	generatePix      string = "/v1/business/pay-in/br-code"
	payInSandbox     string = "/v1/business/mock-pix-pay-in"
	showHistoryPayin string = "/v1/business/pay-in/pix/history"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) CreateAccount(account *models.Account) error {
	url := c.BaseURL + createEndpoint

	_, err := requests.SendRequest(url, account, http.MethodPost, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ValidateAccount(email string) error {
	var token string

	fmt.Print("[INPUT] Enter the token sended by email: ")
	fmt.Scan(&token)
	url := c.BaseURL + concludesCreationEndpoint
	reqBody := map[string]string{
		"email": email,
		"token": token,
	}

	_, err := requests.SendRequest(url, reqBody, http.MethodPatch, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AuthLoginPassword(email, password string) (string, error) {
	url := c.BaseURL + authLoginPasswordEndpoint
	reqBody := map[string]string{
		"email":    email,
		"password": password,
	}

	token, err := requests.SendRequest(url, reqBody, http.MethodPost, nil)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (c *Client) ResetPassword(email string) error {
	url := c.BaseURL + resetPassword
	reqBody := map[string]string{
		"email": email,
	}

	_, err := requests.SendRequest(url, reqBody, http.MethodPost, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ConcludesResetPassword(token, email string) error {
	url := c.BaseURL + concludesResetPassword + token
	reqBody := map[string]string{
		"email": email,
	}

	_, err := requests.SendRequest(url, reqBody, http.MethodPatch, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ChangePassword(
	currentPassword, newPassword, newPasswordConfirm, token string,
) error {
	url := c.BaseURL + changePassword
	reqBody := map[string]string{
		"currentPassword":    currentPassword,
		"newPassword":        newPassword,
		"newPasswordConfirm": newPasswordConfirm,
	}

	_, err := requests.SendRequest(url, reqBody, http.MethodPatch, token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) LoggoutAccount(token string) error {
	url := c.BaseURL + loggoutAccount
	reqBody := "empty"

	_, err := requests.SendRequest(url, reqBody, http.MethodPost, nil)
	if err != nil {
		return err
	}

	return nil
}

// Get functions

func (c *Client) GetAccountInfo(token string) (string, error) {
	url := c.BaseURL + getAccount
	responseBody, err := requests.SendRequestGet(url, nil, token)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

func (c *Client) GetAccountLimit(token string) (string, error) {
	url := c.BaseURL + getAccount
	responseBody, err := requests.SendRequestGet(url, nil, token)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

func (c *Client) GetAccountBalance(token string) (string, error) {
	url := c.BaseURL + getAccount
	responseBody, err := requests.SendRequestGet(url, nil, token)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

func (c *Client) GetAccountFees(token string) (string, error) {
	url := c.BaseURL + getAccount
	responseBody, err := requests.SendRequestGet(url, nil, token)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

func (c *Client) GeneratePaymantSandbox(token string) error {
	url := c.BaseURL + payInSandbox

	_, err := requests.SendRequest(url, "", http.MethodPost, token)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ShowHistoryPayIn(token string) (string, error) {
	url := c.BaseURL + showHistoryPayin
	responseBody, err := requests.SendRequestGet(url, nil, token)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}

func (c *Client) GeneratesPayinCode(token, amount, referenceLabel, id string) (string, error) {
	url := c.BaseURL + generatePix
	query := map[string]string{
		"amount":         amount,
		"referenceLabel": referenceLabel,
		"subaccountId":   id,
	}

	responseBody, err := requests.SendRequestGet(url, query, token)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}
