package structs

import "go.mongodb.org/mongo-driver/mongo"

type DbConn struct {
	Db *mongo.Database
}
