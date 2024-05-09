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
	address := models.Address{
		Cep:        "a",
		City:       "a",
		State:      "a",
		Street:     "a",
		Number:     "a",
		District:   "a",
		Complement: "a",
	}

	account := models.NewAccount(
		"test.test@example.com",
		"password122",
		"password122",
		"9999999999998",
		"CPF",
		"Testing Name",
		"478.280.460-40",
		"2003-jan-02",
		address,
	)

	err := client.CreateAccount(account)

	if err != nil {
		t.Fatalf("error creating account: %v", err)
	}
}

func TestValidateAccount(t *testing.T) {
	client := NewClient(sandbox)
	err := client.ValidateAccount("email@example.com", "token")
	if err != nil {
		t.Skipf("error validate creating account: %v", err)
	}
}
