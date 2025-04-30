package fetching

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
)

type CourseListAll interface {
	ListAll(ctx context.Context) ([]*mooc.Course, error)
}

type CourseService interface {
	CourseListAll
}

type courseService struct {
	courseRepository mooc.CourseRepository
}

var _ CourseService = (*courseService)(nil)

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
