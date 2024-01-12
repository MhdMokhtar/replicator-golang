package controllers

import (
	"encoding/json"
	"my-golang-api/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// BuildingController handles HTTP requests for the Building model.
type BuildingController struct {
	Collection *mongo.Collection
}

// GetAllBuildings returns all buildings.
func (bc *BuildingController) GetAllBuildings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var buildings []models.Building

	// Fetch buildings from MongoDB
	cursor, err := bc.Collection.Find(r.Context(), bson.D{})
	if err != nil {
		http.Error(w, "Failed to fetch buildings from MongoDB", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(r.Context())

	// Iterate through the returned cursor and decode each document into the Building struct
	for cursor.Next(r.Context()) {
		var building models.Building
		if err = cursor.Decode(&building); err != nil {
			http.Error(w, "Failed to decode building from MongoDB", http.StatusInternalServerError)
			return
		}
		buildings = append(buildings, building)
	}

	// Check for errors during cursor iteration
	if err = cursor.Err(); err != nil {
		http.Error(w, "Failed to iterate over buildings from MongoDB", http.StatusInternalServerError)
		return
	}

	// Convert to JSON and send the response
	json.NewEncoder(w).Encode(buildings)
}

func (bc *BuildingController) CreateBuilding(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var building models.Building

	// Decode the request body into a Building struct
	err := json.NewDecoder(r.Body).Decode(&building)
	if err != nil {
		// Handle invalid request body
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	building.ID = primitive.NewObjectID()
	building.Audioguides = make([]models.Audioguide, 0)

	// Insert building into MongoDB
	if _, err := bc.Collection.InsertOne(r.Context(), building); err != nil {
		http.Error(w, "Failed to insert building into MongoDB", http.StatusInternalServerError)
		return
	}

	// Convert to JSON and send the response
	json.NewEncoder(w).Encode(building)
}