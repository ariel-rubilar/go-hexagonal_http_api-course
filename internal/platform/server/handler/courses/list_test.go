package courses_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/server/handler/courses"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_List(t *testing.T) {

	courseRepository := new(mocks.CourseRepositoryMock)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/courses", courses.ListHandler(courseRepository))

	t.Run("return 200", func(t *testing.T) {

		courseRepository.On("ListAll", mock.Anything).Return(make([]*mooc.Course, 0), nil)

		req, err := http.NewRequest(http.MethodGet, "/courses", &bytes.Buffer{})

		require.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, res.StatusCode, http.StatusOK)
	})

	t.Run("return courses", func(t *testing.T) {

		coursesModel := make([]*mooc.Course, 1)
		course, err := mooc.NewCourse(
			"123e4567-e89b-12d3-a456-426614174000",
			"Course Name",
			"3 months",
		)
		require.NoError(t, err)
		coursesModel[0] = course

		courseRepository.On("ListAll", mock.Anything).Return(coursesModel, nil)

		req, err := http.NewRequest(http.MethodGet, "/courses", &bytes.Buffer{})

		require.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		res := w.Result()

		body, err := io.ReadAll(res.Body)

		require.NoError(t, err)

		defer res.Body.Close()

		response := &courses.ListResponse{}

		err = json.Unmarshal(body, response)
		require.NoError(t, err)

		assert.Equal(t, response.Data[0].ID, coursesModel[0].ID().String())

		assert.Equal(t, len(response.Data), 1)
	})

}
