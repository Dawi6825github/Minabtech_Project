package users

import (
	"food-recipe-backend/internal/models"
	"food-recipe-backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	if err := h.Service.CreateUser(&user); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.RespondJSON(c, http.StatusCreated, user)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.Service.GetUserByID(uint(id))
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found")
		return
	}

	utils.RespondJSON(c, http.StatusOK, user)
}
