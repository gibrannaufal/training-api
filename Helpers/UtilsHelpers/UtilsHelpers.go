package UtilsHelpers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Fungsi SuccessResponse untuk mengirimkan respons sukses
func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}
