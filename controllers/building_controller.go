package controllers

import (
	"context"
	"encoding/json"
	"my-golang-api/models"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

func (bc *BuildingController) GetBuildingByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Extract building ID from request parameters
	params := mux.Vars(r)
	buildingID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		// Handle invalid ID
		http.Error(w, "Invalid Building ID", http.StatusBadRequest)
		return
	}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB query to find a building by ID
	filter := bson.M{"_id": buildingID}
	var building models.Building

	// Execute the query and decode the result into the building variable
	err = bc.Collection.FindOne(ctx, filter).Decode(&building)
	if err != nil {
		// Handle errors, e.g., building not found
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Building not found", http.StatusNotFound)
		} else {
			// Handle other errors
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	// Convert to JSON and send the response
	json.NewEncoder(w).Encode(building)
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