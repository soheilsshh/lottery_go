package routes

import (
	"lottery/internal/controllers"
	"lottery/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and configures all API routes
func SetupRouter() *gin.Engine {
	// Initialize router
	router := gin.Default()

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Lottery API is running!",
		})
	})

	// Initialize controllers
	authCtrl := controllers.NewAuthController()

	// API routes group
	api := router.Group("/api")

	// Authentication routes (no auth required)
	auth := api.Group("/auth")
	{
		auth.POST("/register", authCtrl.Register)
		auth.POST("/login", authCtrl.Login)
		auth.POST("/refresh", authCtrl.Refresh)
		auth.POST("/logout", authCtrl.Logout)
	}

	// Protected routes (require JWT authentication)
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// Get current user (test middleware)
		protected.GET("/me", func(c *gin.Context) {
			user, exists := middleware.GetUserFromContext(c)
			if !exists {
				c.JSON(401, gin.H{
					"error":   "unauthorized",
					"message": "User not found in context",
				})
				return
			}

			c.JSON(200, gin.H{
				"id":        user.ID,
				"name":      user.Name,
				"email":     user.Email,
				"is_active": user.IsActive,
			})
		})
	}

	return router
}
