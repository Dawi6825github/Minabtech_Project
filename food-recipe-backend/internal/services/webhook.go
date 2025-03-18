package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// HasuraWebhookPayload defines the structure of the event payload
type HasuraWebhookPayload struct {
	Event struct {
		Op    string `json:"op"`
		Table string `json:"table"`
		Data  struct {
			New map[string]interface{} `json:"new"`
			Old map[string]interface{} `json:"old"`
		} `json:"data"`
	} `json:"event"`
}

// WebhookHandler processes Hasura event triggers
func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload HasuraWebhookPayload

	// Parse JSON payload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	// Log event
	log.Printf("Received webhook for %s operation on %s table\n", payload.Event.Op, payload.Event.Table)

	// Example: Handle "INSERT" event for the "recipes" table
	if payload.Event.Op == "INSERT" && payload.Event.Table == "recipes" {
		newRecipe := payload.Event.Data.New
		log.Printf("New Recipe Added: %v\n", newRecipe)
	}

	// Respond to Hasura
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message": "Webhook received successfully"}`)
}
