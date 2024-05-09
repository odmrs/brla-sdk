package main

import (
	"fmt"

	"github.com/odmrs/brla-sdk/pkg/models"
	"github.com/odmrs/brla-sdk/pkg/sdk"
)

const (
	// Sandbox Environment
	sandbox string = "https://api.brla.digital:4567"

	// Production Environment
	// production string = "https://api.brla.digital:5567"
)

func main() {
	// Create the sdk client
	client := sdk.NewClient(sandbox)

	// Example of creation of account
	account := models.NewAccount(
		"email@example.com",
		"senha",
		"senha",
		"12321321321",
		"CPF",
		"marcos",
		"906.089.050-70",
		"2004-jan-02",
		models.Address{
			Cep:        "a",
			City:       "a",
			State:      "a",
			Street:     "a",
			Number:     "a",
			District:   "a",
			Complement: "a",
		},
	)

	// Send the request
	err := client.CreateAccount(account)
	if err != nil {
		fmt.Printf("Failed to create account, error:\n\t%v", err)
		return
	}

	fmt.Println("Account creation successful")

	// Validate account
	err = client.ValidateAccount("email@example.com", "token")
	if err != nil {
		fmt.Printf("error validate creating account: %v", err)
	}
}
