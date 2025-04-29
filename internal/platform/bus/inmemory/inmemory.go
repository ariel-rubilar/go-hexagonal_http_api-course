package inmemory

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/command"
)

type CommandBus struct {
	handlers map[command.Type]command.Handler
}

var _ command.Bus = (*CommandBus)(nil)

func New() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) (any, error) {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return nil, nil
	}

	return handler.Handle(ctx, cmd)
}

func (b *CommandBus) Register(t command.Type, h command.Handler) {
	b.handlers[t] = h
}
