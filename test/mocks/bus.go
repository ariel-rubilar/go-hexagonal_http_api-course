package mocks

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/command"
	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/event"
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

type EventBusMock struct {
	mock.Mock
}

var _ event.Bus = (*EventBusMock)(nil)

func (m *EventBusMock) Publish(ctx context.Context, events []event.Event) error {
	args := m.Called(ctx, events)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}

func (m *EventBusMock) Subscribe(t event.Type, h event.Handler) error {

	args := m.Called(t, h)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}
