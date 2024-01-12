package controllers

import (
	"encoding/json"
	"my-golang-api/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AudioguideController struct {
	Collection *mongo.Collection
}

func (ac AudioguideController) GetAllAudioguides(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	var audioguides []models.Audioguide

	// Fetch audioguides from MongoDB
	cursor, err := ac.Collection.Find(r.Context(), bson.D{})
	if err != nil {
		http.Error(w, "Failed to fetch audioguides from MongoDB", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	// Iterate through the returned cursor and decode each document into the Audioguide struct
	for cursor.Next(r.Context()) {
		var audioguide models.Audioguide
		if err = cursor.Decode(&audioguide); err != nil {
			http.Error(w, "Failed to decode audioguide from MongoDB", http.StatusInternalServerError)
			return
		}
		audioguides = append(audioguides, audioguide)
	}

	// Check for errors during cursor iteration
	if err = cursor.Err(); err != nil {
		http.Error(w, "Failed to iterate over audioguides from MongoDB", http.StatusInternalServerError)
		return
	}

	// Convert to JSON and send the response
	json.NewEncoder(w).Encode(audioguides)
}

// CreateAudioguide creates a new audioguide.
func (ac *AudioguideController) CreateAudioguide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var audioguide models.Audioguide

	// Decode the incoming Audioguide json
	if err := json.NewDecoder(r.Body).Decode(&audioguide); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Set CreatedAt and UpdatedAt timestamps
	audioguide.CreatedAt = time.Now()
	audioguide.UpdatedAt = time.Now()

	// Insert the audioguide into MongoDB
	if _, err := ac.Collection.InsertOne(r.Context(), audioguide); err != nil {
		http.Error(w, "Failed to insert audioguide into MongoDB", http.StatusInternalServerError)
		return
	}

	// Send the response
	json.NewEncoder(w).Encode(audioguide)
}