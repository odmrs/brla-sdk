package sdk

import (
	"net/http"

	"github.com/odmrs/brla-sdk/pkg/requests"
)

const (
	// BRLA SUPERUSER ENDPOINTS
	validateSuperAccount     string = "/v1/superuser/login"
	createPixSuperUserTicket string = "/v1/superuser/buy/static-pix"
	validateKYCVerify        string = "/v1/superuser/kyc/pf-free/pass-kyc-level1"
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
