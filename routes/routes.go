package routes

import (
	"my-golang-api/controllers"

	"github.com/gorilla/mux"
)

// SetAudioguideRoutes sets up routes for the Audioguide model.
func SetAudioguideRoutes(router *mux.Router, audioguideController *controllers.AudioguideController) {
	router.HandleFunc("/api/audioguides", audioguideController.GetAllAudioguides).Methods("GET")
}

// SetBuildingRoutes sets up routes for the Building model.
func SetBuildingRoutes(router *mux.Router, buildingController *controllers.BuildingController) {
	router.HandleFunc("/api/buildings", buildingController.GetAllBuildings).Methods("GET")
	router.HandleFunc("/api/buildings/{id}", buildingController.GetBuildingByID).Methods("GET")
}