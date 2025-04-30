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

type CreateCourseCommand struct {
	name     string
	id       string
	duration string
}

func NewCreateCourseCommand(id, name, duration string) *CreateCourseCommand {
	return &CreateCourseCommand{
		name:     name,
		id:       id,
		duration: duration,
	}
}

func (c *CreateCourseCommand) Type() command.Type {
	return CreateCourseCommandType
}

type CreateCourseCommandHandler struct {
	courseService CourseCreate
}

func NewCreateCourseCommandHandler(courseService CourseCreate) *CreateCourseCommandHandler {
	return &CreateCourseCommandHandler{
		courseService: courseService,
	}
}

func (h *CreateCourseCommandHandler) Handle(ctx context.Context, cmd command.Command) (any, error) {
	courseCmd, ok := cmd.(*CreateCourseCommand)
	if !ok {
		return nil, errors.New(fmt.Sprintf("invalid command type: %T", cmd))
	}

	c, err := h.courseService.Create(ctx, courseCmd.id, courseCmd.name, courseCmd.duration)
	if err != nil {
		return nil, err
	}

	return c, nil
}
