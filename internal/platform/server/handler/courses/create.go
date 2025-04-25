package courses

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/application/course"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(service course.CourseCreate) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req CreateRequest

		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		c, err := service.Create(ctx, req.ID, req.Name, req.Duration)

		fmt.Printf("err: %v\n", err)
		if err != nil {

			switch {
			case errors.Is(err, mooc.ErrInvalidCourseID):
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid course ID",
				})
			case errors.Is(err, mooc.ErrInvalidCourseName):
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid course name",
				})
			case errors.Is(err, mooc.ErrInvalidCourseDuration):
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "Invalid course duration",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error",
				})
			}
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "Course created successfully",
			"data": gin.H{
				"id":       c.ID().String(),
				"name":     c.Name().String(),
				"duration": c.Duration().String(),
			},
		})
	}
}
