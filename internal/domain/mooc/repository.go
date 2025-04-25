package mooc

import "context"

type CourseRepository interface {
	Save(ctx context.Context, course *Course) error
	ListAll(ctx context.Context) ([]*Course, error)
}
