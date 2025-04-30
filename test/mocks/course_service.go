package mocks

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/application/course/creating"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/stretchr/testify/mock"
)

type CourseServiceMock struct {
	mock.Mock
}

var _ creating.CreatingService = (*CourseServiceMock)(nil)

func (m *CourseServiceMock) Create(ctx context.Context, id, name, duration string) (*mooc.Course, error) {
	args := m.Called(ctx, id, name, duration)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	if args.Get(0) == nil {
		return nil, nil
	}
	return args.Get(0).(*mooc.Course), args.Error(1)
}

func (m *CourseServiceMock) ListAll(ctx context.Context) ([]*mooc.Course, error) {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return nil, nil
	}
	return args.Get(0).([]*mooc.Course), args.Error(1)
}
