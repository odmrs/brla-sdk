package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/odmrs/brla-sdk/models"
	"github.com/odmrs/brla-sdk/pkg/requests"
)

const (
	// BRLA ACCOUNT ENDPOINTS
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

	// Pay out endpoints
	createPayoutOrder string = "/v1/business/pay-out"
	payoutHistory     string = "/v1/business/pay-out/history"

	// Webhooks
	registerWebhook string = "/v1/business/webhooks"

	// Quote
	quoteTokenConversion string = "/v1/business/fast-quote"

	// Convert operation
	convertCurrencies          string = "/v1/business/swap"
	historyOfConvertCurrencies string = "/v1/business/swap/history"
)

type Client struct {
	BaseURL string
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) CreateAccount(account *models.Account) error {
	url := c.BaseURL + createEndpoint

	_, err := requests.SendRequest(url, account, http.MethodPost, nil, nil)
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

	_, err := requests.SendRequest(url, reqBody, http.MethodPatch, nil, nil)
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

	token, err := requests.SendRequest(url, reqBody, http.MethodPost, nil, nil)
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

	_, err := requests.SendRequest(url, reqBody, http.MethodPost, nil, nil)
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

	_, err := requests.SendRequest(url, reqBody, http.MethodPatch, nil, nil)
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

	_, err := requests.SendRequest(url, reqBody, http.MethodPatch, token, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) LoggoutAccount(token string) error {
	url := c.BaseURL + loggoutAccount
	reqBody := "empty"

	_, err := requests.SendRequest(url, reqBody, http.MethodPost, nil, nil)
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

	_, err := requests.SendRequest(url, "", http.MethodPost, token, nil)
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

	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, responseBody, "", " "); err != nil {
		return "", fmt.Errorf("failed to format JSON: %v", err)
	}

	return prettyJson.String(), nil
}

func (c *Client) CreatePayoutOrder(
	token, pixKey, taxId, referenceLabel, name, ispb, branchCode, accountNumber, accountType string,
	amount int,
) (string, error) {
	url := c.BaseURL + createPayoutOrder

	reqBody := map[string]interface{}{
		"pixKey":         pixKey,
		"taxId":          taxId,
		"amount":         amount,
		"referenceLabel": referenceLabel,
		"name":           name,
		"ispb":           ispb,
		"branchCode":     branchCode,
		"accountNumber":  accountNumber,
		"accountType":    accountType,
	}

	id, err := requests.SendRequest(url, reqBody, http.MethodPost, token, nil)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Client) ShowPayoutHistory(token string) (string, error) {
	url := c.BaseURL + payoutHistory
	responseBody, err := requests.SendRequestGet(url, nil, token)
	if err != nil {
		return "", err
	}

	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, responseBody, "", " "); err != nil {
		return "", fmt.Errorf("failed to format JSON: %v", err)
	}

	return prettyJson.String(), nil
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

func (c *Client) RegisterWebhook(token string) (string, error) {
	var webhookUrl string
	url := c.BaseURL + registerWebhook

	fmt.Print("Enter link of your webhook: ")
	fmt.Scan(&webhookUrl)
	fmt.Printf("your link: %v\n", webhookUrl)
	reqBody := map[string]string{
		"url": webhookUrl,
	}

	webhookId, err := requests.SendRequest(url, reqBody, http.MethodPost, token, nil)
	if err != nil {
		return "", err
	}

	return webhookId, err
}

func (c *Client) QuoteToken(
	tokenJWT, operation, inputCoin, outputCoin, chain,
	fixOutPut,
	amount string,
) (string, error) {
	url := c.BaseURL + quoteTokenConversion
	query := map[string]string{
		"operation":  operation,
		"inputCoin":  inputCoin,
		"outputCoin": outputCoin,
		"chain":      chain,
		"fixOutPut":  fixOutPut,
		"amount":     amount,
	}

	responseBody, err := requests.SendRequestGet(url, query, tokenJWT)
	if err != nil {
		return "", err
	}
	var responseMap map[string]interface{}
	if err := json.Unmarshal(responseBody, &responseMap); err == nil {
		if quoteToken, ok := responseMap["token"].(string); ok {
			return quoteToken, nil
		}
	}
	return "", fmt.Errorf("error to get \"TOKEN\" from the request: \n%v\n", string(responseBody))
}

func (c *Client) ConvertBetweenCurrencies(
	tokenJWT, quoteToken, receiverAddress, markupAddress string,
) (string, error) {
	url := c.BaseURL + convertCurrencies

	reqBody := map[string]interface{}{
		"token":           quoteToken,
		"receiverAddress": receiverAddress,
		"makupAddress":    markupAddress,
	}

	id, err := requests.SendRequest(url, reqBody, http.MethodPost, tokenJWT, nil)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Client) HistoryConversionOperations(tokenJWT string) (string, error) {
	url := c.BaseURL + historyOfConvertCurrencies
	responseBody, err := requests.SendRequestGet(url, nil, tokenJWT)
	if err != nil {
		return "", err
	}

	var prettyJson bytes.Buffer
	if err := json.Indent(&prettyJson, responseBody, "", " "); err != nil {
		return "", fmt.Errorf("failed to format JSON: %v", err)
	}

	return prettyJson.String(), nil
}
