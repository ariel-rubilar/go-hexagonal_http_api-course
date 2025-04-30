package creating

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
)

type CourseCreate interface {
	Create(ctx context.Context, id, name, duration string) (*mooc.Course, error)
}

type CourseService interface {
	CourseCreate
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

func (s *courseService) Create(ctx context.Context, id, name, duration string) (*mooc.Course, error) {

	course, err := mooc.NewCourse(id, name, duration)

	if err != nil {
		return nil, err
	}

	if err := s.courseRepository.Save(ctx, course); err != nil {
		return nil, err
	}

	return course, nil
}
