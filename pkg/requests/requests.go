package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ApiError struct {
	Message string `json:"error"`
}

func SendRequest(url string, data interface{}, httpMethod string) error {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal account data: %v", err)
	}

	bodyRender := bytes.NewReader([]byte(jsonBody))

	req, err := http.NewRequest(httpMethod, url, bodyRender)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

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
		var apiErr ApiError
		if err := json.Unmarshal(body, &apiErr); err != nil {
			return fmt.Errorf("failed to unmarshal error response: %v", err)
		}

		return fmt.Errorf("API error: %s", apiErr.Message)
	}

	return nil
}
