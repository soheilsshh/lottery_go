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
	// بارگذاری فایل .env
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("هشدار: فایل .env یافت نشد، از مقادیر پیش‌فرض استفاده می‌شود")
		log.Println("Warning: .env file not found, using default values")
	}

	// اتصال به دیتابیس
	// Connect to database
	database.Connect()

	// اجرای migrations برای تمام مدل‌ها
	// Run migrations for all models
	if err := database.DB.AutoMigrate(
		&entities.User{},
		&entities.Draw{},
		&entities.DrawEntry{},
	); err != nil {
		log.Fatal("خطا در اجرای migrations: ", err)
		log.Fatal("Error running migrations: ", err)
	}

	// ایجاد کاربر ادمین در صورت عدم وجود
	// Create admin user if not exists
	seedAdminUser()

	// راه‌اندازی سرور Gin
	// Setup Gin server
	router := gin.Default()

	// مسیر اصلی برای بررسی سلامت API
	// Root route for health check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Lottery API is running!",
		})
	})

	// شروع سرور روی پورت 8080
	// Start server on port 8080
	log.Println("سرور روی پورت 8080 در حال اجرا است...")
	log.Println("Server is running on port 8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("خطا در راه‌اندازی سرور: ", err)
		log.Fatal("Error starting server: ", err)
	}
}

// seedAdminUser کاربر ادمین را ایجاد می‌کند در صورت عدم وجود
// seedAdminUser creates admin user if it doesn't exist
func seedAdminUser() {
	// بررسی وجود کاربر با ایمیل admin@lottery.com
	// Check if user with email admin@lottery.com exists
	var existingUser entities.User
	result := database.DB.Where("email = ?", "admin@lottery.com").First(&existingUser)

	// اگر کاربر وجود داشت، نیازی به ایجاد نیست
	// If user exists, no need to create
	if result.Error == nil {
		log.Println("کاربر ادمین از قبل وجود دارد")
		log.Println("Admin user already exists")
		return
	}

	// هش کردن رمز عبور
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("خطا در هش کردن رمز عبور: ", err)
		log.Fatal("Error hashing password: ", err)
	}

	// ایجاد کاربر ادمین
	// Create admin user
	adminUser := entities.User{
		Name:     "Admin",
		Email:    "admin@lottery.com",
		Password: string(hashedPassword),
		IsActive: true,
	}

	if err := database.DB.Create(&adminUser).Error; err != nil {
		log.Fatal("خطا در ایجاد کاربر ادمین: ", err)
		log.Fatal("Error creating admin user: ", err)
	}

	log.Println("کاربر ادمین با موفقیت ایجاد شد")
	log.Println("Admin user created successfully")
}
