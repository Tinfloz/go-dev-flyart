package routes

import (
	"go-backend/controllers"
	"go-backend/middlewares"
	"go-backend/structs"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine, dbConn *structs.DbConn) {
	adminAuthController := controllers.NewAdminAuthController(dbConn)
	authMiddleware := middlewares.AuthMiddleware(dbConn)
	adminAuthApi := r.Group("/api")
	adminAuthRoutes := adminAuthApi.Group("/admin-auth")
	{
		adminAuthRoutes.POST("/login", adminAuthController.AdminLogin)
		adminAuthRoutes.POST("/add-user", authMiddleware, adminAuthController.AddAdminUsers)
	}
}
