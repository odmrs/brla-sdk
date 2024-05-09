package sdk

import (
	"testing"

	"github.com/odmrs/brla-sdk/pkg/models"
)

const (
	sandbox string = "https://api.brla.digital:4567"
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
		"password123",
		"password123",
		"9999999999999",
		"CPF",
		"Testing Name",
		"479.280.460-40",
		"2004-jan-02",
		address,
	)

	err := client.CreateAccount(account)

	if err != nil {
		t.Fatalf("error creating account: %v", err)
	}
}
