package mocks

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/stretchr/testify/mock"
)

type CourseRepositoryMock struct {
	mock.Mock
}

var _ mooc.CourseRepository = (*CourseRepositoryMock)(nil)

func (m *CourseRepositoryMock) Save(ctx context.Context, course *mooc.Course) error {
	args := m.Called(ctx, course)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(error)
}

func (m *CourseRepositoryMock) ListAll(ctx context.Context) ([]*mooc.Course, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, nil
	}
	return args.Get(0).([]*mooc.Course), args.Error(1)
}
