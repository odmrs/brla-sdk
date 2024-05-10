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

	// Change account password
	err := client.ChangePassword("currentPassword", "newpassword", "newpassword", "blablabla")

	fmt.Println("[SENDED]\tChange password sended with success")

	if err != nil {
		fmt.Printf("[ERROR] \terror to change your account password: %v\n", err)
	}
}
