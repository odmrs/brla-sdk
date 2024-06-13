package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
)

const sandbox string = "https://api.brla.digital:4567"

func main() {
	client := sdk.NewClient(sandbox)

	// Validate superuser and get tokenJWT
	tokenJWT, err := client.ValidateSuperAccount("castelobranco@brla.digital", "123456789")
	if err != nil {
		fmt.Printf("[error] \tfailed to auth the superuser account, error:\n\t%v", err)
		return
	}
	fmt.Println("\n[SENDED]\tSuper User account auth with successful")
	fmt.Printf("Your token: %v\n", tokenJWT)

	signatureKey, privateKey := signApiKey()

	if signatureKey == "" || privateKey == nil {
		return
	}

	/*
		publicKey, err := os.ReadFile("publickey.pem")
		if err != nil {
			fmt.Printf("[error] \terror to read publickey.pem, error:\n\t%v", err)
			return
		}

		apiKey, err := client.RegisterApiKey(tokenJWT, "brlaapikey", signatureKey, string(publicKey))
		if err != nil {
			fmt.Printf("\n[SENDED]\tfailed to register your api key, error: %v", err)
			return
		}

		fmt.Println("\n[SENDED]\tYour API KEY was registred with successful")
		fmt.Printf("Your API KEY: \n=> %v\n", apiKey)
	*/
	apiKey := "94ad0ddd-0c63-4198-9db5-4cdb1960770a"
	pixCode, err := client.CreateBRLABuyOrderApiKey(apiKey,
		"126.966.870-64",
		"0xd22c5da587869232dfdfe53a60873A2166027b51",
		"Polygon",
		5010,
		privateKey,
	)
	if err != nil {
		fmt.Printf("\n[ERROR] \tfailed to create PixCode ticket, error:\n\t%v", err)
		return
	}

	fmt.Println("\n[SENDED]\tCreate PixCode ticketwith successful")
	fmt.Printf("Your pixCode: \n%v\n", pixCode)
}
