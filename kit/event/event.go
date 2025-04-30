package event

import (
	"context"
	"time"
)

type Type string

type Event interface {
	ID() string
	AggregateID() string
	OccurredOn() time.Time
	Type() Type
}

type BaseEvent struct {
	id          string
	aggregateID string
	occurredOn  time.Time
}

func NewBaseEvent(id, aggregateID string, occurredOn time.Time) BaseEvent {
	return BaseEvent{
		id:          id,
		aggregateID: aggregateID,
		occurredOn:  occurredOn,
	}
}

func (e BaseEvent) ID() string {
	return e.id
}

func (e BaseEvent) AggregateID() string {
	return e.aggregateID
}

func (e BaseEvent) OccurredOn() time.Time {
	return e.occurredOn
}

type Handler interface {
	Handle(context.Context, Event) error
}

type Bus interface {
	Publish(context.Context, []Event) error
	Subscribe(Type, Handler) error
}
