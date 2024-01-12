package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Building represents the model for buildings.
type Building struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	PreviewImageURL  string             `json:"preview_image_url"`
	Latitude         float64            `json:"latitude"`
	Longitude        float64            `json:"longitude"`
	Address          string             `json:"address"`
	ConstructionYear int                `json:"construction_year"`
	TypeOfUse        string             `json:"type_of_use"`
	Tags             map[string]interface{} `json:"tags"`
	Description      string             `json:"description"`
	ImageURLs        []string           `json:"image_urls"`
	Timeline         map[string]string  `json:"timeline"`
	Audioguides      []Audioguide       `json:"audioguides"`
}