package course

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/mooc"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/platform/persistence/memdb"
	"github.com/hyperioxx/memsql"
)

type courseRepository struct {
	db *memsql.Database
}

var _ mooc.CourseRepository = (*courseRepository)(nil)

func NewCourseRepository(db *memsql.Database) *courseRepository {
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

func (c *courseRepository) ListAll(ctx context.Context) ([]*mooc.Course, error) {

	courses := []*mooc.Course{}

	rows, err := c.db.Execute(fmt.Sprintf("SELECT * FROM %s", memdb.CourseTableName))
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		jsonData, err := json.Marshal(row)
		if err != nil {
			return nil, err
		}
		var courseData map[string]any
		if err := json.Unmarshal(jsonData, &courseData); err != nil {
			return nil, err
		}
		id, ok := courseData["id"].(string)
		if !ok {
			return nil, err
		}
		course, err := mooc.NewCourse(
			id,
			id,
			id,
		)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}
