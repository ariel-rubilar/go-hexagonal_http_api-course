package mocks

import (
	"context"

	mooc "github.com/ariel-rubilar/go-hexagonal_http_api-course/internal"
	"github.com/stretchr/testify/mock"
)

type CourseRepositoryMock struct {
	mock.Mock
}

func (m *CourseRepositoryMock) Save(ctx context.Context, course *mooc.Course) error {
	args := m.Called(ctx, course)

	if args.Get(0) != nil {
		return args.Get(0).(error)
	}

	return nil
}
