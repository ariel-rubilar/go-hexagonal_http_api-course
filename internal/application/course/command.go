package course

import (
	"context"
	"errors"
	"fmt"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/command"
)

var (
	CreateCourseCommandType = command.Type("command.create.course")
	ListCoursesCommandType  = command.Type("command.list.courses")
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

type ListCoursesCommand struct {
}

func NewListCoursesCommand() *ListCoursesCommand {
	return &ListCoursesCommand{}
}

func (c *ListCoursesCommand) Type() command.Type {
	return ListCoursesCommandType
}

type ListCoursesCommandHandler struct {
	courseService CourseListAll
}

func NewListCoursesCommandHandler(courseService CourseListAll) *ListCoursesCommandHandler {
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
