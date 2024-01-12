package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Audioguide struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Title       string             `json:"title"`
	AudioURL    string             `json:"audio_url"`
	Description string             `json:"description"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}