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

	// Get account limit information
	data, err := client.GetAccountLimit("JWT TOKEN HERE")

	fmt.Println("[SENDED]\tGet the information of limit account sended with success")
	if data != "" {
		fmt.Printf("[GET RESPONSE API] \t %v\n", data)
	}
	if err != nil {
		fmt.Printf("[ERROR] \terror to get the limit of your account: %v\n", err)
	}
}
