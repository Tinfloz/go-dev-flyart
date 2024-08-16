package configs

import (
	"fmt"
	"go-backend/structs"
	"os"

	"github.com/joho/godotenv"
)

func GetDotEnv() *structs.ConfigEnv {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	config := &structs.ConfigEnv{
		Port:      os.Getenv("PORT"),
		MongoUri:  os.Getenv("MONGO_URI"),
		JwtSecret: os.Getenv("JWT_SECRET"),
	}
	return config
}
