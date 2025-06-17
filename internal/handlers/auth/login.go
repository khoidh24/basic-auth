package auth

import (
	configs "leanGo/config"
	"leanGo/internal/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	req := new(models.User)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid login")
	}

	user := &models.User{}
	err := mgm.Coll(user).First(bson.M{"email": req.Email}, user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Wrong email or password")
	}

	if compareErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); compareErr != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Wrong email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   user.ID.Hex(),
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	})
	t, err := token.SignedString(configs.JWTSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Token generation failed")
	}

	return c.JSON(fiber.Map{"token": t})
}
