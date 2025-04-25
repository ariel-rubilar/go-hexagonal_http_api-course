package course

import (
	"context"
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/persistence/memdb"
)

type courseRepository struct {
	db *memdb.MemDB
}

var _ mooc.CourseRepository = (*courseRepository)(nil)

func NewCourseRepository(db *memdb.MemDB) *courseRepository {
	return &courseRepository{
		db: db,
	}
}

func (c *courseRepository) Save(ctx context.Context, course *mooc.Course) error {

	values := map[string]any{
		"id":       course.ID().String(),
		"name":     course.Name().String(),
		"duration": course.Duration().String(),
	}

	if err := c.db.InsertRow(memdb.CourseTableName, values); err != nil {
		return err
	}

	return nil
}

func (c *courseRepository) ListAll(ctx context.Context) ([]*mooc.Course, error) {

	courses := []*mooc.Course{}

	rows, err := c.db.List(memdb.CourseTableName)
	if err != nil {
		return nil, fmt.Errorf("error listing courses: %w", err)
	}
	for _, row := range rows {

		course, err := mooc.NewCourse(
			row["id"].(string),
			row["name"].(string),
			row["duration"].(string),
		)
		if err != nil {
			return nil, fmt.Errorf("error creating course: %w", err)
		}
		courses = append(courses, course)
	}

	return courses, nil
}
