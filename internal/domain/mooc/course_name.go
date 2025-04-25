package mooc

import (
	"errors"
	"fmt"
)

type CourseName struct {
	value string
}

var (
	ErrEmptyCourseName   = errors.New("invalid course name")
	ErrLengthCourseName  = errors.New("course name must be between 3 and 100 characters")
	ErrInvalidCourseName = errors.New("invalid course name")
)

func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, fmt.Errorf("%w: %w: %s", ErrInvalidCourseName, ErrEmptyCourseName, value)
	}

	if len(value) < 3 || len(value) > 100 {
		return CourseName{}, fmt.Errorf("%w: %w: %s", ErrInvalidCourseName, ErrLengthCourseName, value)
	}

	return CourseName{value: value}, nil
}

func (n CourseName) String() string {
	return n.value
}
