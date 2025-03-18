package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"backend/internal/users"
)

// GetUser retrieves user profile
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := users.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser updates user profile
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var userUpdate users.UserUpdateRequest

	if err := c.ShouldBindJSON(&userUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := users.UpdateUserProfile(id, &userUpdate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
