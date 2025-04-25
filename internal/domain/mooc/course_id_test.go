package mooc_test

import (
	"testing"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/stretchr/testify/assert"
)

func TestCourseId_New(t *testing.T) {

	t.Run("given a invalid id should return error", func(t *testing.T) {
		id, err := mooc.NewCourseID("")

		assert.ErrorIs(t, err, mooc.ErrInvalidCourseID)
		assert.Equal(t, mooc.CourseID{}, id)
	})

	t.Run("given a valid id should return a CourseID", func(t *testing.T) {
		id, err := mooc.NewCourseID("123e4567-e89b-12d3-a456-426614174000")

		assert.NoError(t, err)
		assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", id.String())
	})

}
