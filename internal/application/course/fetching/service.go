package fetching

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
)

type FetchingService interface {
	ListAll(ctx context.Context) ([]*mooc.Course, error)
}

type courseService struct {
	courseRepository mooc.CourseRepository
}

func NewCourseService(courseRepository mooc.CourseRepository) *courseService {
	return &courseService{
		courseRepository: courseRepository,
	}
}

func (s *courseService) ListAll(ctx context.Context) ([]*mooc.Course, error) {
	courses, err := s.courseRepository.ListAll(ctx)
	if err != nil {
		return nil, err
	}
	return courses, nil
}
