package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func signApiKey() (string, *rsa.PrivateKey) {
	// read privatekey file
	privateKeyFile, err := os.ReadFile("keypair.pem")
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to read privatekey.pem file, error:\n\t%v", err)
		return "", nil
	}

	// Decode this privatekey
	privateKeyBlock, _ := pem.Decode(privateKeyFile)
	if privateKeyBlock == nil {
		fmt.Printf("[ERROR] \tfailed to decode privatekey.pem file, error: \n\t%v\n", err)
		return "", nil
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to parse private key, error:\n\t%v\n", err)
		return "", nil
	}

	rsaPrivateKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		fmt.Printf("[ERROR] \tfailed to assert type to *rsa.PrivateKey\n")
		return "", nil
	}

	const keyName = "brlaapikey"

	hasher := sha256.New()
	hasher.Write([]byte(keyName))
	hash := hasher.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaPrivateKey, crypto.SHA256, hash)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to sign your API key, error:\n\t%v", err)
		return "", nil
	}

	base64Signature := base64.StdEncoding.EncodeToString(signature)
	return base64Signature, rsaPrivateKey
}
