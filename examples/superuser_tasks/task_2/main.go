package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	sdk "github.com/odmrs/brla-sdk"
)

const (
	sandbox string = "https://api.brla.digital:4567"
	cpf     string = "126.966.870-64"
)

var (
	secretkey     string
	walletAddress string
)

func setEnv() {
	envFile, err := os.Open("../.env")
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to read .env file, error:\n\t%v", err)
		return
	}
	defer envFile.Close()

	// Scanner
	scanner := bufio.NewScanner(envFile)
	for scanner.Scan() {
		envVar := strings.Split(scanner.Text(), "=")
		os.Setenv(envVar[0], envVar[1])
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("[ERROR] \tscanner .env file error:\n\t%v", err)
		return
	}

	walletAddress = os.Getenv("walletaddress")
	secretkey = os.Getenv("secretkey")
}

func main() {
	setEnv()
	fmt.Printf("YOUR WALLET ADDRESS -> %s\n", walletAddress)

	client := sdk.NewClient(sandbox)

	// Validate superuser and get tokenJWT
	tokenJWT, err := client.ValidateSuperAccount("castelobranco@brla.digital", "123456789")
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to auth the SuperUser account, error:\n\t%v", err)
		return
	}
	fmt.Println("\n[SENDED]\tSuper User account auth with successful")
	fmt.Printf("Your token: %v\n", tokenJWT)

	// Get id order to sell BRLA
	permit := map[string]interface{}{
		"r":        "0xd6d9a9010a380eb6af3fe5e20eadd7c0e695941f116150ddbe6b164b7372397e",
		"s":        "0x4386c16e5b6985eb2f451031b337b4cdb9434ef395f5e0a7871b5d502ea486c0",
		"v":        27,
		"deadline": 1718206478,
		"nonce":    1,
	}
	id, err := client.CreateBRLASellOrder(
		tokenJWT,
		cpf,
		"71918232988",
		walletAddress,
		"Polygon",
		8000,
		permit,
	)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to create brla sell order, error:\n\t%v", err)
		return
	}
	fmt.Println("\n[SENDED]\tCreate sell order with successful")
	fmt.Printf("Your id of this sell: \n%v\n", id)
}
