package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backend/internal/auth"
	"backend/models"
)

// Register handles user registration
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := auth.RegisterUser(&user)
	if err != nil {
		c.J
