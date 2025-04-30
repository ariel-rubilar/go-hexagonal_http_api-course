package creating_test

import (
	"context"
	"errors"
	"testing"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/application/course/creating"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCourseService_Create_Fail(t *testing.T) {

	mockCourseRepository := new(mocks.CourseRepositoryMock)
	eventBus := new(mocks.EventBusMock)

	courseService := creating.NewCreatingService(mockCourseRepository, eventBus)

	t.Run("if course repository fails should return error", func(t *testing.T) {

		id, name, duration := "123e4567-e89b-12d3-a456-426614174000", "Go Programming", "3 months"

		mockCourseRepository.On("Save", mock.Anything, mock.AnythingOfType("*mooc.Course")).Return(errors.New("error")).Once()

		_, err := courseService.Create(context.Background(), id, name, duration)

		mockCourseRepository.AssertExpectations(t)

		assert.Error(t, err)
	})

	t.Run("if new course fails should return error", func(t *testing.T) {

		id, name, duration := "123e4567-e89b-12d3-a456-426614174000", "", "3 months"
		course, err := mooc.NewCourse(id, name, duration)
		require.Error(t, err)

		_, err = courseService.Create(context.Background(), id, name, duration)

		mockCourseRepository.AssertNotCalled(t, "Save", mock.Anything, course)
		assert.Error(t, err)
	})
}

func TestCourseService_Create_Success(t *testing.T) {

	mockCourseRepository := new(mocks.CourseRepositoryMock)
	mockEventBus := new(mocks.EventBusMock)

	courseService := creating.NewCreatingService(mockCourseRepository, mockEventBus)

	id, name, duration := "123e4567-e89b-12d3-a456-426614174000", "Go Programming", "3 months"

	mockCourseRepository.On("Save", mock.Anything, mock.AnythingOfType("*mooc.Course")).Return(nil)
	mockEventBus.On("Publish", mock.Anything, mock.AnythingOfType("[]event.Event")).Return(nil)

	newCourse, err := courseService.Create(context.Background(), id, name, duration)

	assert.NoError(t, err)

	assert.NotNil(t, newCourse)
	assert.Equal(t, id, newCourse.ID().String())
	assert.Equal(t, name, newCourse.Name().String())
	assert.Equal(t, duration, newCourse.Duration().String())
	mockCourseRepository.AssertExpectations(t)
	mockEventBus.AssertExpectations(t)
}
