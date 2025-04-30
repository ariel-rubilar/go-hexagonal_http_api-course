package fetching

import (
	"context"
	"errors"
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/command"
)

var (
	ListCoursesCommandType = command.Type("command.list.courses")
)

type ListCoursesCommand struct {
}

func NewListCoursesCommand() *ListCoursesCommand {
	return &ListCoursesCommand{}
}

func (c *ListCoursesCommand) Type() command.Type {
	return ListCoursesCommandType
}

type ListCoursesCommandHandler struct {
	courseService FetchingService
}

func NewListCoursesCommandHandler(courseService FetchingService) *ListCoursesCommandHandler {
	return &ListCoursesCommandHandler{
		courseService: courseService,
	}
}

func (h *ListCoursesCommandHandler) Handle(ctx context.Context, cmd command.Command) (any, error) {
	_, ok := cmd.(*ListCoursesCommand)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid command type: %T", cmd))
	}

	courses, err := h.courseService.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	return courses, nil
}
