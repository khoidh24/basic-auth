package auth

import (
	configs "leanGo/config"
	authModel "leanGo/internal/models/auth"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	req := new(authModel.User)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  400,
			"message": "Invalid request",
		})
	}

	user := &authModel.User{}
	err := mgm.Coll(user).First(bson.M{"email": req.Email}, user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	if compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); compareErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  401,
			"message": "Unauthorized",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID.Hex(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})
	t, err := token.SignedString(configs.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  500,
			"message": "Token generation failed",
		})
	}

	return c.JSON(fiber.Map{
		"status": 200,
		"token":  t,
	})
}
