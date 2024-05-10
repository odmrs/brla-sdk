package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
)

func main() {
	// Create the sdk client
	client := sdk.NewClient(sandbox)

	// Validate account
	err := client.ValidateAccount("email@example.com", "token")
	if err != nil {
		fmt.Printf("[ERROR] \terror validate creating account: %v\n", err)
	}
	fmt.Println("[SENDED]\tValidate Account sended with success")
}
