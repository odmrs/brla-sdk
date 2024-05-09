package sdk

import (
	"testing"

	"github.com/odmrs/brla-sdk/models"
)

const (
	sandbox string = "https://api.brla.digital:4566"
)

func TestCreateAccount(t *testing.T) {
	client := NewClient(sandbox)

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
		t.Errorf("error creating account: %v", err)
	}
}

func TestValidateAccount(t *testing.T) {
	client := NewClient(sandbox)
	err := client.ValidateAccount("email@example.com", "token")
	if err != nil {
		t.Errorf("error validate creating account: %v", err)
	}
}
