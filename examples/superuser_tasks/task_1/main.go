package main

import (
	"fmt"

	sdk "github.com/odmrs/brla-sdk"
)

const sandbox string = "https://api.brla.digital:4567"

func main() {
	client := sdk.NewClient(sandbox)

	tokenJWT, err := client.ValidateSuperAccount("castelobranco@brla.digital", "123456789")
	if err != nil {
		fmt.Printf("[ERROR] \tfailed to auth the SuperUser account, error:\n\t%v", err)
		return
	}
	fmt.Println("\n[SENDED]\tSuper User account auth with successful")
	fmt.Printf("Your token: %v\n", tokenJWT)
	/*
		id, err := client.SimulateKYCVerify(tokenJWT, "126.966.870-64", "2004-jan-02", "Marcos")
		if err != nil {
			fmt.Printf("\n[ERROR] \tfailed to simulate KYC, error:\n\t%v", err)
			return
		}

		fmt.Println("\n[SENDED]\tSimulate KYC with successful")
		fmt.Printf("Your ID: %v\n", id)
	*/
	pixCode, err := client.CreatePixSuperUserTicket(
		tokenJWT,
		"126.966.870-64",
		"0xd22c5da587869232dfdfe53a60873A2166027b51",
		"Polygon",
		5010,
	)
	if err != nil {
		fmt.Printf("\n[ERROR] \tfailed to create PixCode ticket, error:\n\t%v", err)
		return
	}

	fmt.Println("\n[SENDED]\tCreate PixCode ticketwith successful")
	fmt.Printf("Your pixCode: \n%v\n", pixCode)
}
