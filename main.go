// main.go
package main

import (
	"context"
	"log"
	"my-golang-api/controllers"
	"my-golang-api/routes"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Load environment variables from .env file
	loadEnv()

	// Initialize MongoDB client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
 
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
	   log.Fatal(err)
	}
 
	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
	   log.Fatal(err)
	}

	// Set up a new router
	router := mux.NewRouter()

	// Set up controllers
	audioguideController := &controllers.AudioguideController{
		Collection: client.Database(os.Getenv("anno-amsterdam")).Collection("audioguides"),
	}

	buildingController := &controllers.BuildingController{
		Collection: client.Database(os.Getenv("anno-amsterdam")).Collection("buildings"),
	}

	// Set up routes for the Audioguide model
	routes.SetAudioguideRoutes(router, audioguideController)

	// Set up routes for the Building model
	routes.SetBuildingRoutes(router, buildingController)

	// Start the server on the configured port
	apiPort := os.Getenv("API_PORT")
	log.Printf("Server is running on :%s\n", apiPort)
	log.Fatal(http.ListenAndServe(":"+apiPort, router))
}
