package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
	"github.com/odmrs/brla-sdk/models"
)

const (
	// sandbox environment
	sandbox string = "https://api.brla.digital:4567"

	// production environment
	// production string = "https://api.brla.digital:5567"
)

func main() {
	// Create the sdk client
	client := sdk.NewClient(sandbox)

}
