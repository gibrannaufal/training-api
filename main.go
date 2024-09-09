package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"training-api/models" // Import models dari folder models
)

// Koneksi ke database MySQL
func initDB() *gorm.DB {
	dsn := "root:gibran123@tcp(127.0.0.1:3306)/training-db"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}

	// AutoMigrate untuk membuat tabel jika belum ada
	db.AutoMigrate(&models.User{})

	return db
}

func main() {
	// Inisialisasi Gin
	r := gin.Default()

	// Inisialisasi database
	db := initDB()

	// Route untuk menambahkan user
	r.POST("/add-user", func(c *gin.Context) {
		var user models.User

		// Bind JSON request body ke struct user
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Simpan data user ke database
		result := db.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// Berikan respons sukses
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
	})

	// Jalankan server di port 8001
	r.Run(":8001")
}
