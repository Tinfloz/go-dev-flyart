package structs

import (
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtClaims struct {
	Id primitive.ObjectID
	jwt.RegisteredClaims
}
