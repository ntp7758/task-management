package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	AuthRoleUser  string = "USER"
	AuthRoleAdmin string = "ADMIN"
)

type Auth struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt    time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt    time.Time          `json:"updatedAt" bson:"updatedAt"`
	Username     string             `json:"username" bson:"username"`
	Password     string             `json:"password" bson:"password"`
	Role         string             `json:"role" bson:"role"`
	RefreshToken string             `json:"refreshToken" bson:"refreshToken"`
}
