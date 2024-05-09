package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendRequest(url string, data interface{}, httpMethod string, bearer interface{}) error {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal account data: %v", err)
	}

	bodyRender := bytes.NewReader([]byte(jsonBody))

	req, err := http.NewRequest(httpMethod, url, bodyRender)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %v", err)
	}

	if bearer != nil {
		var tokenBearer string = "Bearer " + bearer.(string)
		req.Header.Add("authorization", tokenBearer)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed %v", err)
	}

	defer req.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}
	if len(body) != 0 {
		return fmt.Errorf("\n\t\t\tAPI error: %s", string(body))
	}

	return nil
}

func SendRequestGet(url, bearerToken string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
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
