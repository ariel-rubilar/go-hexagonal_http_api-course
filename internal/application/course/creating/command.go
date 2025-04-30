package creating

import (
	"context"
	"errors"
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/command"
)

var (
	CreateCourseCommandType = command.Type("command.create.course")
)

type CreateCommand struct {
	name     string
	id       string
	duration string
}

func NewCreateCommand(id, name, duration string) *CreateCommand {
	return &CreateCommand{
		name:     name,
		id:       id,
		duration: duration,
	}
}

func (c *CreateCommand) Type() command.Type {
	return CreateCourseCommandType
}

type CreateCourseCommandHandler struct {
	courseService CreatingService
}

func NewCreateCommandHandler(courseService CreatingService) *CreateCourseCommandHandler {
	return &CreateCourseCommandHandler{
		courseService: courseService,
	}
}

func (h *CreateCourseCommandHandler) Handle(ctx context.Context, cmd command.Command) (any, error) {
	courseCmd, ok := cmd.(*CreateCommand)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid command type: %T", cmd))
	}

	c, err := h.courseService.Create(ctx, courseCmd.id, courseCmd.name, courseCmd.duration)
	if err != nil {
		return nil, err
	}

	return c, nil
}
