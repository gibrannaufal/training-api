package user

import (
	"github.com/gibrannaufal/training-api/Helpers/UtilsHelpers"
	"github.com/gibrannaufal/training-api/Models/UserModels"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UserRoutes(r *gin.Engine, db *gorm.DB) {
	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/get-users", func(c *gin.Context) {
			var users []UserModels.User
			var total int64

			page := c.DefaultQuery("page", "1")
			perPage := c.DefaultQuery("per_page", "5")
			nameQuery := c.DefaultQuery("name", "")

			pageNum, _ := strconv.Atoi(page)
			perPageNum, _ := strconv.Atoi(perPage)
			offset := (pageNum - 1) * perPageNum

			query := db.Model(&UserModels.User{})
			if nameQuery != "" {
				query = query.Where("name LIKE ?", "%"+nameQuery+"%")
			}

			query.Count(&total)
			query.Offset(offset).Limit(perPageNum).Find(&users)

			response := UserModels.PaginatedResponse{
				List: users,
			}
			response.Meta.Total = total

			UtilsHelpers.SuccessResponse(c, "User created successfully", response)
		})

		userRoutes.POST("/add-user", func(c *gin.Context) {
			var user UserModels.User

			// Bind JSON request body ke struct User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			result := db.Create(&user)
			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}

			// Kirim response dengan data
			UtilsHelpers.SuccessResponse(c, "User created successfully", user)
		})

		userRoutes.POST("/update-user", func(c *gin.Context) {
			var user UserModels.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Ambil ID dari model User
			if user.Id == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
				return
			}

			// Update data pengguna di database
			result := db.Model(&UserModels.User{}).Where("id = ?", *user.Id).Updates(map[string]interface{}{
				"name":     user.Name,
				"email":    user.Email,
				"password": user.Password,
				"foto_url": user.FotoURL,
			})

			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}

			UtilsHelpers.SuccessResponse(c, "User updated successfully", user)
		})

		userRoutes.DELETE("/delete-users/:id", func(c *gin.Context) {
			// Ambil ID dari parameter URL
			idStr := c.Param("id")
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
				return
			}

			// Hapus pengguna dari database
			result := db.Delete(&UserModels.User{}, id)

			if result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
				return
			}

			UtilsHelpers.SuccessResponse(c, "User deleted successfully", nil)
		})
	}
}
