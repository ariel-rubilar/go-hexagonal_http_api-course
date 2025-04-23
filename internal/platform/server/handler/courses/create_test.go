package courses_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/server/handler/courses"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_Create(t *testing.T) {

	courseRepository := new(mocks.CourseRepositoryMock)
	courseRepository.On("Create", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", courses.CreateHandler(courseRepository))

	t.Run("Given invalid request it return 400", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/courses", strings.NewReader(`{"id": "", "name": "Course Name", "duration": "3 months"}`))

		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}
