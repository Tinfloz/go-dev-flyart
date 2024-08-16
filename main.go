package main

import (
	"fmt"
	"go-backend/configs"
	"go-backend/routes"
	"go-backend/structs"
	"log"

	"github.com/gin-gonic/gin"
)

func setup() (*structs.DbConn, error) {
	config := configs.GetDotEnv()
	dbConn, err := configs.MongoConnect(config.MongoUri)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(dbConn)
	return dbConn, nil
}

func main() {
	// fmt.Println(dbConn)
	dbConn, err := setup()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Gin server up and running!",
		})
	})
	routes.SetUpRoutes(r, dbConn)
	r.Run()
}
