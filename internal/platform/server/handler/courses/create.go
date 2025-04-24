package courses

import (
	"net/http"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := mooc.NewCourse(req.ID, req.Name, req.Duration)

		if err := courseRepository.Save(ctx, course); err != nil {
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
