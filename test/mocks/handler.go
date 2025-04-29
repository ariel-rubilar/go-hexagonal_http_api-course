package mocks

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/command"
	"github.com/stretchr/testify/mock"
)

type HandlerMock struct {
	mock.Mock
}

var _ command.Handler = (*HandlerMock)(nil)

func (m *HandlerMock) Handle(ctx context.Context, cmd command.Command) (any, error) {
	args := m.Called(ctx, cmd)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0), nil
}
