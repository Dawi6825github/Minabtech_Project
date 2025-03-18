package payments

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"your_project/internal/utils"
)

type PaymentRequest struct {
	Amount     float64 `json:"amount" binding:"required"`
	RecipeID   string  `json:"recipe_id" binding:"required"`
	UserID     string  `json:"user_id" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
}

func HandlePaymentRequest(c *gin.Context) {
	var paymentReq PaymentRequest
	if err := c.ShouldBindJSON(&paymentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to handle payment process
	paymentResponse, err := ProcessPayment(paymentReq.Amount, paymentReq.RecipeID, paymentReq.UserID, paymentReq.PaymentMethod)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentResponse)
}

func ProcessPayment(amount float64, recipeID, userID, paymentMethod string) (map[string]interface{}, error) {
	// Placeholder for Chapa integration
	// Typically, you'll use Chapa's SDK here to initiate the payment.

	// Example response from Chapa API
	paymentResponse := map[string]interface{}{
		"status": "success",
		"message": "Payment successful",
		"data": map[string]interface{}{
			"amount":        am
