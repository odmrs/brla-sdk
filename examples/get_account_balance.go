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

	// Get account balance information
	data, err := client.GetAccountBalance("JWT TOKEN HERE")

	fmt.Println("[SENDED]\tGet the balance of account sended with success")
	if data != "" {
		fmt.Printf("[GET RESPONSE API] \t %v\n", data)
	}
	if err != nil {
		fmt.Printf("[ERROR] \terror to get balance of your account: %v\n", err)
	}
}
