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

func (i CourseID) String() string {
	return i.value
}
