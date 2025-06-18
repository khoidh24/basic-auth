package models

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Note struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string             `json:"noteTitle" bson:"noteTitle"`
	Content          string             `json:"content" bson:"content"`
	CoverImage       string             `json:"coverImage" bson:"coverImage"`
	UserId           primitive.ObjectID `json:"userId" bson:"userId"`
	CategoryID       primitive.ObjectID `json:"categoryId" bson:"categoryId"`
	IsActive         bool               `json:"isActive" bson:"isActive"`
}
