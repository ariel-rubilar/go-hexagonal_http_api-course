package courses_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/server/handler/courses"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {

	courseRepository := new(mocks.CourseRepositoryMock)
	courseRepository.On("Create", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", courses.CreateHandler(courseRepository))

	t.Run("Given invalid request it return 400", func(t *testing.T) {

		createRequest := &courses.CreateRequest{
			Name:     "Course Name",
			Duration: "3 months",
		}

		json, err := json.Marshal(createRequest)

		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(json))

		require.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, res.StatusCode, w.Code)
	})

}
