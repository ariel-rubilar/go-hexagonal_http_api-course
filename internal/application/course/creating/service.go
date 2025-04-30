package creating

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
)

type CreatingService interface {
	Create(ctx context.Context, id, name, duration string) (*mooc.Course, error)
}

type creatingService struct {
	courseRepository mooc.CourseRepository
}

func NewCreatingService(courseRepository mooc.CourseRepository) *creatingService {
	return &creatingService{
		courseRepository: courseRepository,
	}
}

func (s *creatingService) Create(ctx context.Context, id, name, duration string) (*mooc.Course, error) {

	course, err := mooc.NewCourse(id, name, duration)

	if err != nil {
		return nil, err
	}

	if err := s.courseRepository.Save(ctx, course); err != nil {
		return nil, err
	}

	return course, nil
}
