package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Event struct {
	Acknowledged bool      `json:"acknowledged"`
	CreatedAt    int64     `json:"createdAt"`
	ID           string    `json:"id"`
	Subscription string    `json:"subscription"`
	UserID       string    `json:"userId"`
	Data         EventData `json:"data"`
}

type EventData struct {
	Amount         int    `json:"amount,omitempty"`
	Chain          string `json:"chain,omitempty"`
	DiscountedFee  int    `json:"discountedFee,omitempty"`
	ID             string `json:"id,omitempty"`
	Reason         string `json:"reason,omitempty"`
	ReferenceLabel string `json:"referenceLabel,omitempty"`
	Status         string `json:"status,omitempty"`
	TaxID          string `json:"taxId,omitempty"`
	Tx             string `json:"tx,omitempty"`
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	var event Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	eventJSON, err := json.MarshalIndent(event, "", " ")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	fmt.Printf("Received event: \n%s\n", eventJSON)
}

func main() {
	http.HandleFunc("/", handleWebhook)
	fmt.Printf("[INFO] - Listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
