package utils

import (
	auth "leanGo/internal/models/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserByEmailFromContext(c *fiber.Ctx) (*auth.User, error) {
	email := c.Locals("email")
	if email == nil {
		return nil, fiber.ErrUnauthorized
	}

	user := &auth.User{}
	if err := mgm.Coll(user).First(bson.M{"email": email}, user); err != nil {
		return nil, err
	}

	return user, nil
}
