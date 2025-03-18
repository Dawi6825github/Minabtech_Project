package tests

import (
	"bytes"
	"encoding/json"
	"food-recipe-backend/internal/auth"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	payload := map[string]string{
		"email":    "testuser@example.com",
		"password": "securePassword123",
	}
	jsonPayload, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/api/auth/signup", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "applic
