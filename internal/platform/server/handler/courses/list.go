package courses

import (
	"net/http"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/application/course"
	"github.com/gin-gonic/gin"
)

type CourseResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"`
}

type ListResponse struct {
	Data    []CourseResponse `json:"data"`
	Message string           `json:"message"`
}

func ListHandler(service course.CourseListAll) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		courses, err := service.ListAll(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		curseResponse := make([]CourseResponse, len(courses))

		for i, course := range courses {
			curseResponse[i] = CourseResponse{
				ID:       course.ID().String(),
				Name:     course.Name().String(),
				Duration: course.Duration().String(),
			}
		}

		ctx.JSON(http.StatusOK, ListResponse{
			Data:    curseResponse,
			Message: "Courses retrieved successfully",
		})

	}
}
