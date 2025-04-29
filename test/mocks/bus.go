package mocks

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/command"
	"github.com/stretchr/testify/mock"
)

type BuseMock struct {
	mock.Mock
}

var _ command.Bus = (*BuseMock)(nil)

func (m *BuseMock) Dispatch(ctx context.Context, cmd command.Command) (any, error) {
	args := m.Called(ctx, cmd)
	if args.Get(1) != nil {
		return nil, args.Get(1).(error)
	}
	return args.Get(0), nil
}

func (m *BuseMock) Register(t command.Type, h command.Handler) {
	args := m.Called(t, h)
	if args.Get(0) != nil {
		return
	}
}
