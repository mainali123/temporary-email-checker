package main

import (
	"encoding/json"
	"net/http"
	"temp_mail/email"

	"github.com/gorilla/mux"
)

var version = "1.0.0"

// RegisterRoutes sets up all the routes for the application
func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Define GET route
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{
			"message": "Welcome to the Go web server!",
			"version": version,
		}
		json.NewEncoder(w).Encode(response)
	})

	// Define POST route for email validation
	r.HandleFunc("/validate-email", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Parse the request body
		var requestBody struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Check if the email is a Gmail address
		isTemp := email.IsTemporary(requestBody.Email)

		// Respond with the result
		response := map[string]interface{}{
			"email":       requestBody.Email,
			"isTemporary": isTemp,
		}
		json.NewEncoder(w).Encode(response)
	}).Methods("POST")

	return r
}
