package routes

import (
	"go-backend/controllers"
	"go-backend/middlewares"
	"go-backend/structs"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine, dbConn *structs.DbConn) {
	adminAuthController := controllers.NewAdminAuthController(dbConn)
	productController := controllers.NewProductController(dbConn)
	authMiddleware := middlewares.AuthMiddleware(dbConn)
	adminAuthApi := r.Group("/api")
	productApi := r.Group("/api-product")
	adminAuthRoutes := adminAuthApi.Group("/admin-auth")
	{
		adminAuthRoutes.POST("/login", adminAuthController.AdminLogin)
		adminAuthRoutes.POST("/add-user", authMiddleware, adminAuthController.AddAdminUsers)
	}
	productRoutes := productApi.Group("/product")
	{
		productRoutes.POST("/create", authMiddleware, middlewares.CheckAdminStatus(), productController.CreateNewDrawing)
		productRoutes.GET("/get", productController.GetAllProducts)
	}
}
