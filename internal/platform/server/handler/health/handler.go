package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"version": "1.0.0",
		})
	}
}
