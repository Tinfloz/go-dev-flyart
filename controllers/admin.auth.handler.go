package controllers

import (
	"context"
	"go-backend/structs"
	"go-backend/utils"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AdminAuthController struct {
	DbConn *structs.DbConn
}

func NewAdminAuthController(dbConn *structs.DbConn) *AdminAuthController {
	return &AdminAuthController{
		DbConn: dbConn,
	}
}

func (ac *AdminAuthController) AdminLogin(c *gin.Context) {
	var loginRequest structs.LoginBody
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	collection := ac.DbConn.Db.Collection("admin_users")
	var result struct {
		Id       primitive.ObjectID `bson:"_id"`
		Email    string             `bson:"email"`
		Password string             `bson:"password"`
	}
	err := collection.FindOne(context.TODO(), bson.M{"email": loginRequest.Email}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(404, gin.H{
				"message": "No such users found",
			})
		} else {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		}
		return
	}
	errBcrypt := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(loginRequest.Password))
	if errBcrypt != nil {
		if errBcrypt == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(400, gin.H{
				"message": "Passwords don't match",
			})
		} else {
			c.JSON(500, gin.H{
				"message": errBcrypt.Error(),
			})
		}
		return
	}
	token, errJwt := utils.GetJwtToken(result.Id)
	if errJwt != nil {
		c.JSON(500, gin.H{
			"message": errJwt.Error(),
		})
	}
	c.JSON(200, gin.H{
		"token": token,
		"email": result.Email,
	})
}

func (ac *AdminAuthController) AddAdminUsers(c *gin.Context) {
	var createUserBody structs.LoginBody
	if err := c.ShouldBindJSON(&createUserBody); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	passwordSet, errBc := bcrypt.GenerateFromPassword([]byte(createUserBody.Password), 12)
	if errBc != nil {
		log.Fatal(errBc)
		c.JSON(500, gin.H{
			"message": errBc.Error(),
		})
		return
	}
	user := bson.D{
		{Key: "email", Value: createUserBody.Email},
		{Key: "password", Value: string(passwordSet)},
		{Key: "userType", Value: "admin"},
	}
	collection := ac.DbConn.Db.Collection("admin_users")
	_, errMg := collection.InsertOne(context.TODO(), user)
	if errMg != nil {
		log.Fatal(errMg)
		c.JSON(500, gin.H{
			"message": errMg.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "User Created",
		"email":   createUserBody.Email,
	})
}
