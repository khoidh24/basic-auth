package database

import (
	configs "leanGo/config"
	"log"
	"sync"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once

func ConnectMongo() {
	once.Do(func() {
		err := mgm.SetDefaultConfig(nil, configs.DBName, options.Client().ApplyURI(configs.MongoURI))
		if err != nil {
			log.Fatal("Could not connect to MongoDB: ", err)
		}
		log.Println("Connected to MongoDB")
	})
}
