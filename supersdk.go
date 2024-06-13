package sdk

import (
	"crypto/rsa"
	"net/http"

	"github.com/odmrs/brla-sdk/pkg/requests"
)

const (
	// BRLA SUPERUSER ENDPOINTS
	// validate
	validateSuperAccount string = "/v1/superuser/login"
	validateKYCVerify    string = "/v1/superuser/kyc/pf-free/pass-kyc-level1"
	// Buy
	createPixSuperUserTicket string = "/v1/superuser/buy/static-pix"

	// Sell
	createBRLASellOrder string = "/v1/superuser/sell"

	// Register api key
	registerApiKey string = "/v1/superuser/api-keys"
)

func (c *Client) ValidateSuperAccount(email, password string) (string, error) {
	url := c.BaseURL + validateSuperAccount

	reqBody := map[string]string{
		"email":    email,
		"password": password,
	}

	accessToken, err := requests.SendRequest(url, reqBody, http.MethodPost, nil, nil)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (c *Client) SimulateKYCVerify(tokenJWT, cpf, birthDate, fullName string) (string, error) {
	url := c.BaseURL + validateKYCVerify
	reqBody := map[string]string{
		"cpf":       cpf,
		"birthDate": birthDate,
		"fullName":  fullName,
	}

	id, err := requests.SendRequest(url, reqBody, http.MethodPost, tokenJWT, nil)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *Client) CreatePixSuperUserTicket(
	tokenJWT, taxId, walletAddress, chain string,
	amount int,
) (string, error) {
	url := c.BaseURL + createPixSuperUserTicket
	query := map[string]string{
		"taxId": taxId,
	}
	reqBody := map[string]interface{}{
		"amount":        amount,
		"walletAddress": walletAddress,
		"chain":         chain,
	}

	pixCode, err := requests.SendRequest(url, reqBody, http.MethodPost, tokenJWT, query)
	if err != nil {
		return "", err
	}

	return pixCode, nil
}

func (c *Client) CreateBRLASellOrder(
	tokenJWT, taxId, pixKey, walletAddress, chain string,
	amount int,
	permit map[string]interface{},
) (string, error) {
	url := c.BaseURL + createBRLASellOrder
	query := map[string]string{
		"taxId": taxId,
	}
	reqBody := map[string]interface{}{
		"pixKey":        pixKey,
		"walletAddress": walletAddress,
		"chain":         chain,
		"amount":        amount,
		"permit":        permit,
	}

	pixCode, err := requests.SendRequest(url, reqBody, http.MethodPost, tokenJWT, query)
	if err != nil {
		return "", err
	}

	return pixCode, nil
}

func (c *Client) RegisterApiKey(tokenJWT, name, signature, publicKey string) (string, error) {
	url := c.BaseURL + registerApiKey

	reqBody := map[string]interface{}{
		"signature": signature,
		"name":      name,
		"publicKey": publicKey,
	}

	apiKey, err := requests.SendRequest(url, reqBody, http.MethodPost, tokenJWT, nil)
	if err != nil {
		return "", err
	}

	return apiKey, nil
}

func (c *Client) CreateBRLABuyOrderApiKey(
	apiKey, taxId, walletAddress, chain string,
	amount int,
	privateKey *rsa.PrivateKey,
) (string, error) {
	url := c.BaseURL + createPixSuperUserTicket

	query := map[string]string{
		"taxId": taxId,
	}

	reqBody := map[string]interface{}{
		"walletAddress": walletAddress,
		"chain":         chain,
		"amount":        amount,
	}

	pixCode, err := requests.SendRequestApiKey(
		privateKey,
		apiKey,
		url,
		reqBody,
		http.MethodPost,
		query,
	)
	if err != nil {
		return "", err
	}

	return pixCode, nil
}
