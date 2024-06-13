package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"net/http"
	"os"
)

const sandbox string = "https://api.brla.digital:4567/v1/pubkey"

type PublicKeyResponse struct {
	PublicKey string `json:"publickey"`
}

func main() {
	req, _ := http.NewRequest("GET", sandbox, nil)
	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[ERROR] -> failed to send a request to endpoint, error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	pubKeyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[ERROR] -> failed to read response body, error: %v\n", err)
		return
	}

	fmt.Printf("Raw response body: %s\n", pubKeyBytes)

	var pubKeyData PublicKeyResponse

	err = json.Unmarshal(pubKeyBytes, &pubKeyData)
	if err != nil {
		fmt.Printf("[ERROR] -> failed to Unmarshal body, error: %v\n", err)
		return
	}

	blockPub, _ := pem.Decode([]byte(pubKeyData.PublicKey))
	if blockPub == nil {
		fmt.Printf("[ERROR] -> failed to decode pubkeydata to pem format\n")
		return
	}

	pubKeyUntyped, err := x509.ParsePKIXPublicKey(blockPub.Bytes)
	if err != nil {
		fmt.Printf("[ERROR] -> failed to parse blockPub to x509 function, error: %v\n", err)
		return
	}

	pubKey, ok := pubKeyUntyped.(*ecdsa.PublicKey)
	if !ok {
		fmt.Printf("[ERROR] -> failed to parse pubKey to right format\n")
		return
	}

	// Convert to PEM FORMAT
	pubKeyPem := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: blockPub.Bytes,
	})

	if pubKeyPem == nil {
		fmt.Printf("[ERROR] -> failed to encode pubkeydata to pem format, error: %v\n", err)
		return
	}

	err = os.WriteFile("public_key_brla.pem", pubKeyPem, 0644)
	if err != nil {
		fmt.Printf("[ERROR] -> failed to write your pubkey, error: %v\n", err)
		return
	}
	fmt.Println("Your publickey saved to public_key_brla.pem")
	fmt.Printf("BRLA Digital pubkey -> \n%s\n", pubKey)
}
