package main

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	privateKeyFile, err := os.ReadFile("keypair.pem")
	if err != nil {
		log.Fatalf("[ERROR] \tfailed to read privatekey.pem file, error:\n\t%v", err)
	}

	// Decode this private key
	privateKeyBlock, _ := pem.Decode(privateKeyFile)
	if privateKeyBlock == nil {
		log.Fatalf("[ERROR] \tfailed to decode privatekey.pem file, error: \n\t%v\n", err)
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		log.Fatalf("[ERROR] \tfailed to parse private key, error:\n\t%v\n", err)
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		log.Fatalf("[ERROR] \tfailed to assert type to *rsa.PrivateKey\n")
	}

	timestamp := time.Now().UnixMilli()
	requestMethod := http.MethodPost
	endpointPath := "/v1/superuser/buy/static-pix?taxId=12345678900"
	apiURL := "https://api.brla.digital:4567"

	bodyData, _ := json.Marshal(struct {
		WalletAddress string `json:"walletAddress"`
		Chain         string `json:"chain"`
		Amount        int64  `json:"amount"`
	}{
		WalletAddress: "0xd22c5da587869232dfdfe53a60873A2166027b51",
		Chain:         "Polygon",
		Amount:        5000,
	})

	content := fmt.Sprintf(
		"%s%s%s%s",
		strconv.FormatInt(timestamp, 10),
		requestMethod,
		endpointPath,
		string(bodyData),
	)

	fmt.Println("Content to sign:", content) // Log the content to sign

	hasher := sha256.New()
	hasher.Write([]byte(content))
	hash := hasher.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hash)
	if err != nil {
		log.Fatalf("[ERROR] \tfailed to sign content, error:\n\t%v", err)
	}

	base64Signature := base64.StdEncoding.EncodeToString(signature)
	fmt.Println("Base64 Signature:", base64Signature) // Log the base64 signature

	req, _ := http.NewRequest(requestMethod, apiURL+endpointPath, bytes.NewReader(bodyData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", "94ad0ddd-0c63-4198-9db5-4cdb1960770a")
	req.Header.Set("X-API-Timestamp", strconv.FormatInt(timestamp, 10))
	req.Header.Set("X-API-Signature", base64Signature)

	// Log the request headers
	fmt.Println("Headers:")
	fmt.Println("Content-Type:", req.Header.Get("Content-Type"))
	fmt.Println("X-API-Key:", req.Header.Get("X-API-Key"))
	fmt.Println("X-API-Timestamp:", req.Header.Get("X-API-Timestamp"))
	fmt.Println("X-API-Signature:", req.Header.Get("X-API-Signature"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("[ERROR] \tfailed to perform HTTP request, error:\n\t%v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(res.Body)
		log.Fatalf(
			"[ERROR] \trequest failed, status code: %d, body: %s",
			res.StatusCode,
			string(body),
		)
	}

	fmt.Println("Request successful!")
}
