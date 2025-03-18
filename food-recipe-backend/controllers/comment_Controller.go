package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"backend/internal/comments"
	"backend/models"
)

// CreateComment handles creating a new comment for a recipe
func CreateComment(c *gin.Context) {
	var comment models.Comment
	// Bind the incoming JSON to the Comment model
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the recipe ID from the URL parameter
	recipeID, _ := strconv.Atoi(c.Param("recipe_id"))
	comment.RecipeID = recipeID

	// Create the comment using the service layer
	err := comments.CreateComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment created successfully", "comment": comment})
}

// GetComments retrieves all comments for a specific recipe
func GetComments(c *gin.Context) {
	recipeID, _ := strconv.Atoi(c.Param("recipe_id"))

	// Fetch the comments from the service layer
	commentsList, err := comments.GetCommentsByRecipeID(recipeID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Comments not found"})
		return
	}

	c.JSON(http.StatusOK, commentsList)
}

// UpdateComment modifies an existing comment
func UpdateComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Param("comment_id"))
	var updatedComment models.Comment

	// Bind the incoming JSON to the updated comment model
	if err := c.ShouldBindJSON(&updatedComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the comment through the service layer
	err := comments.UpdateComment(commentID, &updatedComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

// DeleteComment removes a comment from the database
func DeleteComment(c *gin.Context) {
	commentID, _ := strconv.Atoi(c.Param("comment_id"))

	// Delete the comment through the service layer
	err := comments.DeleteComment(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
