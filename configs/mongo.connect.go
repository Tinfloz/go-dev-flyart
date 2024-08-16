package configs

import (
	"context"
	"fmt"
	"go-backend/structs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(mongoUri string) (*structs.DbConn, error) {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoUri).SetServerAPIOptions(serverApi)
	clientConnect, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	err = clientConnect.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to mongo!")
	dbConn := &structs.DbConn{
		Db: clientConnect.Database("flyart"),
	}
	return dbConn, nil
}
