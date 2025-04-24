package mooc

import (
	"errors"
	"fmt"
)

type CourseDuration struct {
	value string
}

var (
	ErrInvalidCourseDuration = errors.New("invalid course name")
)

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, fmt.Errorf("%w: %s", ErrInvalidCourseName, value)
	}
	return CourseDuration{value: value}, nil
}

func (n CourseDuration) String() string {
	return n.value
}
