package controllers

import (
	"context"
	"go-backend/structs"
	"go-backend/utils"
	"log"
	"os"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductController struct {
	DbConn *structs.DbConn
}

func NewProductController(dbConn *structs.DbConn) *ProductController {
	return &ProductController{
		DbConn: dbConn,
	}
}

func (pc *ProductController) CreateNewDrawing(c *gin.Context) {
	var createProduct structs.Product
	err := c.ShouldBindJSON(&createProduct)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}
	cloudinaryUrl := "cloudinary://" + os.Getenv("API_KEY_CLOUDINARY") + ":" + os.Getenv("API_SECRET_CLOUDINARY") + "@" + os.Getenv("CLOUD_NAME")
	cld, errCld := cloudinary.NewFromURL(cloudinaryUrl)
	if errCld != nil {
		c.JSON(500, gin.H{
			"message": "Image upload failed",
		})
		return
	}
	reader, errChange := utils.GetB64Image(createProduct.Base64Str)
	if errChange != nil {
		c.JSON(500, gin.H{
			"message": errChange.Error(),
		})
		return
	}
	uploadRes, uploadErr := cld.Upload.Upload(
		context.Background(),
		reader,
		uploader.UploadParams{
			Folder:       os.Getenv("FOLDER"),
			ResourceType: "image",
		},
	)
	if uploadErr != nil {
		log.Fatal(uploadErr)
		c.JSON(500, gin.H{
			"message": "Could not upload image",
		})
		return
	}
	product := bson.D{
		{Key: "name", Value: createProduct.Name},
		{Key: "description", Value: createProduct.Description},
		{Key: "dimension", Value: createProduct.Dimension},
		{Key: "medium", Value: createProduct.Medium},
		{Key: "price", Value: createProduct.Price},
		{Key: "publicId", Value: uploadRes.PublicID},
		{Key: "secureUrl", Value: uploadRes.SecureURL},
		{Key: "uploadedAt", Value: time.Now()},
	}
	collection := pc.DbConn.Db.Collection("products")
	_, errMg := collection.InsertOne(context.TODO(), product)
	if errMg != nil {
		log.Printf("%v", errMg)
		c.JSON(500, gin.H{
			"message": errMg.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"message": "Painting uploaded!",
	})
}
