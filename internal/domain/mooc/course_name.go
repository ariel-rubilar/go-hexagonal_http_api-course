package mooc

import (
	"errors"
	"fmt"
)

type CourseName struct {
	value string
}

var (
	ErrInvalidCourseName = errors.New("invalid course name")
)

func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, fmt.Errorf("%w: %s", ErrInvalidCourseName, value)
	}
	return CourseName{value: value}, nil
}

func (n CourseName) String() string {
	return n.value
}
