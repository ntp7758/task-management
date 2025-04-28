package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	UserRoleEmployee string = "EMPLOYEE"
	UserRoleManager  string = "MANAGER"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
	AuthId    string             `json:"authId" bson:"authId"`
	Role      string             `json:"role" bson:"role"`
	Task      []string           `json:"task" bson:"task"`
}
