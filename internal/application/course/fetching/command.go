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

type ListCommand struct {
}

func NewListCommand() *ListCommand {
	return &ListCommand{}
}

func (c *ListCommand) Type() command.Type {
	return ListCoursesCommandType
}

type ListCommandHandler struct {
	courseService FetchingService
}

func NewListCommandHandler(courseService FetchingService) *ListCommandHandler {
	return &ListCommandHandler{
		courseService: courseService,
	}
}

func (h *ListCommandHandler) Handle(ctx context.Context, cmd command.Command) (any, error) {
	_, ok := cmd.(*ListCommand)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid command type: %T", cmd))
	}

	courses, err := h.courseService.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	return courses, nil
}
