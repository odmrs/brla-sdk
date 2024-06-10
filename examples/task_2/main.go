package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
)

const baseUrl = "https://api.brla.digital:4567"

func main() {
	client := sdk.NewClient(baseUrl)
	/*	account := models.NewAccount(
				"omarcosviniciusdev@gmail.com",
				"SecretBrla",
				"SecretBrla",
				"13321321321",
				"CPF",
				"Marcos",
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

		err := client.ValidateAccount("omarcosviniciusdev@gmail.com")
		if err != nil {
			fmt.Printf("[ERROR] \tfailed to validate account, error:\n\t%v", err)
			return
		}

		fmt.Println("[SENDED]\tAccount validated with successful")
	*/

	// Auth user with login and password -> validate endpoint
	token, err := client.AuthLoginPassword("omarcosviniciusdev@gmail.com", "SecretBrla")
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to auth the account, error:\n\t%v", err)
		return
	}
	fmt.Println("[SENDED]\tAccount auth with successful")
	fmt.Printf("Your token: %v\n", token)

	// Generate PIX code in payin endpoint
	pixCode, err := client.GeneratesPayinCode(token, "0", "", "")
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to auth the account, error:\n\t%v", err)
		return
	}
	fmt.Println("[SENDED]\tGenerate PIX with successful")
	fmt.Printf("Your pix code: %v\n", pixCode)

	// Generate a fake payment inside sandbox mock
	err = client.GeneratePaymantSandbox(token)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to generate payment on sandbox, error: \n\t%v", err)
		return
	}
	fmt.Println("[SENDED]\tPayment simulated with successful")

	body, err := client.ShowHistoryPayIn(token)
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to show history of payments, error: \n\t%v", err)
		return
	}
	fmt.Println("[SENDED]\tGet all history of payment with successful\n")
	fmt.Print(body)
}
