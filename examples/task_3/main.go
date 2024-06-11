package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
)

const baseUrl = "https://api.brla.digital:4567"

func main() {
	client := sdk.NewClient(baseUrl)
	token, err := client.AuthLoginPassword("omarcosviniciusdev@gmail.com", "SecretBrla")
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to auth the account, error:\n\t%v", err)
		return
	}
	fmt.Println("[SENDED]\tAccount auth with successful")
	fmt.Printf("Your token: %v\n\n", token)

	id, err := client.CreatePayoutOrder(
		token,
		"omarcosviniciusdev@gmail.com",
		"906.089.050-70",
		"",
		"Marcos",
		"12345678",
		"0001",
		"123456789",
		"CACC",
		80,
	)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to create payout order, error:\n\t%v", err)
		return
	}
	fmt.Println("[SENDED]\tCreate PayOut Order with successful")
	fmt.Printf("Your order ID: %v\n\n", id)

	history, err := client.ShowPayoutHistory(token)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to create payout order, error:\n\t%v", err)
		return
	}

	fmt.Println("[SENDED]\tGet PayOut history with successful")
	fmt.Printf("Your payout history: %v\n\n", history)
}
