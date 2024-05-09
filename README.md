# BRLA Digital SDK for GOlang

This repository contains the UNOFFICIAL SDK for BRLA Digital, allowing developers to easily integrate their applications with the BRLA Digital platform. This SDK is written in GoLang and provides an easy-to-use interface for accessing BRLA Digital's API resources.

## API Documentation

For more details on how to use the official API, consult the official documentation available on the official [BRLA Digital](https://brla-account-api.readme.io/reference/welcome) website. For the SDK documentation, keep reading this repository.

## Pre-requisites

To install the SDK, run the following command:
```bash
go get github.com/odmrs/brla-sdk-go
```

## Usage
example of creating a new account:

```go
import (
	"github.com/odmrs/brla-sdk/models"
	sdk "github.com/odmrs/brla-sdk"
)

const (
	sandbox string "https://api.brla.digital:4567"
)

func main() {
	// Create client
	client := sdk.NewCLient(sandbox)

	// Prepare the requires to create account
	account := models.NewAccounta(
		"email@example.com",	// EMAIL
		"your-password",	// PASSWORD
		"your-password",	// CONFIRM PASSWORD
		"99999999999",		// NUMBER
		"CPF",			// TAXIDTYPE 
		"BRLA Digital",  	// BRLA Digital
		"999.999.999-99", 	// CPF 
		"2004-jan-02", 		// Birth Day
		models.Address{
			Cep:        "cep",
			City:       "city",
			State:      "state",
			Street:     "street",
			Number:     "number",
			District:   "district",
			Complement: "complement",
		},
	)

	// Finally, send the request

	err := client.CreateAccount(account)

	.... handler the erros here ....
}
```
