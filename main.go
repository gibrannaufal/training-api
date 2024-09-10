package main

import (
	"github.com/gibrannaufal/training-api/Models/UserModels"
	"github.com/gibrannaufal/training-api/Routes/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Koneksi ke database MySQL
func initDB() *gorm.DB {
	dsn := "root:gibran123@tcp(127.0.0.1:3306)/training-db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// AutoMigrate untuk membuat tabel jika belum ada
	db.AutoMigrate(&UserModels.User{})

	return db
}

func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	db := initDB()

	user.UserRoutes(r, db)

	// Jalankan server di port 8001
	r.Run(":8001")
}
