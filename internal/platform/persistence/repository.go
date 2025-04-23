package persistence

import (
	"context"

	mooc "github.com/ariel-rubilar/go-hexagonal_http_api-course/internal"
	"github.com/hyperioxx/memsql"
)

var (
	tableName = "courses"
)

type courseRepository struct {
	db *memsql.Database
}

func NewCourseRepository() mooc.CourseRepository {
	db := memsql.NewDatabase()

	db.CreateTable(tableName, []*memsql.Column{
		{
			Name: "id",
			Kind: "uuid",
		},
		{
			Name: "name",
			Kind: "string",
		},
		{
			Name: "duration",
			Kind: "string",
		},
	})
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

	if err := c.db.InsertRow(tableName, values); err != nil {
		return err
	}

	return nil
}
