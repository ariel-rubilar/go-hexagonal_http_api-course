package mooc

import (
	"errors"
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/internal/domain/uuid"
)

type CourseID struct {
	value string
}

var (
	ErrInvalidCourseID = errors.New("invalid course ID")
)

func NewCourseID(value string) (CourseID, error) {
	uuid, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}
	return CourseID{value: uuid.String()}, nil

}

type Course struct {
	id       CourseID
	name     string
	duration string
}

func NewCourse(id, name, duration string) (*Course, error) {

	courseID, err := NewCourseID(id)

	if err != nil {
		return nil, err
	}

	return &Course{
		id:       courseID,
		name:     name,
		duration: duration,
	}, nil
}

func (c *Course) ID() CourseID {
	return c.id
}

func (c *Course) Name() string {
	return c.name
}

func (c *Course) Duration() string {
	return c.duration
}
