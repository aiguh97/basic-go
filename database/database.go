package database

import (
	"crypto/tls"
	"fmt"
	"log"
	"santrikoding/backend-api/config"
	"santrikoding/backend-api/models"

	gmysql "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load konfigurasi dari .env
	dbUser := config.GetEnv("DB_USER", "")
	dbPass := config.GetEnv("DB_PASS", "")
	dbHost := config.GetEnv("DB_HOST", "")
	dbPort := config.GetEnv("DB_PORT", "4000") // TiDB Cloud default port
	dbName := config.GetEnv("DB_NAME", "")

	// 1️⃣ Daftarkan TLS config
	err := gmysql.RegisterTLSConfig("tidb", &tls.Config{
		InsecureSkipVerify: true, // 🔥 lewati verifikasi sertifikat
	})
	if err != nil {
		log.Fatal("❌ Failed to register TLS config:", err)
	}

	// 2️⃣ Format DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&tls=tidb",
		dbUser, dbPass, dbHost, dbPort, dbName)

	// 3️⃣ Koneksi ke database
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Database connected successfully (TLS secure connection)")

	// 4️⃣ Auto migrate model
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("❌ Failed to migrate database:", err)
	}

	fmt.Println("✅ Database migrated successfully!")
}
