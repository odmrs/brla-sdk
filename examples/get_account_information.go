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

	// Get account general information
	data, err := client.GetAccountInfo("JWT TOKEN HERE")

	fmt.Println("[SENDED]\tGet all information of account  sended with success")
	if data != "" {
		fmt.Printf("[GET RESPONSE API] \t %v\n", data)
	}
	if err != nil {
		fmt.Printf("[ERROR] \terror of try get information of account: %v\n", err)
	}
}
