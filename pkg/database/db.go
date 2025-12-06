package database

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// DB متغیر سراسری اتصال به دیتابیس
// DB is the global database connection variable
var DB *gorm.DB

// Connect اتصال به دیتابیس SQLite را برقرار می‌کند
// Connect establishes a connection to SQLite database
// این تابع در صورت خطا panic می‌کند
// This function panics on error
func Connect() {
	// بارگذاری فایل .env
	// Load .env file
	if err := godotenv.Load(); err != nil {
		// اگر فایل .env وجود نداشت، از مقادیر پیش‌فرض استفاده می‌کنیم
		// If .env file doesn't exist, use default values
	}

	// خواندن مسیر دیتابیس از متغیر محیطی یا استفاده از مقدار پیش‌فرض
	// Read database path from environment variable or use default
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./lottery.db"
	}

	// اتصال به SQLite
	// Connect to SQLite
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("خطا در اتصال به دیتابیس: " + err.Error())
	}

	// تنظیمات connection pool (اختیاری برای SQLite)
	// Connection pool settings (optional for SQLite)
	sqlDB, err := DB.DB()
	if err != nil {
		panic("خطا در دریافت اتصال SQL: " + err.Error())
	}

	// تنظیم حداکثر تعداد اتصالات باز
	// Set maximum number of open connections
	sqlDB.SetMaxOpenConns(1) // SQLite فقط یک اتصال همزمان
	sqlDB.SetMaxIdleConns(1) // SQLite only supports one concurrent connection
}
