package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
	"github.com/odmrs/brla-sdk/models"
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
		fmt.Printf("[ERROR] \tfailed to create account, error:\n\t%v", err)
		return
	}
	fmt.Println("[SENDED]\tAccount creation successful")

	// Validate account
	err = client.ValidateAccount("email@example.com", "token")
	if err != nil {
		fmt.Printf("[ERROR] \terror validate creating account: %v\n", err)
	}
	fmt.Println("[SENDED]\tValidate Account sended with success")

	// Authenticates with login and password
	err = client.AuthLoginPassword("email@example.com", "password")
	if err != nil {
		fmt.Printf("[ERROR] \terror validate authenticates account: %v\n", err)
	}
	fmt.Println("[SENDED]\tAuthentication sended with success")

	// Reset password
	err = client.ResetPassword("email@example.com")
	if err != nil {
		fmt.Printf("[ERROR] \terror reset password: %v\n", err)
	}
	fmt.Println("[SENDED]\tReset password sended with success")

	// Concludes password reset process
	err = client.ConcludesResetPassword("tokenblablabla", "email@gmail.com")
	if err != nil {
		fmt.Printf("[ERROR] \terror concludes reset password: %v\n", err)
	}
	fmt.Println("[SENDED]\tConcludes password sended with success")

	// Change account password
	err = client.ChangePassword("currentPassword", "newpassword", "newpassword", "blablabla")

	fmt.Println("[SENDED]\tChange password sended with success")

	if err != nil {
		fmt.Printf("[ERROR] \terror to change your account password: %v\n", err)
	}

	// Logs account out
	err = client.LoggoutAccount("JWT TOKEN HERE")

	fmt.Println("[SENDED]\tLoggout account sended with success")

	if err != nil {
		fmt.Printf("[ERROR] \terror to loggout your account: %v\n", err)
	}

}
