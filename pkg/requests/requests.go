package requests

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
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
			if apiKey, ok := responseMap["apiKey"].(string); ok {
				return apiKey, nil
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

func SendRequestApiKey(
	privateKey *rsa.PrivateKey,
	apiKey,
	urlStr string,
	data interface{},
	httpMethod string,
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

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal data: %v", err)
	}

	// Get sign
	timestamp := time.Now().UnixMilli()
	content := fmt.Sprintf(
		"%s%s%s",
		strconv.FormatInt(timestamp, 10),
		httpMethod,
		baseUrl.Path, // Corrected to use baseUrl.Path instead of urlStr
	)

	hasher := sha256.New()
	hasher.Write([]byte(content))
	hash := hasher.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
	if err != nil {
		log.Fatal(err)
	}

	base64Signature := base64.StdEncoding.EncodeToString(signature)
	fmt.Printf("[INFO] -> Your signature: \n%v\n", base64Signature)
	fmt.Printf("[INFO] -> Your timestamp: \n%v\n", timestamp)

	bodyRender := bytes.NewReader(jsonData) // Use the marshaled JSON data here

	req, err := http.NewRequest(httpMethod, baseUrl.String(), bodyRender)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %v", err)
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("X-API-KEY", apiKey)
	req.Header.Set("X-API-Timestamp", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-API-Signature", base64Signature)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed %v", err)
	}

	defer resp.Body.Close()

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
			if apiKey, ok := responseMap["apiKey"].(string); ok {
				return apiKey, nil
			}
		}
		return "", fmt.Errorf("\n\t\t\tAPI error: %s", string(body))
	}

	return "", nil
}
