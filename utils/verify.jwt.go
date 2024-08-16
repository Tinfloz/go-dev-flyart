package utils

import (
	"errors"
	"go-backend/structs"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyJwtToken(tokenString string) (*structs.JwtClaims, error) {
	var JwtSecret []byte
	token, err := jwt.ParseWithClaims(tokenString, &structs.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("could not parse jwt")
		}
		return JwtSecret, nil
	})
	if err != nil {
		return nil, errors.New("something failed at jwt verification")
	}
	if claims, ok := token.Claims.(*structs.JwtClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
