package utils

import (
	"go-backend/structs"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetJwtToken(id primitive.ObjectID) (string, error) {
	expirationTime := time.Now().Add(120 * time.Hour)
	claims := &structs.JwtClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	var JwtSecret []byte
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tokenString, nil
}
