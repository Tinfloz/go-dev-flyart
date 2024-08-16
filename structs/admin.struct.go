package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type AdminUser struct {
	Id       primitive.ObjectID
	Email    string
	Password string
	UserType string
}
