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

	// Concludes password reset process
	err := client.ConcludesResetPassword("tokenblablabla", "email@gmail.com")
	if err != nil {
		fmt.Printf("[ERROR] \terror concludes reset password: %v\n", err)
	}
	fmt.Println("[SENDED]\tConcludes password sended with success")
}
