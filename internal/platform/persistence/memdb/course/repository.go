package course

import (
	"context"

	mooc "github.com/ariel-rubilar/go-hexagonal_http_api-course/internal"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/persistence/memdb"
	"github.com/hyperioxx/memsql"
)

type courseRepository struct {
	db *memsql.Database
}

func NewCourseRepository(db *memsql.Database) mooc.CourseRepository {
	return &courseRepository{
		db: db,
	}
}

func (c *courseRepository) Save(ctx context.Context, course *mooc.Course) error {

	values := map[string]any{
		"id":       course.ID(),
		"name":     course.Name(),
		"duration": course.Duration(),
	}

	if err := c.db.InsertRow(memdb.CourseTableName, values); err != nil {
		return err
	}

	return nil
}
