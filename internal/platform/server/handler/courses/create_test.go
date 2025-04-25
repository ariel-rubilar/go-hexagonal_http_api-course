package courses_test

import (
	"bytes"
	"encoding/json"
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

func TestHandler_Create(t *testing.T) {

	courseService := new(mocks.CourseServiceMock)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", courses.CreateHandler(courseService))

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

	t.Run("Given invalid  id request it return 400", func(t *testing.T) {

		id, name, duration := "invalid-id", "Course Name", "3 months"

		courseService.On("Create", mock.Anything, id, name, duration).Return(nil, mooc.ErrInvalidCourseID)

		createRequest := &courses.CreateRequest{
			ID:       id,
			Name:     name,
			Duration: duration,
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

	t.Run("Given valid request it return 201", func(t *testing.T) {

		id, name, duration := "123e4567-e89b-12d3-a456-426614174000", "Course Name", "3 months"
		course, err := mooc.NewCourse(id, name, duration)
		require.NoError(t, err)
		courseService.On("Create", mock.Anything, id, name, duration).Return(course, nil)
		createRequest := &courses.CreateRequest{
			ID:       id,
			Name:     name,
			Duration: duration,
		}

		json, err := json.Marshal(createRequest)

		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(json))

		require.NoError(t, err)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		res := w.Result()
		defer res.Body.Close()

		assert.Equal(t, res.StatusCode, http.StatusCreated)
	})

}
