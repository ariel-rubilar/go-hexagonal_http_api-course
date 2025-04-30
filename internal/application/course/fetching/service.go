package fetching

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
)

type FetchingService interface {
	ListAll(ctx context.Context) ([]*mooc.Course, error)
}

type fetchingService struct {
	courseRepository mooc.CourseRepository
}

func NewFetchingService(courseRepository mooc.CourseRepository) *fetchingService {
	return &fetchingService{
		courseRepository: courseRepository,
	}
}

func (s *fetchingService) ListAll(ctx context.Context) ([]*mooc.Course, error) {
	courses, err := s.courseRepository.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	return courses, nil
}
