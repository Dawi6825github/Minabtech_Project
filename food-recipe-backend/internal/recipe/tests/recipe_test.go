package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRecipe(t *testing.T) {
	payload := map[string]string{
		"title":        "Spaghetti Bolognese",
		"ingredients":  "Pasta, Tomato Sauce, Ground Beef",
		"instructions": "Cook pasta, add sauce, mix with beef",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/recipes", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(recipes.CreateRecipeHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Recipe created successfully")
}

func TestGetRecipes(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/recipes", nil)
	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(recipes.GetRecipesHandler)
	handler.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Spaghetti Bolognese")
}
