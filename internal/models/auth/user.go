package models

import (
	"github.com/kamva/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Email            string `json:"email" bson:"email"`
	Password         string `json:"password,omitempty" bson:"password"`
	ConfirmPassword  string `json:"confirmPassword,omitempty" bson:"-"`
	IsActive         bool   `json:"isActive" bson:"isActive"`
}
