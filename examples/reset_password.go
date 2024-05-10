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

	err := client.AuthLoginPassword("email@example.com", "password")
	if err != nil {
		fmt.Printf("[ERROR] \terror validate authenticates account: %v\n", err)
	}
	fmt.Println("[SENDED]\tAuthentication sended with success")
}
