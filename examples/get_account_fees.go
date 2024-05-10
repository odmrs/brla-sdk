package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
)

const (
	// sandbox environment
	sandbox string = "https://api.brla.digital:4567"

	// production environment
	// production string = "https://api.brla.digital:5567"
)

func main() {
	// Create the sdk client
	client := sdk.NewClient(sandbox)

	// Get account fees
	data, err := client.GetAccountFees("JWT TOKEN HERE")

	fmt.Println("[SENDED]\tGet account fees with success")
	if data != "" {
		fmt.Printf("[GET RESPONSE API] \t %v\n", data)
	}
	if err != nil {
		fmt.Printf("[ERROR] \terror get the fees of account: %v\n", err)
	}
}
