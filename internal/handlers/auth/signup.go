package auth

import (
	authModel "leanGo/internal/models/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	user := new(authModel.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request",
		})
	}

	if user.Password != user.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Passwords do not match",
		})
	}

	existing := &authModel.User{}
	err := mgm.Coll(existing).First(bson.M{"email": user.Email}, existing)
	if err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  409,
			"message": "Email already exists",
		})
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hash)
	user.IsActive = true

	if err := mgm.Coll(user).Create(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Internal Server Error: Creating user failed",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  201,
		"message": "Signup successful",
	})
}
