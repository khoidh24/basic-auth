package configs

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	MongoURI  string = "mongodb://localhost:27017"
	Port      string = "3000"
	Domain    string = "http://localhost"
	JWTSecret string
	DBName    string = "habify"
)

func LoadConfig() {
	_ = godotenv.Load(".env")
	MongoURI = os.Getenv("MONGO_URI")
	Port = os.Getenv("PORT")
	Domain = os.Getenv("DOMAIN")
	JWTSecret = os.Getenv("JWT_SECRET")
	DBName = os.Getenv("DB_NAME")
}
