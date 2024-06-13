package main

import (
	"crypto/ecdsa"
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
)

type Event struct {
	Acknowledged bool      `json:"acknowledged"`
	CreatedAt    int64     `json:"createdAt"`
	ID           string    `json:"id"`
	Subscription string    `json:"subscription"`
	UserID       string    `json:"userId"`
	Data         EventData `json:"data"`
}

type EventData struct {
	Amount         int    `json:"amount,omitempty"`
	Chain          string `json:"chain,omitempty"`
	DiscountedFee  int    `json:"discountedFee,omitempty"`
	ID             string `json:"id,omitempty"`
	Reason         string `json:"reason,omitempty"`
	ReferenceLabel string `json:"referenceLabel,omitempty"`
	Status         string `json:"status,omitempty"`
	TaxID          string `json:"taxId,omitempty"`
	Tx             string `json:"tx,omitempty"`
}

/*
	func handleWebhook(w http.ResponseWriter, r *http.Request) {
		var event Event

		if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		eventJSON, err := json.MarshalIndent(event, "", " ")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		fmt.Printf("Received event: \n%s\n", eventJSON)
	}
*/
func loadECDSAPublicKey(filePath string) (*ecdsa.PublicKey, error) {
	pemData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("[!ERR!] -> failed to read PEM file: %v", err)
	}

	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("[!ERR!] -> failed to read PEM file: %v", err)
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("[!ERR!] -> failed to parse public key, error: %v", err)
	}

	ecdsaPub, ok := pub.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("[!ERR!] -> your file is not ECDSA public key, error: %v", err)
	}

	return ecdsaPub, nil
}

func handleWebHookCrypt(w http.ResponseWriter, r *http.Request) {
	ecdsaPublicKey, err := loadECDSAPublicKey("public_key_brla.pem")
	if err != nil {
		fmt.Print(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	base64Signature := r.Header.Get("Signature")

	// ECDSA Validation
	sign, err := base64.StdEncoding.DecodeString(base64Signature)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	hash := sha256.New()
	hash.Write(bodyBytes)
	hashedBody := hash.Sum(nil)

	if !ecdsa.VerifyASN1(ecdsaPublicKey, hashedBody, sign) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println("Received data:", string(bodyBytes))

	var event Event
	if err = json.Unmarshal(bodyBytes, &event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Exibir os dados do evento
	eventJSON, err := json.MarshalIndent(event, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Printf("Received event: \n%s\n", eventJSON)

	// Responder com os dados decodificados
	w.Write(eventJSON)
}

func main() {
	http.HandleFunc("/", handleWebHookCrypt)
	fmt.Printf("[INFO] - Listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
