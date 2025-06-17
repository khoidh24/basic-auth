package configs

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	MongoURI  string = "mongodb://localhost:27017"
	Port      string = "3000"
	JWTSecret string
)

func LoadConfig() {
	_ = godotenv.Load(".env")
	MongoURI = os.Getenv("MONGO_URI")
	Port = os.Getenv("PORT")
	JWTSecret = os.Getenv("JWT_SECRET")
}
