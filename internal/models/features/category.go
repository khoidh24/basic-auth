package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string               `json:"categoryName" bson:"categoryName"`
	Description      string               `json:"description" bson:"description"`
	UserID           primitive.ObjectID   `json:"userId" bson:"userId"`
	NoteIDs          []primitive.ObjectID `json:"noteIds" bson:"noteIds"`
	IsPublic         bool                 `json:"isPublic" bson:"isPublic"`
	IsActive         bool                 `json:"isActive" bson:"isActive"`
}
