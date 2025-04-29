package mocks

import (
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/command"
	"github.com/stretchr/testify/mock"
)

type CommandMock struct {
	mock.Mock
}

var _ command.Command = (*CommandMock)(nil)

func (m *CommandMock) Type() command.Type {
	args := m.Called()
	if args.Get(0) == nil {
		return ""
	}
	return args.Get(0).(command.Type)
}
