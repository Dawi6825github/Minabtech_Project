package controllers

import (
	"net/http"
	"backend/internal/payments"
	"github.com/gin-gonic/gin"
)

// ProcessPayment handles Chapa payment request
func ProcessPayment(c *gin.Context) {
	var paymentRequest payments.PaymentRequest
	if err := c.ShouldBindJSON(&paymentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentResponse, err := payments.ProcessPayment(paymentRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment failed"})
		return
	}

	c.JSON(http.StatusOK, paymentResponse)
}

// VerifyPayment checks Chapa payment status
func VerifyPayment(c *gin.Context) {
	transactionID := c.Param("transaction_id")

	status, err := payments.VerifyPayment(transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify payment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}
