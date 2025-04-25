package mooc_test

import (
	"testing"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/stretchr/testify/assert"
)

func TestCourseDuration_New(t *testing.T) {

	t.Run("given a empty duration should return error", func(t *testing.T) {
		duration, err := mooc.NewCourseDuration("")

		assert.ErrorIs(t, err, mooc.ErrInvalidCourseDuration)
		assert.Equal(t, mooc.CourseDuration{}, duration)
	})

	t.Run("given a valid duration should return a CourseDuration", func(t *testing.T) {
		duration, err := mooc.NewCourseDuration("3 months")

		assert.NoError(t, err)
		assert.Equal(t, "3 months", duration.String())
	})
}
