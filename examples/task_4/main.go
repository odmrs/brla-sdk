package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
)

const sandbox string = "https://api.brla.digital:4567"

func main() {
	client := sdk.NewClient(sandbox)
	// Auth user with login and password -> validate endpoint
	token, err := client.AuthLoginPassword("omarcosviniciusdev@gmail.com", "SecretBrla")
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to auth the account, error:\n\t%v", err)
		return
	}
	fmt.Println("[SENDED]\tAccount auth with successful")
	fmt.Printf("Your token: %v\n", token)

	webhookId, err := client.RegisterWebhook(token)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to register webhook, error:\n\t%v", err)
		return
	}
	fmt.Println("[SENDED]\tWebhook registred successful")
	fmt.Printf("Your webhook ID: %v\n", webhookId)
}
