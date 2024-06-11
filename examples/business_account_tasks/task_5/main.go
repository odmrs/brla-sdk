package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
)

const sandbox string = "https://api.brla.digital:4567"

func main() {
	client := sdk.NewClient(sandbox)
	// Auth user with login and password -> validate endpoint
	tokenJWT, err := client.AuthLoginPassword("omarcosviniciusdev@gmail.com", "SecretBrla")
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to auth the account, error:\n\t%v", err)
		return
	}
	fmt.Println("\n[SENDED]\tAccount auth with successful")
	fmt.Printf("Your token: %v\n", tokenJWT)

	// Get quote token:
	quoteToken, err := client.QuoteToken(
		tokenJWT,
		"swap",
		"BRLA",
		"USDT",
		"Polygon",
		"true",
		"4500",
	)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to get quote token, error:\n\t%v", err)
		return
	}
	fmt.Println("\n[SENDED]\tReceive quote toke with successful")
	fmt.Println("[INFO] Your token expiry in 10 seconds...")
	fmt.Printf("Your token: %v\n", quoteToken)

	// Using this quoteToken to simulate a convert
	id, err := client.ConvertBetweenCurrencies(
		tokenJWT,
		quoteToken,
		"0x599d804403ec5e7A48aa1C685A7BFb0C20B9f24d",
		"0",
	)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to convert between currencies, error:\n\t%v", err)
		return
	}
	fmt.Println("\n[SENDED]\tReceive id by convert between currencies with successful")
	fmt.Printf("Your convert ID: %v\n", id)

	history, err := client.HistoryConversionOperations(tokenJWT)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to get your history of conversions, error:\n\t%v", err)
		return
	}

	fmt.Println("\n[SENDED]\tget history of conversions with successful")
	fmt.Printf("Your history: %v\n", history)
}
