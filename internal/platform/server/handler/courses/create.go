package courses

import (
	"net/http"

	mooc "github.com/ariel-rubilar/go-hexagonal_http_api-course/internal"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/persitence"
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

		course := mooc.NewCourse(req.ID, req.Name, req.Duration)
		repo := persitence.NewCourseRepository()
		if err := repo.Save(course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Course created successfully",
			"data": gin.H{
				"id":       course.ID(),
				"name":     course.Name(),
				"duration": course.Duration(),
			},
		})
	}
}
