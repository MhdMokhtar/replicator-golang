// main.go
package main

import (
	"context"
	"log"
	"my-golang-api/controllers"
	"my-golang-api/routes"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
   // Initialize MongoDB client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	// Set up a new router
	router := mux.NewRouter()

	// Set up controllers
	audioguideController := &controllers.AudioguideController{
		Collection: client.Database("my-golang-api").Collection("audioguides"),
	}

	buildingController := &controllers.BuildingController{
		Collection: client.Database("my-golang-api").Collection("buildings"),
	}

	// Set up routes for the Audioguide model
	routes.SetAudioguideRoutes(router, audioguideController)

	// Set up routes for the Building model
	routes.SetBuildingRoutes(router, buildingController)

	// Start the server
	log.Println("Server is running on :9000")
	log.Fatal(http.ListenAndServe(":9000", router))
}
