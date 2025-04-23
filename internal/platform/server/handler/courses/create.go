package courses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type creeateRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateCourse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req creeateRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.Status(http.StatusCreated)
	}
}
