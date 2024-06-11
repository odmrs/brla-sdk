package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func SendRequest(
	urlStr string,
	data interface{},
	httpMethod string,
	bearer interface{},
	queryParams map[string]string,
) (string, error) {
	baseUrl, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("error to parse URL: %v", err)
	}

	if len(queryParams) > 0 {
		params := url.Values{}
		for key, value := range queryParams {
			params.Add(key, value)
		}

		baseUrl.RawQuery = params.Encode()
	}
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal account data: %v", err)
	}

	bodyRender := bytes.NewReader([]byte(jsonBody))

	req, err := http.NewRequest(httpMethod, baseUrl.String(), bodyRender)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %v", err)
	}

	if bearer != nil {
		tokenBearer := "Bearer " + bearer.(string)
		req.Header.Add("authorization", tokenBearer)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed %v", err)
	}

	defer req.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	if len(body) != 0 {
		var responseMap map[string]interface{}
		if err := json.Unmarshal(body, &responseMap); err == nil {
			if accessToken, ok := responseMap["accessToken"].(string); ok {
				return accessToken, nil
			}
			if pix, ok := responseMap["brCode"].(string); ok {
				return pix, nil
			}
			if id, ok := responseMap["id"].(string); ok {
				return id, nil
			}
			if webhookId, ok := responseMap["webhookId"].(string); ok {
				return webhookId, nil
			}
		}
		return "", fmt.Errorf("\n\t\t\tAPI error: %s", string(body))
	}

	return "", nil
}

func SendRequestGet(
	urlStr string,
	queryParams map[string]string,
	bearerToken string,
) ([]byte, error) {
	baseUrl, err := url.Parse(urlStr)
	if err != nil {
		return nil, fmt.Errorf("error to parse URL: %v", err)
	}

	if len(queryParams) > 0 {
		params := url.Values{}
		for key, value := range queryParams {
			params.Add(key, value)
		}

		baseUrl.RawQuery = params.Encode()
	}
	req, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("error to create request: %v", err)
	}

	req.Header.Add("Authorization", "Bearer "+bearerToken)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error to send request: %v", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error to read the body of response: %v", err)
	}

	return body, nil
}
