package tests

import (
	"bytes"
	"encoding/json"
	"food-recipe-backend/internal/payments"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessPayment(t *testing.T) {
	payload := map[string]string{
		"amount":    "100",
		"currency":  "ETB",
		"email":     "customer@example.com",
		"reference": "chapa12345",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/payments/process", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(payments.ProcessPaymentHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Payment successful")
}
