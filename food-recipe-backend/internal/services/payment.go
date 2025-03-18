package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// ChapaPaymentRequest defines the request payload for Chapa API
type ChapaPaymentRequest struct {
	Amount       float64 `json:"amount"`
	Currency     string  `json:"currency"`
	Email        string  `json:"email"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	TxRef        string  `json:"tx_ref"`
	CallbackURL  string  `json:"callback_url"`
	ReturnURL    string  `json:"return_url"`
	CustomFields struct {
		OrderID string `json:"order_id"`
	} `json:"custom_fields"`
}

// ChapaPaymentResponse defines the response payload from Chapa
type ChapaPaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}

// ProcessChapaPayment initiates a payment request to Chapa
func ProcessChapaPayment(amount float64, email, firstName, lastName, orderID string) (string, error) {
	apiURL := "https://api.chapa.co/v1/transaction/initialize"
	apiKey := os.Getenv("CHAPA_SECRET_KEY")

	requestData := ChapaPaymentRequest{
		Amount:      amount,
		Currency:    "ETB",
		Email:       email,
		FirstName:   firstName,
		LastName:    lastName,
		TxRef:       orderID,
		CallbackURL: "https://yourdomain.com/payment/callback",
		ReturnURL:   "https://yourdomain.com/payment/success",
	}

	// Convert request data to JSON
	requestBody, err := json.Marshal(requestData)
	if err != nil {
		return "", fmt.Errorf("failed to create payment request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// Set Headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to process payment: %v", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	// Parse response
	var paymentResponse ChapaPaymentResponse
	if err := json.Unmarshal(body, &paymentResponse); err != nil {
		return "", fmt.Errorf("failed to parse payment response: %v", err)
	}

	if paymentResponse.Status != "success" {
		return "", fmt.Errorf("payment request failed: %s", paymentResponse.Message)
	}

	return paymentResponse.Data.CheckoutURL, nil
}
