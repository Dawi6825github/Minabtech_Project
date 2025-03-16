package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/google/uuid"
	"github.com/dgrijalva/jwt-go"
	"log"
)

// PaymentRequest represents the structure of a payment request
type PaymentRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Email    string  `json:"email"`
	FullName string  `json:"full_name"`
}

// PaymentResponse represents the response from Chapa API
type PaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    struct {
		CheckoutURL string `json:"checkout_url"`
	} `json:"data"`
}

// Initialize environment variables
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

// ProcessPayment handles payment requests
func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	// Decode request
	var req PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a unique transaction reference
	transactionRef := uuid.New().String()

	// Chapa API URL
	chapaURL := "https://api.chapa.co/v1/transaction/initialize"

	// Create payload
	payload := map[string]string{
		"amount":       fmt.Sprintf("%.2f", req.Amount),
		"currency":     req.Currency,
		"email":        req.Email,
		"first_name":   req.FullName,
		"tx_ref":       transactionRef,
		"callback_url": os.Getenv("CHAPA_CALLBACK_URL"),
	}

	// Marshal payload into JSON
	jsonPayload, _ := json.Marshal(payload)

	// Create HTTP request
	reqChapa, _ := http.NewRequest("POST", chapaURL, bytes.NewBuffer(jsonPayload))
	reqChapa.Header.Set("Authorization", "Bearer "+os.Getenv("CHAPA_SECRET_KEY"))
	reqChapa.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{}
	resp, err := client.Do(reqChapa)
	if err != nil {
		http.Error(w, "Failed to connect to payment gateway", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Parse response
	var paymentResp PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResp); err != nil {
		http.Error(w, "Invalid payment response", http.StatusInternalServerError)
		return
	}

	// Return checkout URL
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paymentResp)
}
