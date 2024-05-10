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

	// Logs account out
	err := client.LoggoutAccount("JWT TOKEN HERE")

	fmt.Println("[SENDED]\tLoggout account sended with success")

	if err != nil {
		fmt.Printf("[ERROR] \terror to loggout your account: %v\n", err)
	}
}
