package mooc_test

import (
	"testing"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/stretchr/testify/assert"
)

func TestCourseName_New(t *testing.T) {

	t.Run("given a empty name should return error", func(t *testing.T) {
		name, err := mooc.NewCourseName("")

		assert.ErrorIs(t, err, mooc.ErrEmptyCourseName)
		assert.Equal(t, mooc.CourseName{}, name)
	})

	t.Run("given a valid name should return a CourseName", func(t *testing.T) {
		name, err := mooc.NewCourseName("Go Programming")

		assert.NoError(t, err)
		assert.Equal(t, "Go Programming", name.String())
	})

	t.Run("given a name with less than 3 characters should return error", func(t *testing.T) {
		name, err := mooc.NewCourseName("Go")

		assert.ErrorIs(t, err, mooc.ErrLengthCourseName)
		assert.Equal(t, mooc.CourseName{}, name)
	})
	t.Run("given a name with more than 100 characters should return error", func(t *testing.T) {
		name, err := mooc.NewCourseName("Go Programming Go Programming Go Programming Go Programming Go Programming Go Programming Go Programming Go Programming")

		assert.ErrorIs(t, err, mooc.ErrLengthCourseName)
		assert.Equal(t, mooc.CourseName{}, name)
	})

}
