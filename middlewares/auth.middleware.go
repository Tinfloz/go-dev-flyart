package middlewares

import (
	"context"
	"fmt"
	"go-backend/structs"
	"go-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MiddlewareController struct {
	DbConn *structs.DbConn
}

func NewMiddlewareController(dbConn *structs.DbConn) *MiddlewareController {
	return &MiddlewareController{
		DbConn: dbConn,
	}
}

func AuthMiddleware(dbConn *structs.DbConn) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(400, gin.H{"message": "Bearer token absent"})
			c.Abort()
			return
		}
		bearerToken := strings.Split(token, " ")
		if bearerToken[0] != "Bearer" {
			c.JSON(400, gin.H{"message": "Invalid token format"})
			c.Abort()
			return
		}
		claims, err := utils.VerifyJwtToken(bearerToken[1])
		if err != nil {
			c.JSON(400, gin.H{
				"message": "Something went wrong in token authorization",
			})
			c.Abort()
			return
		}
		var result struct {
			Id       primitive.ObjectID `bson:"_id"`
			Email    string             `bson:"email"`
			UserType string             `bson:"userType"`
		}
		projection := bson.M{"password": 0}
		errDb := dbConn.Db.Collection("admin_users").FindOne(context.Background(), bson.M{"_id": claims.Id}, options.FindOne().SetProjection(projection)).Decode(&result)
		if errDb != nil {
			c.JSON(400, gin.H{
				"message": "User not found",
			})
			c.Abort()
			return
		}
		fmt.Printf("%v", result)
		c.Set("userClaims", result)
		c.Next()
	}
}

func CheckAdminStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, ok := c.Get("userClaims")
		if !ok {
			c.JSON(500, gin.H{
				"message": "Something went wrong",
			})
			c.Abort()
			return
		}
		user, okUser := userClaims.(struct {
			Id       primitive.ObjectID `bson:"_id"`
			Email    string             `bson:"email"`
			UserType string             `bson:"userType"`
		})
		if !okUser {
			c.JSON(500, gin.H{"message": "User claims type assertion failed"})
			c.Abort()
			return
		}
		if user.UserType != "admin" {
			c.JSON(403, gin.H{"message": "Unauthorised access"})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}

}
