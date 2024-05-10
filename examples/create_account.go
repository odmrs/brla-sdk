package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
	"github.com/odmrs/brla-sdk/models"
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

	err := client.CreateAccount(account)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to create account, error:\n\t%v", err)
		return
	}
	fmt.Println("[SENDED]\tAccount creation successful")
}
