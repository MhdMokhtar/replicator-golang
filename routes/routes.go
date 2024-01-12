package routes

import (
	"my-golang-api/controllers"

	"github.com/gorilla/mux"
)

// SetAudioguideRoutes sets up routes for the Audioguide model.
func SetAudioguideRoutes(router *mux.Router, audioguideController *controllers.AudioguideController) {
	router.HandleFunc("/audioguides", audioguideController.GetAllAudioguides).Methods("GET")
	router.HandleFunc("/audioguides", audioguideController.CreateAudioguide).Methods("POST")
}

// SetBuildingRoutes sets up routes for the Building model.
func SetBuildingRoutes(router *mux.Router, buildingController *controllers.BuildingController) {
	router.HandleFunc("/buildings", buildingController.GetAllBuildings).Methods("GET")
	router.HandleFunc("/buildings", buildingController.CreateBuilding).Methods("POST")
}