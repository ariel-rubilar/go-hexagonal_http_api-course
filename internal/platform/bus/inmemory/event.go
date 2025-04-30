package inmemory

import (
	"context"

	"github.com/ariel-rubilar/go-hexagonal_http_api-course/kit/event"
)

type EventBus struct {
	handlers map[event.Type]event.Handler
}

var _ event.Bus = (*EventBus)(nil)

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event.Type]event.Handler),
	}
}

func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, e := range events {
		if handler, ok := b.handlers[e.Type()]; ok {
			if err := handler.Handle(ctx, e); err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *EventBus) Subscribe(t event.Type, h event.Handler) error {
	if _, ok := b.handlers[event.Type(t)]; ok {
		return nil
	}
	b.handlers[event.Type(t)] = h
	return nil
}
