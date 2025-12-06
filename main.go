package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"lottery/internal/entities"
	"lottery/pkg/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default values")
	}


	database.Connect()


	if err := database.DB.AutoMigrate(
		&entities.User{},
		&entities.Draw{},
		&entities.DrawEntry{},
	); err != nil {
		log.Fatal("Error running migrations: ", err)
	}


	seedAdminUser()

	// Setup Gin server
	router := gin.Default()


	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Lottery API is running!",
		})
	})


	// Start server on port 8080
	log.Println("Server is running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}


// seedAdminUser creates admin user if it doesn't exist
func seedAdminUser() {

	// Check if user with email admin@lottery.com exists
	var existingUser entities.User
	result := database.DB.Where("email = ?", "admin@lottery.com").First(&existingUser)


	// If user exists, no need to create
	if result.Error == nil {
		log.Println("Admin user already exists")
		return
	}


	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error hashing password: ", err)
	}


	// Create admin user
	adminUser := entities.User{
		Name:     "Admin",
		Email:    "admin@lottery.com",
		Password: string(hashedPassword),
		IsActive: true,
	}

	if err := database.DB.Create(&adminUser).Error; err != nil {
		log.Fatal("Error creating admin user: ", err)
	}

	log.Println("Admin user created successfully")
}
